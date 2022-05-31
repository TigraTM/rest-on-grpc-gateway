// Package user contains all logic for work with user's.
package user

import (
	"context"
	"fmt"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"

	"rest-on-grpc-gateway/modules/user/internal/api"
	"rest-on-grpc-gateway/modules/user/internal/app"
	"rest-on-grpc-gateway/modules/user/internal/config"
	"rest-on-grpc-gateway/modules/user/internal/repo"
	"rest-on-grpc-gateway/pkg/password"
	"rest-on-grpc-gateway/pkg/serve"

	"go.uber.org/zap"

	// Driver.
	_ "github.com/golang-migrate/migrate/v4/source/file"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
)

const appName = "user"

// Service ...
type Service struct {
	cfg *config.Config
	log *zap.Logger
	db  *repo.Repo
}

// Name return service name.
func (*Service) Name() string {
	return appName
}

// Init service initialization.
func (s *Service) Init(ctx context.Context, log *zap.Logger) (err error) {
	s.log = log

	s.cfg, err = config.LoadDevConfig()
	if err != nil {
		s.log.Fatal("couldn't get envConfig: %+v \n", zap.Error(err))
	}

	s.db, err = repo.New(ctx, &s.cfg.Database, appName)
	if err != nil {
		s.log.Fatal("failed repo.NewExternal: %+v \n", zap.Error(err))
	}
	//defer func() {  // TODO: Check db close
	//	err := s.db.DB.Close()
	//	if err != nil {
	//		log.Error("close database connection", zap.Error(err))
	//	}
	//}()

	return nil
}

// RunServe start service.
func (s *Service) RunServe(ctx context.Context) error {
	grpc_zap.ReplaceGrpcLoggerV2(s.log)

	hashSvc := password.New()

	appl := app.New(s.db, hashSvc)
	grpcExternalAPI := api.NewExternal(s.log, appl)
	grpcInternalAPI := api.NewInternal(s.log, appl)

	gwCfg := serve.GateWayConfig{
		FS:             userpb.OpenAPI,
		GRPCServerPort: s.cfg.Transport.GRPCExternalPort,
		Namespace:      "user",
		GRPCGWPattern:  "/",
		OpenAPIPattern: "/openapi/",
		Register:       userpb.RegisterUserExternalAPIHandler,
	}

	err := serve.Start(
		ctx,
		serve.GRPC(s.log.With(zap.String("serve", "gRPC-external")), s.cfg.Transport.Host, s.cfg.Transport.GRPCExternalPort, grpcExternalAPI),
		serve.GRPC(s.log.With(zap.String("serve", "gRPC-internal")), s.cfg.Transport.Host, s.cfg.Transport.GRPCInternalPort, grpcInternalAPI),
		serve.GRPCGateWay(s.log.With(zap.String("serve", "gRPC-Gateway")), s.cfg.Transport.Host, s.cfg.Transport.GRPCGWPort, gwCfg),
	)
	if err != nil {
		return fmt.Errorf("serve.Start: %w", err)
	}

	return nil
}
