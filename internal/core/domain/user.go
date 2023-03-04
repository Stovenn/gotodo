package domain

import (
	"time"
)

// User represents an todolist user
type User struct {
	ID             string    `db:"id"`
	FullName       string    `db:"full_name"`
	HashedPassword string    `db:"hashed_password"`
	Email          string    `db:"email"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
