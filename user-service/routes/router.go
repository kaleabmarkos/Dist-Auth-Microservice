package routes

import (
	"net/http"
	"Dist-Auth-MicroService/user-service/handlers"
	"Dist-Auth-MicroService/user-service/middleware"
	

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler{
	r:=chi.NewRouter()
	r.Use(middleware.AuthMiddleware)

	r.Post("/", handlers.CreateUser)
	r.Get("/{id}", handlers.GetUserById)
	r.Put("/{id}/role", handlers.UpdateUserRole)

	return r

}