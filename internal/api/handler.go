package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	"net/http"
)

type Handler struct {
	S ports.TodoService
}

func NewHandler(todoService ports.TodoService) *Handler {
	return &Handler{S: todoService}
}

func (t *Handler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var request domain.TodoCreationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, err)
		return
	}

	response, err := t.S.AddTodo(request)
	if err != nil {
		handleError(w, err)
		return
	}
	withJSON(w, http.StatusCreated, response)
}

func (t *Handler) HandleListTodo(w http.ResponseWriter, r *http.Request) {
	response, err := t.S.ListTodos()
	if err != nil {
		handleError(w, err)
		return
	}
	withJSON(w, http.StatusOK, response...)
}

func (t *Handler) HandleFindTodoByID(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]
	response, err := t.S.FindTodoByID(todoId)
	if err != nil {
		handleError(w, err)
		return
	}
	withJSON(w, http.StatusOK, response)
}

func (t *Handler) HandlePatchTodo(w http.ResponseWriter, r *http.Request) {
	//TODO
	panic("implement me")
}

func (t *Handler) HandlePutTodo(w http.ResponseWriter, r *http.Request) {
	//TODO
	panic("implement me")
}

func (t *Handler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]
	err := t.S.DeleteTodo(todoId)
	if err != nil {
		handleError(w, err)
	}
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func withJSON(w http.ResponseWriter, statusCode int, response ...*domain.TodoResponse) {
	b, err := json.Marshal(&response)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(b)
}
