// Package client provide to internal method of service user.
package client

import (
	"context"
	"errors"
	"fmt"

	"rest-on-grpc-gateway/pkg/grpc_helper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

// Errors.
var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInternal        = errors.New("internal error")
)

// Client to user service.
type Client struct {
	conn userpb.UserInternalAPIClient
}

// New build and returns new client to service user.
func New(ctx context.Context, addr string) (*Client, error) {
	conn, err := grpc_helper.Dial(ctx, addr, []grpc.UnaryClientInterceptor{}, []grpc.DialOption{})
	if err != nil {
		return nil, fmt.Errorf("grpc_helper.Dial: %w", err)
	}

	return &Client{conn: userpb.NewUserInternalAPIClient(conn)}, nil
}

func convertError(err error) error {
	switch {
	case status.Code(err) == codes.NotFound:
		return fmt.Errorf("%w: %s", ErrNotFound, err)
	case status.Code(err) == codes.DeadlineExceeded:
		return fmt.Errorf("%w: %s", context.DeadlineExceeded, err)
	case status.Code(err) == codes.Canceled:
		return fmt.Errorf("%w: %s", context.Canceled, err)
	case status.Code(err) == codes.InvalidArgument:
		return fmt.Errorf("%w: %s", ErrInvalidArgument, err)
	default:
		st, ok := status.FromError(err)
		if !ok {
			return err
		}

		return fmt.Errorf("%w: %s", ErrInternal, st.Message())
	}
}
