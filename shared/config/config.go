package config

import ("os"
		"log"
		"github.com/joho/godotenv"
)

func LoadEnv(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading env file")
	}
}
