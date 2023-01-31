package psqlrepo

import (
	"github.com/stovenn/gotodo/internal/core/domain"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (t *userRepository) FindAll() ([]*domain.User, error) {
	rows, err := db.Queryx("SELECT id,full_name, email FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Email,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (t *userRepository) FindByID(id string) (*domain.User, error) {
	var foundUser domain.User
	row := db.QueryRowx("SELECT id,full_name, email FROM users WHERE id = $1;", id)

	err := row.Scan(
		&foundUser.ID,
		&foundUser.FullName,
		&foundUser.Email)
	if err != nil {
		return nil, err
	}
	return &foundUser, nil
}

func (t *userRepository) Create(user *domain.User) (*domain.User, error) {
	var newUser domain.User

	row := db.QueryRowx("INSERT INTO users (full_name, email, hashed_password) VALUES ($1, $2, $3) RETURNING id, full_name, email", user.FullName, user.Email, user.HashedPassword)
	err := row.Scan(
		&newUser.ID,
		&newUser.FullName,
		&newUser.Email)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (t *userRepository) Update(user *domain.User) (*domain.User, error) {
	row := db.QueryRowx("UPDATE users SET full_name = $1, hashed_password = $2 where id = $3 RETURNING id, full_name, email", user.FullName, user.HashedPassword, user.ID)
	err := row.Scan(
		&user.ID,
		&user.FullName,
		&user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (t *userRepository) DeleteByID(id string) error {
	_, err := db.Exec("DELETE FROM users where id = $1;", id)
	if err != nil {
		return err
	}
	return nil
}
