package main

import (
	"fmt"
	"github.com/stovenn/gotodo/internal/api"
	"github.com/stovenn/gotodo/internal/core/services/todoservice"
	"github.com/stovenn/gotodo/internal/repositories/inmemrepo"
	"os"
)

func main() {
	//branching adapters to ports
	repository := inmemrepo.NewTodoRepository()
	service := todoservice.NewTodoService(repository)
	handler := api.NewHandler(service)
	server := api.NewServer(handler)

	fmt.Println("Server listening on port :8080")
	err := server.Start()
	if err != nil {
		fmt.Printf("an error occured on the server: %v", err)
		os.Exit(1)
	}
}
