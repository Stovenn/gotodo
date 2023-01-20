package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	"net/http"
)

type todoHandler struct {
	S ports.TodoService
}

func NewTodoHandler(todoService ports.TodoService) *todoHandler {
	return &todoHandler{S: todoService}
}

func (t *todoHandler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
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

func (t *todoHandler) HandleListTodo(w http.ResponseWriter, r *http.Request) {
	response, err := t.S.ListTodos()
	if err != nil {
		handleError(w, err)
		return
	}
	withJSON(w, http.StatusOK, response...)
}

func (t *todoHandler) HandlePatchTodo(w http.ResponseWriter, r *http.Request) {
	//TODO
	panic("implement me")
}

func (t *todoHandler) HandlePutTodo(w http.ResponseWriter, r *http.Request) {
	//TODO
	panic("implement me")
}

func (t *todoHandler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
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
