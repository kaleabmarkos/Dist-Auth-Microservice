package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongo(){

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil{
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err!=nil{
		panic(err)
	}

	fmt.Println("Connected to MongoDB")
	MongoClient = client

}

func GetCollection(collection string) *mongo.Collection{
	return MongoClient.Database(os.Getenv("MONGO_DB")).Collection(collection)
}