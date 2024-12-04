package repos

import (
	"context"

	"chai/database/sqlc"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(username string, email string, password string) error
	GetUserByEmail(email string) (*sqlc.User, error)
	GetUserByID(userID int32) (*sqlc.User, error)
}

type userRepositoryImpl struct {
	db *sqlc.Queries
}

func NewUserRepository(db *sqlc.Queries) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) CreateUser(name string, email string, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	_, err = r.db.CreateAccount(context.Background(), sqlc.CreateAccountParams{
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
	})

	return err
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*sqlc.User, error) {
	users, err := r.db.SelectAccountByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}

	if len(users) != 0 {
		return nil, nil
	}

	return &users[0], nil
}

func (r *userRepositoryImpl) GetUserByID(userID int32) (*sqlc.User, error) {
	user, err := r.db.SelectUserByID(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
