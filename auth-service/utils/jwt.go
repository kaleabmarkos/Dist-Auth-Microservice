package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//authenticates API requests (short lifetime) , with every api req
func GenerateToken(user_ID string) (string,error){
	claims := jwt.MapClaims{
		"user_id": user_ID,
		"exp":	time.Now().Add(time.Minute*15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

//used ot get new access token (longer) , to renew session safely
func GenerateRefreshToken(user_ID string) (string, error){
	claims := jwt.MapClaims{
		"user_id": user_ID,
		"exp" : time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
}