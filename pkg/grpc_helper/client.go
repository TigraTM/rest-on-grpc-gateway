package grpc_helper

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Dial creates a gRPC client connection to the given target.
func Dial(ctx context.Context,
	log *zap.Logger,
	addr string,
	extraUnary []grpc.UnaryClientInterceptor,
	extraDialOption []grpc.DialOption,
) (*grpc.ClientConn, error) {
	dialOptions := append([]grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}, extraDialOption...)

	unaryInterceptor := append([]grpc.UnaryClientInterceptor{
		grpc_zap.UnaryClientInterceptor(log),
		grpc_zap.PayloadUnaryClientInterceptor(log, alwaysLoggingDeciderClient),
		grpc_validator.UnaryClientInterceptor(),
	}, extraUnary...)

	dialOptions = append(dialOptions,
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(unaryInterceptor...),
		),
	)

	return grpc.DialContext(ctx, addr, dialOptions...)
}
