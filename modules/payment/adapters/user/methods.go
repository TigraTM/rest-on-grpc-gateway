package user

import (
	"context"
	"errors"
	"rest-on-grpc-gateway/modules/payment/internal/app"
	"rest-on-grpc-gateway/modules/user/client"
)

// ExistUserByID check exist user in user service.
func (c *Client) ExistUserByID(ctx context.Context, userID int) error {
	_, err := c.user.GetByID(ctx, userID)
	switch {
	case err == nil:
		return nil
	case errors.Is(err, client.ErrNotFound):
		return app.ErrNotFound
	default:
		return err
	}
}
