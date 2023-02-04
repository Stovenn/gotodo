package domain

import (
	"encoding/json"
	"time"
)

type User struct {
	ID             string    `db:"id"`
	FullName       string    `db:"full_name"`
	HashedPassword string    `db:"hashed_password"`
	Email          string    `db:"email"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type UserResponse struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func (tr UserResponse) JSON() []byte {
	var b []byte
	var err error

	b, err = json.Marshal(&tr)
	if err != nil {
		return nil
	}
	return b
}

type UserCreationRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserCredentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

// ToResponse maps a Todo to a TodoResponse
func (t User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:       t.ID,
		FullName: t.FullName,
		Email:    t.Email,
	}
}
