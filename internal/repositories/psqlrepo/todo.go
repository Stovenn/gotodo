package psqlrepo

import (
	"github.com/stovenn/gotodo/internal/core/domain"
)

type todoRepository struct {
}

func NewTodoRepository() *todoRepository {
	return &todoRepository{}
}

func (t todoRepository) FindAll() ([]*domain.Todo, error) {
	rows, err := db.Queryx("SELECT id,title, completed , item_order FROM todos;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []*domain.Todo
	for rows.Next() {
		var todo domain.Todo
		if err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.Order,
		); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (t todoRepository) FindByID(id string) (*domain.Todo, error) {
	var foundTodo domain.Todo
	row := db.QueryRowx("SELECT id,title, completed , item_order FROM todos WHERE id = $1;", id)

	err := row.Scan(&foundTodo.ID, &foundTodo.Title, &foundTodo.Completed, &foundTodo.Order)
	if err != nil {
		return nil, err
	}
	return &foundTodo, nil
}

func (t todoRepository) FindByOrder(order int) (*domain.Todo, error) {
	var foundTodo domain.Todo
	row := db.QueryRowx("SELECT id,title, completed , item_order FROM todos WHERE item_order = $1;", order)

	err := row.Scan(&foundTodo.ID, &foundTodo.Title, &foundTodo.Completed, &foundTodo.Order)
	if err != nil {
		return nil, err
	}
	return &foundTodo, nil
}

func (t todoRepository) Create(todo *domain.Todo) (*domain.Todo, error) {
	var newTodo domain.Todo

	row := db.QueryRowx("INSERT INTO todos (title, completed, item_order) VALUES ($1, false, (SELECT count(item_order) FROM todos) +1) RETURNING id, title, completed, item_order", todo.Title)
	err := row.Scan(
		&newTodo.ID,
		&newTodo.Title,
		&newTodo.Completed,
		&newTodo.Order)
	if err != nil {
		return nil, err
	}

	return &newTodo, nil
}

func (t todoRepository) Update(id string, todo *domain.Todo) (*domain.Todo, error) {
	row := db.QueryRowx("UPDATE todos SET title = $1, completed = $2, item_order = $3 where id = $4 RETURNING id, title, completed, item_order ", todo.Title, todo.Completed, todo.Order, todo.ID)
	err := row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.Order)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t todoRepository) DeleteByID(id string) error {
	_, err := db.Exec("DELETE FROM todos where id = $1;", id)
	if err != nil {
		return err
	}
	return nil
}
