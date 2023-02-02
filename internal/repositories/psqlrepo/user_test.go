package psqlrepo

import (
	"database/sql"
	"fmt"
	"github.com/stovenn/gotodo/internal/core/domain"
	"github.com/stovenn/gotodo/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createRandomUser(t *testing.T) *domain.User {
	arg := &domain.User{
		FullName:       util.RandomString(12),
		Email:          util.RandomEmail(10),
		HashedPassword: "secret",
	}

	createdUser, err := userRepo.Create(arg)
	assertUserCreation(t, arg, createdUser, err)

	return createdUser
}

func TestUserRepository_Create(t *testing.T) {
	createRandomUser(t)
}

func assertUserCreation(t *testing.T, expected, got *domain.User, err error) {
	t.Helper()

	assert.NotEmpty(t, got)
	assert.NoError(t, err)
	assert.Equal(t, expected.FullName, got.FullName)
	assert.Equal(t, expected.Email, got.Email)
	assert.Equal(t, expected.HashedPassword, got.HashedPassword)
	assert.NotZero(t, got.ID)
}

func TestUserRepository_Update(t *testing.T) {
	user := createRandomUser(t)
	arg := &domain.User{
		ID:             user.ID,
		FullName:       "new fullname",
		HashedPassword: "newsecret",
	}
	expected := &domain.User{
		ID:             arg.ID,
		FullName:       arg.FullName,
		Email:          user.Email,
		HashedPassword: arg.HashedPassword,
	}

	fmt.Println(user)
	updatedUser, err := userRepo.Update(arg)
	fmt.Println(updatedUser)
	assertUserUpdate(t, expected, updatedUser, err)
}

func assertUserUpdate(t *testing.T, expected, got *domain.User, err error) {
	t.Helper()

	assert.NotEmpty(t, got)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestUserRepository_FindByID(t *testing.T) {
	todo := createRandomUser(t)

	foundUser, err := userRepo.FindByID(todo.ID)
	assert.NotEmpty(t, foundUser)
	assert.NoError(t, err)
	assert.Equal(t, todo, foundUser)
}

func TestUserRepository_FindAll(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomUser(t)
	}

	users, err := userRepo.FindAll()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
}

func TestUserRepository_DeleteByID(t *testing.T) {
	user := createRandomUser(t)

	err := todoRepo.DeleteByID(user.ID)
	assert.NoError(t, err)

	foundTodo, err := todoRepo.FindByID(user.ID)
	assert.Empty(t, foundTodo)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}
