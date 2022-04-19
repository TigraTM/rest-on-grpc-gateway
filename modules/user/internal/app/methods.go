package app

import (
	"context"
	"fmt"

	"rest-on-grpc-gateway/modules/user/internal/domain"
)

// CreateUser create user.
func (a *App) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	u, err := a.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("a.repo.CreateUser: %w", err)
	}

	return u, nil
}

// GetUserByID get user by id.
func (a *App) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	u, err := a.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetUserByID: %w", err)
	}

	return u, nil
}

// UpdateUserByID check exist user and update data about him by user id.
func (a *App) UpdateUserByID(ctx context.Context, userID int, name, email string) (u *domain.User, err error) {
	_, err = a.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("a.GetUserByID: %w", err)
	}

	u, err = a.repo.UpdateUserByID(ctx, userID, name, email)
	if err != nil {
		return nil, fmt.Errorf("a.repo.UpdateUserByID: %w", err)
	}

	return u, nil
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
		fmt.Println(user)
		fmt.Println(fmt.Sprintf("pass: %s, oldPass: %s", user.Password, oldPass))
		return ErrInvalidPassword
	}

	err = a.repo.UpdateUserPasswordByID(ctx, userID, newPass)
	if err != nil {
		return fmt.Errorf("a.repo.UpdateUserPasswordByID: %w", err)
	}

	return nil
}

// DeleteUserByID delete user by id if user exist.
func (a *App) DeleteUserByID(ctx context.Context, userID int) error {
	_, err := a.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("a.GetUserByID: %w", err)
	}

	err = a.repo.DeleteUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("a.repo.DeleteUserByID: %w", err)
	}

	return nil
}
