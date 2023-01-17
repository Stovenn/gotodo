package todorepo

import "database/sql"

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository() *todoRepository {
	return &todoRepository{}
}
