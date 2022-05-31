package api

import (
	"context"
	"errors"
	"fmt"
	"rest-on-grpc-gateway/modules/user/internal/app"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

// CreateUser implements userpb.UserAPIServer.
func (a *apiExternal) CreateUser(ctx context.Context, in *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user, err := a.app.CreateUser(ctx, in.Name, in.Email, in.Password)
	switch {
	case err == nil:
		return &userpb.CreateUserResponse{
			User: toPBUser(user),
		}, nil
	case errors.Is(err, app.ErrEmailExist):
		return nil, errEmailExist
	default:
		return nil, fmt.Errorf("a.app.CreateUser: %w", err)
	}
}

// GetUserByID implements userpb.UserAPIServer.
func (a *apiExternal) GetUserByID(ctx context.Context, in *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	user, err := a.app.GetUserByID(ctx, int(in.Id))
	switch {
	case err == nil:
		return &userpb.GetUserByIDResponse{
			User: toPBUser(user),
		}, nil
	case errors.Is(err, app.ErrNotFound):
		return nil, errUserNotFound
	default:
		return nil, fmt.Errorf("a.app.GetUserByID: %w", err)
	}
}

// UpdateUserByID implements userpb.UserAPIServer.
func (a *apiExternal) UpdateUserByID(ctx context.Context, in *userpb.UpdateUserByIDRequest) (*userpb.UpdateUserByIDResponse, error) {
	user, err := a.app.UpdateUserByID(ctx, int(in.Id), in.Name, in.Email)
	switch {
	case err == nil:
		return &userpb.UpdateUserByIDResponse{
			User: toPBUser(user),
		}, nil
	case errors.Is(err, app.ErrNotFound):
		return nil, errUserNotFound
	default:
		return nil, fmt.Errorf("a.app.UpdateUserByID: %w", err)
	}
}

// UpdateUserPasswordByID implements userpb.UserAPIServer.
func (a *apiExternal) UpdateUserPasswordByID(ctx context.Context, in *userpb.UpdateUserPasswordByIDRequest) (*userpb.UpdateUserPasswordByIDResponse, error) {
	err := a.app.UpdateUserPasswordByID(ctx, int(in.Id), in.OldPassword, in.NewPassword)
	switch {
	case err == nil:
		return &userpb.UpdateUserPasswordByIDResponse{}, nil
	case errors.Is(err, app.ErrNotFound):
		return nil, errUserNotFound
	case errors.Is(err, app.ErrInvalidPassword):
		return nil, errInvalidPassword
	case errors.Is(err, app.ErrMustDifferent):
		return nil, errMustDifferent
	default:
		return nil, fmt.Errorf("a.app.UpdateUserPasswordByID: %w", err)
	}
}

// DeleteUserByID implements userpb.UserAPIServer.
func (a *apiExternal) DeleteUserByID(ctx context.Context, in *userpb.DeleteUserByIDRequest) (*userpb.DeleteUserByIDResponse, error) {
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

func (a *apiInternal) UserByID(ctx context.Context, in *userpb.UserByIDRequest) (*userpb.UserByIDResponse, error) {
	user, err := a.app.GetUserByID(ctx, int(in.Id))
	switch {
	case err == nil:
		return &userpb.UserByIDResponse{
			User: toPBUser(user),
		}, nil
	case errors.Is(err, app.ErrNotFound):
		return nil, errUserNotFound
	default:
		return nil, fmt.Errorf("a.app.GetUserByID: %w", err)
	}
}
