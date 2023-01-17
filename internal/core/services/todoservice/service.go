package todoservice

import "github.com/stovenn/gotodo/internal/core/ports"

type todoService struct {
	R ports.TodoRepository
}

func NewTodoService(r ports.TodoRepository) *todoService {
	return &todoService{R: r}
}
