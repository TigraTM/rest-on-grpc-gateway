// Package app contains all business logic for work with user.
package app

import (
	"errors"
)

// Errors.
var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidPassword = errors.New("invalid password")
)

type (
	// Repo interface for user database.
	Repo interface{}
)

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
