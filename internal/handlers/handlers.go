package handlers

import (
	"encoding/json"
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

	b, err := json.Marshal(&response)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(201)
	w.Write(b)
}

func (t *todoHandler) HandleListTodo(w http.ResponseWriter, r *http.Request) {
	//TODO
	panic("implement me")
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
	//TODO
	panic("implement me")
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}
