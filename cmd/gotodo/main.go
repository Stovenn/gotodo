package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stovenn/gotodo/internal/api"
	"github.com/stovenn/gotodo/internal/core/services"
	"github.com/stovenn/gotodo/internal/repositories/psqlrepo"
	"github.com/stovenn/gotodo/pkg/util"
	"log"
)

func main() {
	config, err := util.SetupConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v\n", err)
	}

	err = psqlrepo.OpenDB(config.DBDriver, config.DBUrl)
	if err != nil {
		log.Fatalf("cannot connect to DB: %v\n", err)
	}

	//branching adapters to ports
	todoRepository := psqlrepo.NewTodoRepository()
	userRepository := psqlrepo.NewUserRepository()

	todoService := services.NewTodoService(todoRepository)
	userService := services.NewUserService(userRepository)

	server, err := api.NewServer(config, todoService, userService)
	if err != nil {
		log.Fatalf("cannot create server: %v\n", err)
	}

	fmt.Printf("Server listening on port %s\n", viper.Get("PORT"))
	err = server.Start()
	if err != nil {
		log.Fatalf("an error occured on the server: %v", err)
	}
}
