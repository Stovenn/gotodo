package inmemrepo

import "github.com/stovenn/gotodo/internal/core/domain"

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

func (u userRepository) Create(todo *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Update(todo *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) DeleteByID(id string) error {
	//TODO implement me
	panic("implement me")
}
