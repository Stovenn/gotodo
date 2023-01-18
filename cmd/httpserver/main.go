package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stovenn/gotodo/internal/core/services/todoservice"
	"github.com/stovenn/gotodo/internal/handlers"
	"github.com/stovenn/gotodo/internal/repositories/inmemrepo"
	"net/http"
	"os"
	"time"
)

const (
	serverPort = ":8080"
)

func main() {
	//branching adapters to ports
	repository := inmemrepo.NewTodoRepository()
	service := todoservice.NewTodoService(repository)
	handler := handlers.NewTodoHandler(service)

	//configure routes
	r := mux.NewRouter().PathPrefix("/api/").Subrouter()
	todoRoutes := r.PathPrefix("/todos").Subrouter()
	todoRoutes.HandleFunc("/", handler.HandleCreateTodo).Methods("POST")
	todoRoutes.HandleFunc("/", handler.HandleListTodo).Methods("GET")

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
