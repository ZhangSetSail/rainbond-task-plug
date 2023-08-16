package api_router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func InitRouter() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return http.ListenAndServe(":12345", r)
}
