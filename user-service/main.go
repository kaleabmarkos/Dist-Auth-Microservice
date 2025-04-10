package main

import (
	"Dist-Auth-MicroService/auth-service/routes"
	"Dist-Auth-MicroService/shared/config"
	"Dist-Auth-MicroService/shared/db"
	"log"
	"net/http"
)

func main(){
	config.LoadEnv()
	db.InitMongo()

	routes := routes.NewRouter()

	log.Println("Serving on port 8081")
	log.Fatal(http.ListenAndServe(":8081", routes))

}