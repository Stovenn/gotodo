package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

const (
	serverPort = ":8080"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	server := &http.Server{
		Addr:         serverPort,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	fmt.Printf("Server listening on port %s\n", serverPort)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("an error occured on the server: %v", err)
		os.Exit(1)
	}
}
