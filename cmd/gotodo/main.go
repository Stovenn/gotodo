package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
		errLogger.Fatalf("cannot create server: %v\n", err)
	}

	go func() {
		fmt.Printf("Server listening on port %s\n", config.Port)
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errLogger.Fatalf("an error occured on the server: %v", err)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	infoLogger.Println("received terminate, graceful shutdown", sig)
	if err = server.Shutdown(ctx); err != nil {
		errLogger.Printf("error on server shutdown: %v", err)
	}
}
