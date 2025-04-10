package main

import (
	"log"
	"net/http"

	"Dist-Auth-MicroService/auth-service/routes"
	"Dist-Auth-MicroService/shared/config"
	"Dist-Auth-MicroService/shared/db"
)

func main(){

	config.LoadEnv()
	db.InitMongo()

	router := routes.NewRouter()

	log.Println("Service running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}