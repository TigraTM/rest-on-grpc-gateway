package repo

import (
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

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
