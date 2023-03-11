package inmemrepo

import (
	"time"

	"github.com/google/uuid"
	"github.com/stovenn/gotodo/internal/core/domain"
)

type userRepository struct {
	db []*domain.User
}

func NewUserRepository() *userRepository {
	return &userRepository{db: []*domain.User{}}
}

func (u userRepository) FindAll() ([]*domain.User, error) {
	return u.db, nil
}

func (u userRepository) FindByID(id string) (*domain.User, error) {
	for _, user := range u.db {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, ErrNotFound
}

func (u *userRepository) Create(user *domain.User) (*domain.User, error) {
	id := uuid.New().String()
	created := &domain.User{
		ID:             id,
		FullName:       user.FullName,
		HashedPassword: user.HashedPassword,
		Email:          user.Email,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	u.db = append(u.db, created)
	return created, nil
}

func (u *userRepository) Update(user *domain.User) (*domain.User, error) {
	found, _ := u.FindByID(user.ID)
	*found = *user
	return user, nil
}

func (u *userRepository) DeleteByID(id string) error {
	for i, user := range u.db {
		if user.ID == id {
			u.db = append(u.db[:i], u.db[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
