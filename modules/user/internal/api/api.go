package api

import (
	"context"

	"google.golang.org/grpc"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

type application interface {
}

type api struct {
	app application
}

func New(ctx context.Context, app application) *grpc.Server {
	srv := grpc.NewServer()

	userpb.RegisterUserAPIServer(srv, &api{app: app})
	
	return srv
}
