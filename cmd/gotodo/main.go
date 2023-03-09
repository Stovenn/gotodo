package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/stovenn/gotodo/internal/api"
	"github.com/stovenn/gotodo/internal/core/services"
	"github.com/stovenn/gotodo/internal/repositories/psqlrepo"
	"github.com/stovenn/gotodo/pkg/util"
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

	infoLogger := log.New(os.Stdout, "[INFO] ", log.LstdFlags)
	errLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)

	server, err := api.NewServer(config, todoService, userService, infoLogger, errLogger)
	if err != nil {
		log.Fatalf("cannot create server: %v\n", err)
	}

	fmt.Printf("Server listening on port %s\n", viper.Get("PORT"))
	err = server.Start()
	if err != nil {
		log.Fatalf("an error occured on the server: %v", err)
	}
}
