package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/goodrain/rainbond-task-plug/task-plug-producer/controller"
)

// NormativeRouter -
func NormativeRouter() chi.Router {
	r := chi.NewRouter()
	r.Get("/retrieve_data", controller.GetManager().RetrieveData)
	return r
}
