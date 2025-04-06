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

	route := routes.NewRouter()

	log.Println("Auth service running on 8080")
	log.Fatal(http.ListenAndServe(":8080", route))
}