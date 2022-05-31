// Package apiExternal contains user handlers for work gRPC-gateway.
package api

import (
	"context"
	"errors"
	"rest-on-grpc-gateway/modules/user/internal/domain"
	"rest-on-grpc-gateway/pkg/grpc_helper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

// Errors.
var (
	errUserNotFound    = errors.New("user not found")
	errInvalidPassword = errors.New("invalid password")
	errMustDifferent   = errors.New("value of the existing password and the new password must be different")
	errEmailExist      = errors.New("user with this email already exists")
)

//go:generate mockgen -source=api.go -destination mock.application.contracts_test.go -package api_test

// application interface business logic.
type application interface {
	CreateUser(ctx context.Context, name, email, password string) (*domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	UpdateUserByID(ctx context.Context, userID int, name, email string) (*domain.User, error)
	UpdateUserPasswordByID(ctx context.Context, userID int, oldPass, newPass string) error
	DeleteUserByID(ctx context.Context, userID int) error
}

// apiExternal structure apiExternal.
type apiExternal struct {
	app application
}

// apiInternal structure internal.
type apiInternal struct {
	app application
}

// NewExternal build and return new grpc.Server.
func NewExternal(log *zap.Logger, app application) *grpc.Server {
	srv := grpc_helper.NewServer(log, apiError, []grpc.UnaryServerInterceptor{})

	userpb.RegisterUserExternalAPIServer(srv, &apiExternal{app: app})

	return srv
}

// NewInternal build and return new grpc.Server.
func NewInternal(log *zap.Logger, app application) *grpc.Server {
	srv := grpc_helper.NewServer(log, apiError, []grpc.UnaryServerInterceptor{})

	userpb.RegisterUserInternalAPIServer(srv, &apiInternal{app: app})

	return srv
}

// apiError convert err in status code.
func apiError(err error) *status.Status {
	if err == nil {
		return nil
	}

	code := codes.Internal
	switch {
	case errors.Is(err, errUserNotFound):
		code = codes.NotFound
	case errors.Is(err, errInvalidPassword):
		code = codes.InvalidArgument
	case errors.Is(err, errMustDifferent):
		code = codes.InvalidArgument
	case errors.Is(err, errEmailExist):
		code = codes.InvalidArgument
	case errors.Is(err, context.DeadlineExceeded):
		code = codes.DeadlineExceeded
	case errors.Is(err, context.Canceled):
		code = codes.Canceled
	}

	return status.New(code, err.Error())
}
