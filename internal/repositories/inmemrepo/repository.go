package inmemrepo

import (
	"github.com/google/uuid"
	"github.com/stovenn/gotodo/internal/core/domain"
)

type todoRepository struct {
	db map[string]*domain.Todo
}

func NewTodoRepository() *todoRepository {
	return &todoRepository{
		db: map[string]*domain.Todo{},
	}
}

func (r todoRepository) FindAll() (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (r todoRepository) FindByID(id string) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (r todoRepository) Create(todo domain.Todo) (*domain.Todo, error) {
	id := uuid.New().String()
	created := &domain.Todo{Id: id, Title: todo.Title, Order: len(r.db) + 1, Completed: false, Url: ""}

	r.db[id] = created

	return created, nil
}

func (r todoRepository) Update(todo domain.Todo) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (r todoRepository) DeleteByID(todo domain.Todo) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}
