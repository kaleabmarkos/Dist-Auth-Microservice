package routes

import (
	"Dist-Auth-MicroService/rbac-service/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRoute() http.Handler{
	r := chi.NewRouter()

	r.Post("/role", handlers.CreateRole)
	r.Get("/role", handlers.GetRole)
	r.Post("/check", handlers.CheckPermission)

	return r
}