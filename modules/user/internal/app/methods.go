package app

import (
	"context"
	"fmt"

	"rest-on-grpc-gateway/modules/user/internal/domain"
)

var (
	count int
	users = make(map[int]*domain.User)
)

// CreateUser create user.
func (a *App) CreateUser(_ context.Context, user *domain.User) (*domain.User, error) {
	count++
	user.ID = count

	users[count] = user

	return user, nil
}

// GetUserByID get user by id.
func (a *App) GetUserByID(_ context.Context, id int) (*domain.User, error) {
	user, ok := users[id]
	if !ok {
		return nil, ErrNotFound
	}

	return user, nil
}

// UpdateUserByID check exist user and update data about him by user id.
func (a *App) UpdateUserByID(ctx context.Context, userID int, name, email string) (*domain.User, error) {
	user, err := a.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("a.GetUserByID: %w", err)
	}

	user.ID = userID
	user.Name = name
	user.Email = email
	users[userID] = user

	return user, nil
}

// UpdateUserPasswordByID update user password if:
// - user exist
// - old password compares with the current password.
func (a *App) UpdateUserPasswordByID(ctx context.Context, userID int, oldPass, newPass string) error {
	user, err := a.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("a.GetUserByID: %w", err)
	}

	if user.Password != oldPass {
		return ErrInvalidPassword
	}

	user.Password = newPass
	users[userID] = user

	return nil
}

// DeleteUserByID delete user by id if user exist.
func (a *App) DeleteUserByID(ctx context.Context, userID int) error {
	_, err := a.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("a.GetUserByID: %w", err)
	}

	delete(users, userID)

	return nil
}
