package repo

import (
	"database/sql"
	"errors"

	"rest-on-grpc-gateway/modules/user/internal/app"
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

func convertErr(err error) error {
	//var pqErr *pq.Error

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return app.ErrNotFound
	//case errors.As(err, &pqErr):
	//	return constraint(pqErr)
	default:
		return err
	}
}

func toDomain(user *User) *domain.User {
	return &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func toRepo(user *domain.User) *User {
	return &User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
