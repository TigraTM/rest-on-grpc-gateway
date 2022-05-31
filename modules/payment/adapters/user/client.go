// Package user needed for internal interaction with the service user.
package user

import (
	"context"
	"rest-on-grpc-gateway/modules/user/client"
)

//go:generate mockgen -source=client.go -destination mock.client.contracts_test.go -package user_test

// for test.
type userClient interface {
	GetByID(ctx context.Context, userID int) (*client.User, error)
}

// Client wrapper for user service.
type Client struct {
	user userClient
}

// New build and return new user `Client.
func New(user userClient) *Client {
	return &Client{
		user: user,
	}
}
