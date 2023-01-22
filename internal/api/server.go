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
	todoRoutes.HandleFunc("/", handler.HandleCreateTodo).Methods("POST")
	todoRoutes.HandleFunc("/", handler.HandleListTodo).Methods("GET")
	todoRoutes.HandleFunc("/{id}", handler.HandleFindTodoByID).Methods("GET")
	todoRoutes.HandleFunc("/{id}", handler.HandlePutTodo).Methods("PUT")
	todoRoutes.HandleFunc("/{id}", handler.HandleDeleteTodo).Methods("DELETE")

	return &Server{router: r}
}

func (server Server) Start() error {
	return http.ListenAndServe(":"+viper.GetString("PORT"), server.router)

}
