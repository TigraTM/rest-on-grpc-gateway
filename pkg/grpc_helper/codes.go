// nolint:stylecheck // name is more readable
package grpc_helper

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// GRPCCodesConverterHandler is a function that convert your error to gRPC codes.
type GRPCCodesConverterHandler = func(error) *status.Status

// UnaryConvertCodesServerInterceptor returns a new unary server interceptor that converting returns error.
func UnaryConvertCodesServerInterceptor(converter GRPCCodesConverterHandler) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, converter(err).Err()
		}

		return resp, err
	}
}
