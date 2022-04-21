package api

import (
	userpb "rest-on-grpc-gateway/api/proto/user/v1"
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

func toDomain(req *userpb.CreateUserRequest) *domain.User {
	return &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
