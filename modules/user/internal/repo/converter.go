package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"

	"rest-on-grpc-gateway/modules/user/internal/app"
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

const duplEmail = "users_email_key"

func convertErr(err error) error {
	var pqErr *pq.Error

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return app.ErrNotFound
	case errors.As(err, &pqErr):
		return constraint(pqErr)
	default:
		return err
	}
}

func constraint(pqErr *pq.Error) error {
	switch {
	case strings.HasSuffix(pqErr.Message, fmt.Sprintf("unique constraint \"%s\"", duplEmail)):
		return app.ErrEmailExist
	default:
		return pqErr
	}
}

func toDomain(user *User) *domain.User {
	return &domain.User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func toRepo(user *domain.User) *User {
	return &User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}
