// Package app contains all business logic for work with user.
package app

import (
	"context"
	"errors"

	"rest-on-grpc-gateway/modules/user/internal/domain"
)

// Errors.
var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrMustDifferent   = errors.New("values must be different")
	ErrEmailExist      = errors.New("email exist")
)

//go:generate mockgen -source=app.go -destination mock.app.contracts_test.go -package app_test

type (
	// Repo interface for user database.
	Repo interface {
		CreateUser(ctx context.Context, newUser *domain.User) (*domain.User, error)
		GetUserByID(ctx context.Context, id int) (*domain.User, error)
		UpdateUserByID(ctx context.Context, id int, name, email string) (*domain.User, error)
		UpdateUserPasswordByID(ctx context.Context, id int, password []byte) error
		DeleteUserByID(ctx context.Context, id int) error
	}
	// PasswordHash module responsible for hashing password.
	PasswordHash interface {
		Hashing(password string) ([]byte, error)
		Compare(hashedPassword []byte, password []byte) bool
	}
)

// App structure business logic, contains all app methods.
type App struct {
	repo Repo
	hash PasswordHash
}

// New build and return new App.
func New(repo Repo, hash PasswordHash) *App {
	return &App{
		repo: repo,
		hash: hash,
	}
}
