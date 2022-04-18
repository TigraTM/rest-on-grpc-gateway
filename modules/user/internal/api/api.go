// Package api contains handlers for work gRPC-gateway.
package api

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"rest-on-grpc-gateway/modules/user/internal/app"
	"rest-on-grpc-gateway/modules/user/internal/domain"
	"rest-on-grpc-gateway/pkg/grpc_helper"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

// application interface business logic.
type application interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	UpdateUserByID(ctx context.Context, userID int, name, email string) (*domain.User, error)
	UpdateUserPasswordByID(ctx context.Context, userID int, oldPass, newPass string) error
	DeleteUserByID(ctx context.Context, userID int) error
}

// api structure api,.
type api struct {
	app application
}

// New build and return new grpc.Server.
func New(log *zap.SugaredLogger, app application) *grpc.Server {
	srv := grpc_helper.NewServer(log, apiError)

	userpb.RegisterUserAPIServer(srv, &api{app: app})

	return srv
}

func apiError(err error) *status.Status {
	if err == nil {
		return nil
	}

	code := codes.Internal
	switch {
	case errors.Is(err, app.ErrNotFound):
		code = codes.NotFound
	case errors.Is(err, app.ErrInvalidPassword):
		code = codes.InvalidArgument
	case errors.Is(err, context.DeadlineExceeded):
		code = codes.DeadlineExceeded
	case errors.Is(err, context.Canceled):
		code = codes.Canceled
	}

	return status.New(code, err.Error())
}
