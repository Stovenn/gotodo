package services

import (
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/core/ports"
	"github.com/stovenn/gotodo/internal/dto/request"
	"github.com/stovenn/gotodo/internal/dto/response"
	"github.com/stovenn/gotodo/pkg/bcrypt"
	"github.com/stovenn/gotodo/pkg/token"
	"github.com/stovenn/gotodo/pkg/util"
)

type userService struct {
	R ports.UserRepository
}

func NewUserService(r ports.UserRepository) *userService {
	return &userService{R: r}
}

func (t *userService) SignUp(r request.UserCreationRequest) (*response.UserResponse, error) {
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

	return response.ToResponse(createdUser), nil
}

func (t *userService) Login(uc request.UserCredentials, m token.Maker, c util.Config) (*response.LoginResponse, error) {
	foundUser, err := t.R.FindByEmail(uc.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CheckPassword(uc.Password, foundUser.HashedPassword)
	if err != nil {
		return nil, err
	}

	token, err := m.CreateToken(uc.Email, c.TokenDuration)
	if err != nil {
		return nil, err
	}

	response := &response.LoginResponse{
		AccessToken: token,
		User:        response.ToResponse(foundUser),
	}

	return response, err
}
