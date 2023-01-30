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
	arg := &domain.Todo{ID: "", Title: util.RandomString(12), Order: 0, Completed: false}

	createdTodo, err := r.Create(arg)
	expected := &domain.Todo{ID: "", Title: arg.Title, Order: len(r.db), Completed: false}

	assertCreation(t, expected, createdTodo, err)

	return createdTodo
}

func TestTodoRepository_Create(t *testing.T) {
	t.Cleanup(func() {
		r.db = []*domain.Todo{}
	})
	t.Run("creation with empty db", func(t *testing.T) {
		createRandomTodo(t)
	})

	t.Run("creation with existing todos in db", func(t *testing.T) {
		r.db = []*domain.Todo{
			{ID: "1", Title: "todo 1", Order: 1, Completed: false},
			{ID: "2", Title: "todo 2", Order: 2, Completed: false},
		}
		createRandomTodo(t)
	})
}

func assertCreation(t *testing.T, expected, got *domain.Todo, err error) {
	t.Helper()

	assert.NotEmpty(t, got)
	assert.NoError(t, err)

	assert.Equal(t, expected.Title, got.Title)
	assert.Equal(t, expected.Order, got.Order)
	assert.NotZero(t, got.ID)
}
func TestTodoRepository_Update(t *testing.T) {
	t.Cleanup(func() {
		r.db = []*domain.Todo{}
	})

	t.Run("update", func(t *testing.T) {
		r.db = []*domain.Todo{
			{ID: "1", Title: "todo 1", Order: 1, Completed: false},
			{ID: "2", Title: "todo 2", Order: 2, Completed: false},
		}
		arg := &domain.Todo{ID: "1", Title: "updated title", Order: 1, Completed: true}

		updatedTodo, err := r.Update(arg.ID, arg)
		expected := &domain.Todo{ID: arg.ID, Title: arg.Title, Order: arg.Order, Completed: arg.Completed}

		assertUpdate(t, expected, updatedTodo, err)
	})
}

func assertUpdate(t *testing.T, expected, got *domain.Todo, err error) {
	t.Helper()

	assert.NotEmpty(t, got)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestTodoRepository_FindAll(t *testing.T) {
	todo1 := createRandomTodo(t)
	todo2 := createRandomTodo(t)

	expected := []*domain.Todo{
		{ID: todo1.ID, Title: todo1.Title, Order: todo1.Order, Completed: todo1.Completed},
		{ID: todo2.ID, Title: todo2.Title, Order: todo2.Order, Completed: todo2.Completed},
	}

	todos, err := r.FindAll()

	assert.NotEmpty(t, todos)
	assert.NoError(t, err)
	assert.Equal(t, expected, todos)
}

func TestTodoRepository_FindByID(t *testing.T) {
	t.Cleanup(func() {
		r.db = []*domain.Todo{}
	})
	t.Run("given a todo id should return associated todo item", func(t *testing.T) {
		todo := createRandomTodo(t)
		expected := &domain.Todo{ID: todo.ID, Title: todo.Title, Order: todo.Order, Completed: todo.Completed}

		foundTodo, err := r.FindByID(todo.ID)

		assert.NotEmpty(t, foundTodo)
		assert.NoError(t, err)
		assert.Equal(t, expected, foundTodo)
	})

	t.Run("given an unknown todo id should not return a ErrNotFound error", func(t *testing.T) {
		id := "unknown"
		notFoundTodo, err := r.FindByID(id)

		assert.Empty(t, notFoundTodo)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrNotFound)
	})
}

func TestTodoRepository_FindByOrder(t *testing.T) {
	t.Run("given an order number should return associated todo item", func(t *testing.T) {
		t.Cleanup(func() {
			r.db = []*domain.Todo{}
		})

		_ = createRandomTodo(t)
		todo2 := createRandomTodo(t)

		expected := &domain.Todo{ID: todo2.ID, Title: todo2.Title, Order: todo2.Order, Completed: todo2.Completed}

		foundTodo, err := r.FindByOrder(todo2.Order)

		assert.NotEmpty(t, foundTodo)
		assert.NoError(t, err)
		assert.Equal(t, expected, foundTodo)
	})
	t.Run("given an unknown order should return an ErrNotFound error", func(t *testing.T) {
		foundTodo, err := r.FindByOrder(3)

		assert.Empty(t, foundTodo)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrNotFound)
	})
}

func TestTodoRepository_DeleteByID(t *testing.T) {
	t.Run("given a todo id should delete associated todo item", func(t *testing.T) {
		t.Cleanup(func() {
			r.db = []*domain.Todo{}
		})

		newtodo := createRandomTodo(t)
		err := r.DeleteByID(newtodo.ID)
		assert.NoError(t, err)

		notFoundTodo, err := r.FindByID(newtodo.ID)

		assert.Empty(t, notFoundTodo)
		assert.Error(t, err)
		assert.EqualError(t, err, "todo not found")
	})
	t.Run("given an unknown todo id should return an ErrNotFound error", func(t *testing.T) {
		err := r.DeleteByID("unknown")
		assert.Error(t, err)
		assert.EqualError(t, err, "todo not found")
	})
}
