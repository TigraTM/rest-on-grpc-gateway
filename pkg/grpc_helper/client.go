package grpc_helper

import (
	"context"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Dial creates a gRPC client connection to the given target.
func Dial(ctx context.Context,
	addr string,
	extraUnary []grpc.UnaryClientInterceptor,
	extraDialOption []grpc.DialOption,
) (*grpc.ClientConn, error) {
	log := ctxzap.Extract(ctx)

	dialOptions := append([]grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}, extraDialOption...)

	unaryInterceptor := append([]grpc.UnaryClientInterceptor{
		grpc_zap.UnaryClientInterceptor(log),
		grpc_validator.UnaryClientInterceptor(),
	}, extraUnary...)

	dialOptions = append(dialOptions,
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(unaryInterceptor...),
		),
	)

	return grpc.DialContext(ctx, addr, dialOptions...)
}
