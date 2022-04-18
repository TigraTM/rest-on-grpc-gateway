package api

import (
	"context"
	"fmt"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
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
	if err != nil {
		return nil, fmt.Errorf("a.app.GetUserByID: %w", err)
	}

	return &userpb.GetUserByIDResponse{
		Id:    int64(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// UpdateUserByID implements userpb.UserAPIServer.
func (a *api) UpdateUserByID(ctx context.Context, in *userpb.UpdateUserByIDRequest) (*userpb.UpdateUserByIDResponse, error) {
	user, err := a.app.UpdateUserByID(ctx, int(in.Id), in.Name, in.Email)
	if err != nil {
		return nil, fmt.Errorf("a.app.UpdateUserByID: %w", err)
	}

	return &userpb.UpdateUserByIDResponse{
		Id:    int64(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// UpdateUserPasswordByID implements userpb.UserAPIServer.
func (a *api) UpdateUserPasswordByID(ctx context.Context, in *userpb.UpdateUserPasswordByIDRequest) (*userpb.UpdateUserPasswordByIDResponse, error) {
	err := a.app.UpdateUserPasswordByID(ctx, int(in.Id), in.OldPassword, in.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("a.app.UpdateUserPasswordByID: %w", err)
	}

	return &userpb.UpdateUserPasswordByIDResponse{}, nil
}

// DeleteUserByID implements userpb.UserAPIServer.
func (a *api) DeleteUserByID(ctx context.Context, in *userpb.DeleteUserByIDRequest) (*userpb.DeleteUserByIDResponse, error) {
	err := a.app.DeleteUserByID(ctx, int(in.Id))
	if err != nil {
		return nil, fmt.Errorf("a.app.DeleteUserByIDL: %w", err)
	}

	return &userpb.DeleteUserByIDResponse{}, nil
}
