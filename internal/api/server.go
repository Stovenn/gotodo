package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/stovenn/gotodo/internal/core/ports"
	"github.com/stovenn/gotodo/pkg/token"
	"github.com/stovenn/gotodo/pkg/util"
)

type Server struct {
	config     util.Config
	router     *mux.Router
	tokenMaker token.Maker
	infoLogger *log.Logger
	errLogger  *log.Logger

	ports.TodoService
	ports.UserService
}

var validate *validator.Validate

func NewServer(config util.Config, todoService ports.TodoService, userService ports.UserService, infoLogger, errLogger *log.Logger) (*Server, error) {
	validate = validator.New()

	server := &Server{
		config:      config,
		infoLogger:  infoLogger,
		errLogger:   errLogger,
		TodoService: todoService,
		UserService: userService,
	}
	maker, err := token.NewPasetoMaker(config.SymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("error : %w", err)
	}

	server.tokenMaker = maker

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	r := mux.NewRouter().PathPrefix("/api/").Subrouter()

	todoRoutes := r.PathPrefix("/todos").Subrouter()
	todoRoutes.Use(authMiddleware(server.tokenMaker))
	todoRoutes.HandleFunc("/", server.HandleCreateTodo).Methods(http.MethodPost)
	todoRoutes.HandleFunc("/", server.HandleListTodo).Methods(http.MethodGet)
	todoRoutes.HandleFunc("/{id}", server.HandleFindTodoByID).Methods(http.MethodGet)
	todoRoutes.HandleFunc("/{id}", server.HandlePutTodo).Methods(http.MethodPut)
	todoRoutes.HandleFunc("/{id}", server.HandlePatchTodo).Methods(http.MethodPatch)
	todoRoutes.HandleFunc("/{id}", server.HandleDeleteTodo).Methods(http.MethodDelete)

	userRoutes := r.PathPrefix("/users").Subrouter()
	userRoutes.HandleFunc("/register", server.HandleSignUp).Methods(http.MethodPost)
	userRoutes.HandleFunc("/login", server.HandleLogin).Methods(http.MethodPost)

	server.router = r
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+viper.GetString("PORT"), s.router)
}
