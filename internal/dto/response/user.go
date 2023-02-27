package response

import (
	"encoding/json"

	"github.com/stovenn/gotodo/internal/core/domain"
)

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

func ToResponse(u *domain.User) *UserResponse {
	return &UserResponse{
		ID:       u.ID,
		FullName: u.FullName,
		Email:    u.Email,
	}
}

type LoginResponse struct {
	AccessToken string        `json:"token"`
	User        *UserResponse `json:"user"`
}

func (tr LoginResponse) JSON() []byte {
	var b []byte
	var err error

	b, err = json.Marshal(&tr)
	if err != nil {
		return nil
	}
	return b
}
