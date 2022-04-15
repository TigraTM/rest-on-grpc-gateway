package api

import (
	"context"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

func (*api) CreateUser(_ context.Context, _ *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	return &userpb.CreateUserResponse{Id: "hello Andrey"}, nil
}
