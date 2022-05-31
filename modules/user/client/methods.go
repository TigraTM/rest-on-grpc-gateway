package client

import (
	"context"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

// GetByID search user by id.
func (c *Client) GetByID(ctx context.Context, userID int) (*User, error) {
	user, err := c.conn.UserByID(ctx, &userpb.UserByIDRequest{
		Id: int64(userID),
	})
	if err != nil {
		return nil, convertError(err)
	}

	return &User{
		ID:    int(user.User.Id),
		Name:  user.User.Name,
		Email: user.User.Email,
	}, nil
}
