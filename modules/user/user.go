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

	// Driver.
	_ "github.com/golang-migrate/migrate/v4/source/file"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

// Service ...
type Service struct {
	ctx context.Context
	cfg *config.Config
	log *zap.SugaredLogger
	db  *repo.Repo
}

// Init service initialization.
func (s *Service) Init(ctx context.Context, log *zap.SugaredLogger) (err error) {
	s.log = log
	s.ctx = ctx

	s.cfg, err = config.LoadDevConfig()
	if err != nil {
		s.log.Fatalf("couldn't get envConfig: %+v \n", err)
	}

	s.db, err = repo.New(s.ctx, &s.cfg.Database)
	if err != nil {
		s.log.Fatalf("failed repo.New: %+v \n", err)
	}

	return nil
}

// RunServe start service.
func (s *Service) RunServe() error {
	appl := app.New(s.db)
	grpcAPI := api.New(s.log, appl)

	gwCfg := serve.GateWayConfig{
		FS:             userpb.OpenAPI,
		GRPCServerPort: s.cfg.Transport.GRPCPort,
		Namespace:      "user",
		GRPCGWPattern:  "/",
		OpenAPIPattern: "/openapi/",
		Register:       userpb.RegisterUserAPIHandler,
	}

	err := serve.Start(
		s.ctx,
		serve.GRPC(s.log.With("serve", "gRPC"), s.cfg.Transport.Host, s.cfg.Transport.GRPCPort, grpcAPI),
		serve.GRPCGateWay(s.log.With("serve", "gRPC-Gateway"), s.cfg.Transport.Host, s.cfg.Transport.GRPCGWPort, gwCfg),
	)
	if err != nil {
		return fmt.Errorf("serve.Start: %w", err)
	}

	return nil
}
