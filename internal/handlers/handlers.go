package handlers

import (
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
	//TODO
	panic("implement me")
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
