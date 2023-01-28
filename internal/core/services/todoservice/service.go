package todoservice

import (
	"errors"
	"fmt"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
)

var (
	ErrOrderConflict = errors.New("todo error conflict")
)

type todoService struct {
	R ports.TodoRepository
}

func NewTodoService(r ports.TodoRepository) *todoService {
	return &todoService{R: r}
}

func (t *todoService) DisplayAllTodos() ([]*domain.TodoResponse, error) {
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

func (t *todoService) DisplayTodo(id string) (*domain.TodoResponse, error) {
	todo, err := t.R.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("todoservice.DisplayTodo: %v", err)
	}
	return todo.ToResponse(), nil
}

func (t *todoService) CreateTodo(r domain.TodoCreationRequest) (*domain.TodoResponse, error) {
	todo := &domain.Todo{Title: r.Title}

	createdTodo, err := t.R.Create(todo)
	if err != nil {
		return nil, fmt.Errorf("todoservice.CreateTodo: %v", err)
	}
	return createdTodo.ToResponse(), nil
}

func (t *todoService) UpdateTodo(id string, r domain.TodoUpdateRequest) (*domain.TodoResponse, error) {
	foundTodo, err := t.R.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("todoservice.UpdateTodo: %v", err)
	}
	foundTodo.Title = r.Title
	foundTodo.Completed = r.Completed
	if t.isOrderConflict(r.Order, foundTodo.ID) {
		return nil, fmt.Errorf("todoservice.PartiallyUpdateTodo: %v", ErrOrderConflict)
	}
	foundTodo.Order = r.Order

	updatedTodo, err := t.R.Update(foundTodo)
	if err != nil {
		return nil, fmt.Errorf("todoservice.UpdateTodo: %v", err)
	}
	return updatedTodo.ToResponse(), nil
}

func (t *todoService) PartiallyUpdateTodo(id string, r domain.TodoPartialUpdateRequest) (*domain.TodoResponse, error) {
	foundTodo, err := t.R.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("todoservice.PartiallyUpdateTodo: %v", err)
	}
	if r.Title != "" {
		foundTodo.Title = r.Title
	}
	foundTodo.Completed = r.Completed
	if r.Order != 0 {
		if t.isOrderConflict(r.Order, foundTodo.ID) {
			return nil, fmt.Errorf("todoservice.PartiallyUpdateTodo: %v", err)
		}
		foundTodo.Order = r.Order
	}

	updatedTodo, err := t.R.Update(foundTodo)
	if err != nil {
		return nil, fmt.Errorf("todoservice.UpdateTodo: %v", err)
	}
	return updatedTodo.ToResponse(), nil
}

func (t *todoService) DeleteTodo(id string) error {
	err := t.R.DeleteByID(id)
	if err != nil {
		return fmt.Errorf("todoservice.DeleteOneTodo: %v", err)
	}
	return nil
}

func (t *todoService) DeleteAllTodos() error {
	panic("implement me")
}

func (t *todoService) isOrderConflict(order int, updatedTodoID string) bool {
	foundTodo, _ := t.R.FindByOrder(order)
	if foundTodo != nil && foundTodo.ID != updatedTodoID {
		return true
	}
	return false
}
