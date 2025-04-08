package handlers

import (
	"Dist-Auth-MicroService/shared/db"
	"Dist-Auth-MicroService/user-service/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	user.CreateAt = time.Now()

	collection := db.GetCollection("users")

	res, err := collection.InsertOne(context.TODO(), user)
	if err!=nil{
		http.Error(w, "Failed to create the user", http.StatusInternalServerError)
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(user)
}

func GetUserById(w http.ResponseWriter, r *http.Request){
	idParam := chi.URLParam(r, "id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err!=nil{
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	collection := db.GetCollection("users")
	var user models.User

	err = collection.FindOne(context.TODO(), bson.M{"id": objectID}).Decode(&user)
	if err!=nil{
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func UpdateUserRole(w http.ResponseWriter, r *http.Request){
	idParam := chi.URLParam(r, "id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err!=nil{
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var payload struct{
		Role	string	`json: "role"`
	}

	json.NewDecoder(r.Body).Decode(&payload)

	collection := db.GetCollection("users")
	_, err = collection.UpdateOne(context.TODO(),
		bson.M{"id":objectID},
		bson.M{"$set": bson.M{"role": payload.Role}})
	if err!=nil{
		http.Error(w, "Falied to update Role. ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}