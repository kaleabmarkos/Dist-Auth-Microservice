package handlers

import (
	"Dist-Auth-MicroService/rbac-service/models"
	"Dist-Auth-MicroService/shared/db"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateRole(w http.ResponseWriter, r *http.Request){
	var role models.Role
	json.NewDecoder(r.Body).Decode(&role)

	collection := db.GetCollection("roles")
	res, err := collection.InsertOne(context.TODO(), role)
	if err!=nil{
		http.Error(w, "Failed to create a Role", http.StatusInternalServerError)
		return
	}

	role.ID = res.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(role)
	
}

func GetRole(w http.ResponseWriter, r *http.Request){
	roleName := r.URL.Query().Get("role")
	if roleName == ""{
		http.Error(w, "Missing role param", http.StatusBadRequest)
		return
	}

	collection := db.GetCollection("roles")
	var role models.Role

	err := collection.FindOne(context.TODO(), bson.M{"name": roleName}).Decode(&role)
	if err != nil{
		http.Error(w, "Role can't be found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(role)

}

func CheckPermission(w http.ResponseWriter, r *http.Request){
	var req struct{
		Role		string	`json: "role"`
		Permission	string	`json: "permission"`
	}

	var role models.Role
	json.NewDecoder(r.Body).Decode(&req)

	collection := db.GetCollection("roles")

	err := collection.FindOne(context.TODO(), bson.M{"name": req.Role}).Decode(&role)
	if err != nil{
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}

	for _,perm := range role.Permissions{
		if perm == req.Permission{
			json.NewEncoder(w).Encode(map[string]bool{"allowed": true})
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]bool{"allowed":false})

}