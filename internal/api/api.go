package api

import (
	"net/http"
)

type Response interface {
	JSON()
}

func handleError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}

func withJSON(w http.ResponseWriter, statusCode int, b []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(b)
}
