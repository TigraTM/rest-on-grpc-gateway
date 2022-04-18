package repo

import (
	"context"

	"rest-on-grpc-gateway/modules/user/internal/domain"
)

// CreateUser create user in db.
func (r *Repo) CreateUser(ctx context.Context, newUser *domain.User) (*domain.User, error) {
	return nil, nil
}

// GetUserByID search user in db.
func (r *Repo) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return nil, nil
}

// UpdateUserByID update user in db by id.
func (r *Repo) UpdateUserByID(ctx context.Context, id int, name, email string) (*domain.User, error) {
	return nil, nil
}

// UpdateUserPasswordByID update user password in db by id.
func (r *Repo) UpdateUserPasswordByID(ctx context.Context, id int, oldPass, newPass string) error {
	return nil
}

// DeleteUserByID delete user in db by id.
func (r *Repo) DeleteUserByID(ctx context.Context, id int) error {
	return nil
}
