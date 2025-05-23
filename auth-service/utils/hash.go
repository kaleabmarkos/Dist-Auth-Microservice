package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed_bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed_bytes), err
}

func CheckPassword(password, hash string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}