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
)

//go:generate mockgen -source=app.go -destination mock.app.contracts_test.go -package app_test

// Repo interface for user database.
type Repo interface {
	CreateUser(ctx context.Context, newUser *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	UpdateUserByID(ctx context.Context, id int, name, email string) (*domain.User, error)
	UpdateUserPasswordByID(ctx context.Context, id int, password string) error
	DeleteUserByID(ctx context.Context, id int) error
}

// App structure business logic, contains all app methods.
type App struct {
	repo Repo
}

// New build and return new App.
func New(repo Repo) *App {
	return &App{
		repo: repo,
	}
}
