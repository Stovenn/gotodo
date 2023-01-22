package inmemrepo

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stovenn/gotodo/internal/core/domain"
)

var (
	ErrNotFound = errors.New("todo not found")
)

type todoRepository struct {
	db []*domain.Todo
}

func NewTodoRepository() *todoRepository {
	return &todoRepository{db: []*domain.Todo{}}
}

func (r *todoRepository) FindAll() ([]*domain.Todo, error) {
	return r.db, nil
}

func (r *todoRepository) FindByID(id string) (*domain.Todo, error) {
	for _, todo := range r.db {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, ErrNotFound
}

func (r *todoRepository) Create(todo domain.Todo) (*domain.Todo, error) {
	id := uuid.New().String()
	created := &domain.Todo{ID: id, Title: todo.Title, Order: len(r.db) + 1, Completed: false, Url: ""}

	r.db = append(r.db, created)
	return created, nil
}

func (r *todoRepository) Save(todo *domain.Todo) (*domain.Todo, error) {
	if todo.ID == "" {
		id := uuid.New().String()
		created := &domain.Todo{ID: id, Title: todo.Title, Order: len(r.db) + 1, Completed: false, Url: ""}

		r.db = append(r.db, created)
		return created, nil
	} else {
		found, _ := r.FindByID(todo.ID)
		found = todo
		return found, nil
	}
}

func (r *todoRepository) FindByOrder(order int) (*domain.Todo, error) {
	for _, todo := range r.db {
		if todo.Order == order {
			return todo, nil
		}
	}
	return nil, ErrNotFound
}

func (r *todoRepository) DeleteByID(id string) error {
	for i, todo := range r.db {
		if todo.ID == id {
			r.db = append(r.db[:i], r.db[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
