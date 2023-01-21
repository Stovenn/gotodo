package todoservice

import (
	"fmt"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
)

type todoService struct {
	R ports.TodoRepository
}

func NewTodoService(r ports.TodoRepository) *todoService {
	return &todoService{R: r}
}

func (t *todoService) ListTodos() ([]*domain.TodoResponse, error) {
	todos, err := t.R.FindAll()
	if err != nil {
		return nil, fmt.Errorf("todoservice.ListTodo: %v", err)
	}
	if len(todos) == 0 {
		return []*domain.TodoResponse{}, nil
	}

	var todoResponses []*domain.TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, todo.ToResponse())
	}
	return todoResponses, nil
}

func (t *todoService) FindTodoByID(id string) (*domain.TodoResponse, error) {
	todo, err := t.R.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("todoservice.FindTodoByID: %v", err)
	}
	return todo.ToResponse(), nil
}

func (t *todoService) AddTodo(r domain.TodoCreationRequest) (*domain.TodoResponse, error) {
	todo := domain.Todo{Title: r.Title}

	createdTodo, err := t.R.Create(todo)
	if err != nil {
		return nil, fmt.Errorf("todoservice.AddTodo: %v", err)
	}
	return createdTodo.ToResponse(), nil
}

func (t *todoService) UpdateTodo(r domain.TodoUpdateRequest) (*domain.TodoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoService) PartiallyUpdateTodo(r domain.TodoPartialUpdateRequest) (*domain.TodoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoService) DeleteTodo(id string) error {
	err := t.R.DeleteByID(id)
	if err != nil {
		return fmt.Errorf("todoservice.DeleteTodo: %v", err)
	}
	return nil
}
