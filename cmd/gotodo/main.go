package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stovenn/gotodo/internal/api"
	"github.com/stovenn/gotodo/internal/core/services/todoservice"
	"github.com/stovenn/gotodo/internal/repositories/psqlrepo"
	"github.com/stovenn/gotodo/pkg/util"
	"log"
)

func main() {
	err := util.SetupConfig()
	if err != nil {
		log.Fatalf("an error occured on the server: %v\n", err)
	}
	//branching adapters to ports
	repository := psqlrepo.NewTodoRepository()
	service := todoservice.NewTodoService(repository)
	handler := api.NewHandler(service)
	server := api.NewServer(handler)

	fmt.Printf("Server listening on port %s\n", viper.Get("PORT"))
	err = server.Start()
	if err != nil {
		log.Fatalf("an error occured on the server: %v", err)
	}
}
