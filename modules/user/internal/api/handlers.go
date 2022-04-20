package api

import (
	"context"
	"errors"
	"fmt"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
	"rest-on-grpc-gateway/modules/user/internal/app"
)

// CreateUser implements userpb.UserAPIServer.
func (a *api) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user, err := a.app.CreateUser(ctx, toDomain(in))
	if err != nil {
		return nil, fmt.Errorf("a.app.CreateUser: %w", err)
	}

	return &userpb.CreateUserResponse{
		Id:       int64(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

// GetUserByID implements userpb.UserAPIServer.
func (a *api) GetUserByID(ctx context.Context, in *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	user, err := a.app.GetUserByID(ctx, int(in.Id))
	switch {
	case err == nil:
		return &userpb.GetUserByIDResponse{
			Id:    int64(user.ID),
			Name:  user.Name,
			Email: user.Email,
		}, nil
	case errors.Is(err, app.ErrNotFound):
		return nil, errUserNotFound
	default:
		return nil, fmt.Errorf("a.app.GetUserByID: %w", err)
	}
}

// UpdateUserByID implements userpb.UserAPIServer.
func (a *api) UpdateUserByID(ctx context.Context, in *userpb.UpdateUserByIDRequest) (*userpb.UpdateUserByIDResponse, error) {
	user, err := a.app.UpdateUserByID(ctx, int(in.Id), in.Name, in.Email)
	switch {
	case err == nil:
		return &userpb.UpdateUserByIDResponse{
			Id:    int64(user.ID),
			Name:  user.Name,
			Email: user.Email,
		}, nil
	case errors.Is(err, app.ErrNotFound):
		return nil, errUserNotFound
	default:
		return nil, fmt.Errorf("a.app.UpdateUserByID: %w", err)
	}
}

// UpdateUserPasswordByID implements userpb.UserAPIServer.
func (a *api) UpdateUserPasswordByID(ctx context.Context, in *userpb.UpdateUserPasswordByIDRequest) (*userpb.UpdateUserPasswordByIDResponse, error) {
	err := a.app.UpdateUserPasswordByID(ctx, int(in.Id), in.OldPassword, in.NewPassword)
	switch {
	case err == nil:
		return &userpb.UpdateUserPasswordByIDResponse{}, nil
	case errors.Is(err, app.ErrNotFound):
		return nil, errUserNotFound
	case errors.Is(err, app.ErrInvalidPassword):
		return nil, errInvalidPassword
	default:
		return nil, fmt.Errorf("a.app.UpdateUserPasswordByID: %w", err)
	}
}

// DeleteUserByID implements userpb.UserAPIServer.
func (a *api) DeleteUserByID(ctx context.Context, in *userpb.DeleteUserByIDRequest) (*userpb.DeleteUserByIDResponse, error) {
	err := a.app.DeleteUserByID(ctx, int(in.Id))
	switch {
	case err == nil:
		return &userpb.DeleteUserByIDResponse{}, nil
	case errors.Is(err, app.ErrNotFound):
		return nil, errUserNotFound
	default:
		return nil, fmt.Errorf("a.app.DeleteUserByID: %w", err)
	}
}
