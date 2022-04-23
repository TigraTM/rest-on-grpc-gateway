package app

import (
	"context"
	"fmt"
	"rest-on-grpc-gateway/modules/user/internal/domain"
	"strings"
)

// CreateUser create user.
func (a *App) CreateUser(ctx context.Context, name, email, password string) (*domain.User, error) {
	passHash, err := a.hash.Hashing(password)
	if err != nil {
		return nil, fmt.Errorf("a.hash.Hashing: %w", err)
	}

	user := domain.User{
		Name:         name,
		Email:        strings.ToLower(email),
		PasswordHash: passHash,
	}

	return a.repo.CreateUser(ctx, user)
}

// GetUserByID get user by id.
func (a *App) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return a.repo.GetUserByID(ctx, id)
}

// UpdateUserByID check exist user and update data about him by user id.
func (a *App) UpdateUserByID(ctx context.Context, userID int, name, email string) (u *domain.User, err error) {
	_, err = a.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("a.GetUserByID: %w", err)
	}

	return a.repo.UpdateUserByID(ctx, userID, name, email)
}

// UpdateUserPasswordByID update user password if:
// - user exist
// - old password compares with the current password
// - old password does not equal the new password.
func (a *App) UpdateUserPasswordByID(ctx context.Context, userID int, oldPass, newPass string) error {
	user, err := a.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("a.GetUserByID: %w", err)
	}

	if !a.hash.Compare(user.PasswordHash, []byte(oldPass)) {
		return ErrInvalidPassword
	}

	if a.hash.Compare(user.PasswordHash, []byte(newPass)) {
		return ErrMustDifferent
	}

	newPassHash, err := a.hash.Hashing(newPass)
	if err != nil {
		return fmt.Errorf("a.hash.Hashing: %w", err)
	}

	return a.repo.UpdateUserPasswordByID(ctx, userID, newPassHash)
}

// DeleteUserByID delete user by id if user exist.
func (a *App) DeleteUserByID(ctx context.Context, userID int) error {
	_, err := a.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("a.GetUserByID: %w", err)
	}

	return a.repo.DeleteUserByID(ctx, userID)
}
