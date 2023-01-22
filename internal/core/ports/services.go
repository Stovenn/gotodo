package ports

import "github.com/stovenn/gotodo/internal/core/domain"

type TodoService interface {
	ListTodos() ([]*domain.TodoResponse, error)
	FindTodoByID(id string) (*domain.TodoResponse, error)
	AddTodo(r domain.TodoCreationRequest) (*domain.TodoResponse, error)
	UpdateTodo(id string, r domain.TodoUpdateRequest) (*domain.TodoResponse, error)
	PartiallyUpdateTodo(id string, r domain.TodoPartialUpdateRequest) (*domain.TodoResponse, error)
	DeleteTodo(id string) error
}
