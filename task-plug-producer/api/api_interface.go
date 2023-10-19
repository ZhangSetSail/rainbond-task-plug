package api

import "net/http"

type ProducerInterface interface {
	SendTask(w http.ResponseWriter, r *http.Request)
}

type NormativeInterface interface {
	RetrieveData(w http.ResponseWriter, r *http.Request)
}
