package inmemrepo

import (
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var r *todoRepository

func init() {
	r = NewTodoRepository()
}

func createRandomTodo(t *testing.T) *domain.Todo {
	arg := domain.Todo{ID: "", Title: util.RandomString(12), Order: 0, Completed: false, Url: ""}

	createdTodo, err := r.Create(arg)
	expected := &domain.Todo{ID: "", Title: arg.Title, Order: len(r.db), Completed: false, Url: ""}

	assertCreation(t, expected, createdTodo, err)

	return createdTodo
}

func TestTodoRepository_Create(t *testing.T) {
	t.Cleanup(func() {
		r.db = []*domain.Todo{}
	})
	t.Run("with empty db", func(t *testing.T) {
		createRandomTodo(t)
	})

	t.Run("with existing todos in db", func(t *testing.T) {
		r.db = []*domain.Todo{
			{ID: "1", Title: "todo 1", Order: 1, Completed: false, Url: ""},
			{ID: "2", Title: "todo 2", Order: 2, Completed: false, Url: ""},
		}
		createRandomTodo(t)
	})
}

func assertCreation(t *testing.T, expected, got *domain.Todo, err error) {
	assert.NotEmpty(t, got)
	assert.NoError(t, err)

	assert.Equal(t, expected.Title, got.Title)
	assert.Equal(t, expected.Order, got.Order)
	assert.NotZero(t, got.ID)
}

func TestTodoRepository_FindAll(t *testing.T) {
	todo1 := createRandomTodo(t)
	todo2 := createRandomTodo(t)

	expected := []*domain.Todo{
		{ID: todo1.ID, Title: todo1.Title, Order: todo1.Order, Completed: todo1.Completed, Url: todo1.Url},
		{ID: todo2.ID, Title: todo2.Title, Order: todo2.Order, Completed: todo2.Completed, Url: todo2.Url},
	}

	todos, err := r.FindAll()

	assert.NotEmpty(t, todos)
	assert.NoError(t, err)

	assert.Equal(t, expected, todos)
}
