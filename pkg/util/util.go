package util

import (
	"github.com/stovenn/gotodo/internal/core/domain"
	"math/rand"
	"strings"
	"time"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
func CreateRandomTodo(order int) *domain.Todo {
	return &domain.Todo{
		ID:        RandomString(15),
		Title:     RandomString(25),
		Order:     order,
		Completed: false,
		Url:       RandomString(50),
	}
}

func CreateRandomTodos(n int) []*domain.Todo {
	var todos []*domain.Todo
	for i := 0; i < n; i++ {
		todos = append(todos, CreateRandomTodo(i))
	}
	return todos
}

func CreateRandomTodoResponse(order int) *domain.TodoResponse {
	return CreateRandomTodo(order).ToResponse()
}

func CreateRandomTodoResponses(n int) []*domain.TodoResponse {
	var todoResponses []*domain.TodoResponse
	for i := 0; i < n; i++ {
		todoResponses = append(todoResponses, CreateRandomTodoResponse(i))
	}
	return todoResponses
}
