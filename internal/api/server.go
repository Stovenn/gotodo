package api

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
)

type Server struct {
	router *mux.Router
}

func NewServer(handler *Handler) *Server {
	r := mux.NewRouter().PathPrefix("/api/").Subrouter()
	todoRoutes := r.PathPrefix("/todos").Subrouter()
	todoRoutes.HandleFunc("/", handler.HandleCreateTodo).Methods(http.MethodPost)
	todoRoutes.HandleFunc("/", handler.HandleListTodo).Methods(http.MethodGet)
	todoRoutes.HandleFunc("/{id}", handler.HandleFindTodoByID).Methods(http.MethodGet)
	todoRoutes.HandleFunc("/{id}", handler.HandlePutTodo).Methods(http.MethodPut)
	todoRoutes.HandleFunc("/{id}", handler.HandlePatchTodo).Methods(http.MethodPatch)
	todoRoutes.HandleFunc("/{id}", handler.HandleDeleteTodo).Methods(http.MethodDelete)

	return &Server{router: r}
}

func (server Server) Start() error {
	return http.ListenAndServe(":"+viper.GetString("PORT"), server.router)
}
