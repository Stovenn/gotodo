package services

import (
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
)

type userService struct {
	R ports.UserRepository
}

func NewUserService(r ports.UserRepository) *userService {
	return &userService{R: r}
}

func (t *userService) SignUp(r domain.UserCreationRequest) {
	//TODO implement me
	panic("implement me")
}

func (t *userService) Login(uc domain.UserCredentials) error {
	//TODO implement me
	panic("implement me")
}
