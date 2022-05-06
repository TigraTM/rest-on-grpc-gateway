// Package grpc_helper contains auxiliary functions for gRPC.
package grpc_helper

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewServer(log *zap.SugaredLogger, converter GRPCCodesConverterHandler, addUnary []grpc.UnaryServerInterceptor) *grpc.Server {
	grpc_zap.ReplaceGrpcLoggerV2(log.Desugar())

	unaryInterceptor := append([]grpc.UnaryServerInterceptor{
		grpc_zap.UnaryServerInterceptor(log.Desugar()),
		grpc_recovery.UnaryServerInterceptor(),
		grpc_validator.UnaryServerInterceptor(),
		UnaryConvertCodesServerInterceptor(converter),
	}, addUnary...)

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(unaryInterceptor...),
		),
	)

	return srv
}
