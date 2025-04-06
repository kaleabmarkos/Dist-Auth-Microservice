package routes

import (
	"Dist-Auth-MicroService/auth-service/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler{
	r:= chi.NewRouter()
	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)
	return r
}