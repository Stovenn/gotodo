package services

import (
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	"github.com/stovenn/gotodo/pkg/bcrypt"
)

type userService struct {
	R ports.UserRepository
}

func NewUserService(r ports.UserRepository) *userService {
	return &userService{R: r}
}

func (t *userService) SignUp(r domain.UserCreationRequest) (*domain.UserResponse, error) {
	hashedPassword, err := bcrypt.HashPassword(r.Password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		FullName:       r.FullName,
		Email:          r.Email,
		HashedPassword: hashedPassword,
	}
	createdUser, err := t.R.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser.ToResponse(), nil
}

func (t *userService) Login(uc domain.UserCredentials) error {
	//foundUser, err := t.R.FindByEmail(uc.Email)
	//if err != nil {
	//	return err
	//}
	//if err = bcrypt.CheckPassword(uc.Password, foundUser.HashedPassword); err != nil {
	//	return err
	//}
	//
	panic("implement me")
}
