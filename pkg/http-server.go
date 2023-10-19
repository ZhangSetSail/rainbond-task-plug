package pkg

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func InitHttpServer(r http.Handler, port string) *http.Server {
	s := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			logrus.Errorf("Listen err: %s\n", err)
		}
	}()
	return s
}
