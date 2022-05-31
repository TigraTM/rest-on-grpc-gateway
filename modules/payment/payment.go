// Package payment contains all logic for work with payment's.
package payment

import (
	"context"
	"fmt"
	"net"
	"strconv"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"

	"rest-on-grpc-gateway/modules/payment/adapters/apilayer"
	"rest-on-grpc-gateway/modules/payment/adapters/user"
	"rest-on-grpc-gateway/modules/payment/internal/api"
	"rest-on-grpc-gateway/modules/payment/internal/app"
	"rest-on-grpc-gateway/modules/payment/internal/config"
	"rest-on-grpc-gateway/modules/payment/internal/repo"
	"rest-on-grpc-gateway/pkg/serve"

	user_client "rest-on-grpc-gateway/modules/user/client"

	"go.uber.org/zap"

	paymentpb "rest-on-grpc-gateway/api/proto/payment/v1"
)

const appName = "payment"

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
		s.log.Fatal("failed repo.New: %+v \n", zap.Error(err))
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

	apiLayerClient := apilayer.New(s.cfg.Clients.APILayerAPIKey, s.cfg.Clients.APILayerBasePath)

	userClient, err := user_client.New(ctx, s.log, net.JoinHostPort(s.cfg.Clients.UserClientHost, strconv.Itoa(s.cfg.Clients.UserClientPort)))
	if err != nil {
		return fmt.Errorf("user_client.New: %w", err)
	}

	userSvc := user.New(userClient)

	appl := app.New(s.db, s.db, apiLayerClient, userSvc)
	grpcAPI := api.New(s.log, appl)

	gwCfg := serve.GateWayConfig{
		FS:             paymentpb.OpenAPI,
		GRPCServerPort: s.cfg.Transport.GRPCPort,
		Namespace:      "payment",
		GRPCGWPattern:  "/",
		OpenAPIPattern: "/openapi/",
		Register:       paymentpb.RegisterPaymentExternalAPIHandler,
	}

	err = serve.Start(
		ctx,
		serve.GRPC(s.log.With(zap.String("serve", "gRPC")), s.cfg.Transport.Host, s.cfg.Transport.GRPCPort, grpcAPI),
		serve.GRPCGateWay(s.log.With(zap.String("serve", "gRPC-Gateway")), s.cfg.Transport.Host, s.cfg.Transport.GRPCGWPort, gwCfg),
	)
	if err != nil {
		return fmt.Errorf("serve.Start: %w", err)
	}

	return nil
}
