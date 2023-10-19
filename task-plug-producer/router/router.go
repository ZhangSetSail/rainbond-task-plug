package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

var r *chi.Mux

func InitRouterCli() {
	logrus.Infof("init router cli")
	r = chi.NewRouter()
	r.Use(middleware.Logger)
	initRouter()
	logrus.Infof("init router success")
}

func GetRouter() *chi.Mux {
	return r
}

func initRouter() {
	r.Mount("/producer", ProducerRouter())
	r.Mount("/normative", NormativeRouter())
}
