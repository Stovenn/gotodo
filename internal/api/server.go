package api

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func NewServer(todoHandler *TodoHandler, userHandler *UserHandler) *Server {
	r := mux.NewRouter().PathPrefix("/api/").Subrouter()
	todoRoutes := r.PathPrefix("/todos").Subrouter()
	todoRoutes.HandleFunc("/", todoHandler.HandleCreateTodo).Methods(http.MethodPost)
	todoRoutes.HandleFunc("/", todoHandler.HandleListTodo).Methods(http.MethodGet)
	todoRoutes.HandleFunc("/{id}", todoHandler.HandleFindTodoByID).Methods(http.MethodGet)
	todoRoutes.HandleFunc("/{id}", todoHandler.HandlePutTodo).Methods(http.MethodPut)
	todoRoutes.HandleFunc("/{id}", todoHandler.HandlePatchTodo).Methods(http.MethodPatch)
	todoRoutes.HandleFunc("/{id}", todoHandler.HandleDeleteTodo).Methods(http.MethodDelete)

	userRoutes := r.PathPrefix("/users").Subrouter()
	userRoutes.HandleFunc("/", userHandler.HandleSignUp).Methods(http.MethodPost)
	userRoutes.HandleFunc("/", userHandler.HandleLogin).Methods(http.MethodPost)
	return &Server{router: r}
}

func (server Server) Start() error {
	return http.ListenAndServe(":"+viper.GetString("PORT"), server.router)
}
