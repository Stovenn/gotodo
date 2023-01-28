package psqlrepo

import (
	"github.com/jmoiron/sqlx"
	"github.com/stovenn/gotodo/internal/core/domain"
	"log"

	_ "github.com/lib/pq"
)

type todoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(driver, url string) *todoRepository {
	conn, err := sqlx.Connect(driver, url)
	if err != nil {
		log.Panicln(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Panicln(err)
	}

	return &todoRepository{db: conn}
}

func (t todoRepository) FindAll() ([]*domain.Todo, error) {
	rows, err := t.db.Queryx("SELECT id,title, completed , item_order FROM todos;")
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
	row := t.db.QueryRowx("SELECT id,title, completed , item_order FROM todos WHERE id = $1;", id)

	err := row.Scan(&foundTodo.ID, &foundTodo.Title, &foundTodo.Completed, &foundTodo.Order)
	if err != nil {
		return nil, err
	}
	return &foundTodo, nil
}

func (t todoRepository) FindByOrder(order int) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoRepository) Create(todo *domain.Todo) (*domain.Todo, error) {
	var newTodo domain.Todo

	row := t.db.QueryRowx("INSERT INTO todos (title, completed, item_order) VALUES ($1, false, (SELECT count(item_order) FROM todos) +1) RETURNING id, title, completed, item_order", todo.Title)
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

func (t todoRepository) Update(todo *domain.Todo) (*domain.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (t todoRepository) DeleteByID(id string) error {
	//TODO implement me
	panic("implement me")
}
