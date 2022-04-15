package user

import (
	"context"
	"fmt"
	logStd "log"

	"go.uber.org/zap"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
	"rest-on-grpc-gateway/modules/user/internal/api"
	"rest-on-grpc-gateway/modules/user/internal/config"
	"rest-on-grpc-gateway/pkg/serve"
)

type Service struct {
	Log *zap.SugaredLogger
}

func (s *Service) Init() error {
	// add auto migration

	return nil
}

func (s *Service) RunServe(parentCtx context.Context) error {
	cfg, err := config.LoadDevConfig()
	if err != nil {
		logStd.Printf("couldn't get envConfig: %+v \n", err)
	}

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	grpcAPI := api.New(ctx, nil)

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
