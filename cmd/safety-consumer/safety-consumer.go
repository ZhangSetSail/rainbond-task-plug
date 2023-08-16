package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/goodrain/rainbond-safety/model"
	"github.com/nats-io/nats.go"
	"net/http"
)

func main() {
	nc, _ := nats.Connect("47.93.219.143:10007")
	nc.QueueSubscribe("foo", "queue", func(m *nats.Msg) {
		var cdm model.CodeDetectionModel
		err := json.Unmarshal(m.Data, &cdm)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Received a message: %v\n", cdm)
	})
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	http.ListenAndServe(":12345", r)
}
