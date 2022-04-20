// Package user contains all logic for work with user's.
package user

import (
	"context"
	"fmt"
	"rest-on-grpc-gateway/modules/user/internal/api"
	"rest-on-grpc-gateway/modules/user/internal/app"
	"rest-on-grpc-gateway/modules/user/internal/config"
	"rest-on-grpc-gateway/modules/user/internal/repo"
	"rest-on-grpc-gateway/pkg/serve"

	"go.uber.org/zap"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

// Service ...
type Service struct {
	Log *zap.SugaredLogger
}

// Init service initialization.
func (*Service) Init() error {
	// add auto migration

	return nil
}

// RunServe start service.
func (s *Service) RunServe(parentCtx context.Context) error {
	cfg, err := config.LoadDevConfig()
	if err != nil {
		s.Log.Fatalf("couldn't get envConfig: %+v \n", err)
	}

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	db, err := repo.New(ctx, cfg.Database.DSN())
	if err != nil {
		s.Log.Fatalf("failed repo.New: %+v \n", err)
	}

	appl := app.New(db)
	grpcAPI := api.New(s.Log, appl)

	gwCfg := serve.GateWayConfig{
		FS:             userpb.OpenAPI,
		GRPCServerPort: cfg.Transport.GRPCPort,
		Namespace:      "user",
		GRPCGWPattern:  "/",
		OpenAPIPattern: "/openapi/",
		Register:       userpb.RegisterUserAPIHandler,
	}

	err = serve.Start(
		ctx,
		serve.GRPC(s.Log.With("serve", "gRPC"), cfg.Transport.Host, cfg.Transport.GRPCPort, grpcAPI),
		serve.GRPCGateWay(s.Log.With("serve", "gRPC-Gateway"), cfg.Transport.Host, cfg.Transport.GRPCGWPort, gwCfg),
	)
	if err != nil {
		return fmt.Errorf("serve.Start: %w", err)
	}

	return nil
}
