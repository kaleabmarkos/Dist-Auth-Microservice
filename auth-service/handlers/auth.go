package handlers

import (
	"Dist-Auth-MicroService/auth-service/models"
	"Dist-Auth-MicroService/auth-service/utils"
	"Dist-Auth-MicroService/shared/db"
	"context"
	"encoding/json"
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

func Login(w http.ResponseWriter, r *http.Request){
	var input AuthInput

	json.NewDecoder(r.Body).Decode(&input)

	collection := db.GetCollection("user")

	var user models.User

	err := collection.FindOne(context.TODO(), bson.M{"email":input.Email}).Decode(&user)
	if err !=nil{
		http.Error(w, "Wrong Credentials", http.StatusBadRequest)
		return 
	}

	if !utils.CheckPassword(input.Password, user.Password){
		http.Error(w, "Wrong Credentials", http.StatusUnauthorized)
		return
	}

	accessToken, _ := utils.GenerateToken(user.ID.Hex())
	refershToken, _ := utils.GenerateRefreshToken(user.ID.Hex())

	json.NewEncoder(w).Encode(map[string]string{
		"access_token" : accessToken,
		"refresh_token" : refershToken,
	})
}