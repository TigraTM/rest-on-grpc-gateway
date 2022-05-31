package api

import (
	userpb "rest-on-grpc-gateway/api/proto/user/v1"
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

func toPBUser(user *domain.User) *userpb.User {
	return &userpb.User{
		Id:    int64(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
}
