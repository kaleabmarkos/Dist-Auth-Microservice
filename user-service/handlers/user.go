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


}

func UpdateUserRole(w http.ResponseWriter, r *http.Request){


}