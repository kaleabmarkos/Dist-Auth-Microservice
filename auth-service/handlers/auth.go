package handlers

import (
	"Dist-Auth-MicroService/auth-service/models"
	"Dist-Auth-MicroService/auth-service/utils"
	"Dist-Auth-MicroService/shared/db"
	"auth-service/models"
	"auth-service/utils"
	"context"
	"encoding/json"
	"go/printer"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthInput struct {
	Email		string `json:"string"`
	Password	string `json:"String"`
}

func Register(w http.ResponseWriter, r *http.Request){
	var input AuthInput
	json.NewDecoder(r.Body).Decode(&input)

	collection := db.GetCollection("users")

	count, _ := collection.CountDocuments(context.TODO(), bson.M{"email": input.Email})

	if count > 0{
		http.Error(w, "User already Exists", http.StatusBadRequest)
		return
	}

	hashedPassword, _ := utils.HashPassword(input.Password)
	user := models.User{
		Email: input.Email,
		Password: hashedPassword,
	}

	res, _ := collection.InsertOne(context.TODO(), user)
	user.ID = res.InsertedID.(primitive.ObjectID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}
