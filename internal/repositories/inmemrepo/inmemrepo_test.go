package inmemrepo

import (
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

var r *todoRepository

func TestTodoRepository_Create(t *testing.T) {
	r = NewTodoRepository()
	arg := domain.Todo{
		Id:        "",
		Title:     "new todo",
		Order:     0,
		Completed: false,
		Url:       "",
	}
	createdTodo, err := r.Create(arg)

	expected := domain.Todo{
		Id:        "",
		Title:     "new todo",
		Order:     1,
		Completed: false,
		Url:       "",
	}

	assert.NotEmpty(t, createdTodo)
	assert.NoError(t, err)

	assert.Equal(t, expected.Title, createdTodo.Title)
	assert.Equal(t, expected.Order, createdTodo.Order)
	assert.NotZero(t, createdTodo.Id)
}
