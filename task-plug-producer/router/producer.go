package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/goodrain/rainbond-task-plug/task-plug-producer/controller"
)

// ProducerRouter -
func ProducerRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/send_task", controller.GetManager().SendTask)
	return r
}
