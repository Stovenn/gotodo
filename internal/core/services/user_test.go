package services

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/internal/dto/request"
	"github.com/stovenn/gotodo/internal/dto/response"
	mockdb "github.com/stovenn/gotodo/internal/repositories/mock"
	"github.com/stovenn/gotodo/pkg/bcrypt"
	"github.com/stretchr/testify/assert"
)

func TestUserService_SignUp(t *testing.T) {
	arg := request.UserCreationRequest{
		FullName: "John Doe",
		Email:    "unknown@email.com",
		Password: "password",
	}
	hashedPassword, err := bcrypt.HashPassword(arg.Password)
	assert.NoError(t, err)
	user := &domain.User{
		ID:             "1",
		FullName:       arg.FullName,
		Email:          arg.Email,
		HashedPassword: hashedPassword,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	expectedResponse := response.ToResponse(user)
	ctrl := gomock.NewController(t)
	repository := mockdb.NewMockUserRepository(ctrl)
	us = NewUserService(repository)
	repository.EXPECT().Create(gomock.Any()).Times(1).Return(user, nil)

	response, err := us.SignUp(arg)

	assert.NotEmpty(t, response)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}
