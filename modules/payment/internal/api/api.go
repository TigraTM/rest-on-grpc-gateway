// Package api contains payment handlers for work gRPC-gateway.
package api

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	paymentpb "rest-on-grpc-gateway/api/proto/payment/v1"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"
	"rest-on-grpc-gateway/pkg/grpc_helper"
)

var (
	errUncorrectedFilter = errors.New("uncorrected filter")
	errUncorrectedPaging = errors.New("uncorrected paging")
)

// application for easy test.
type application interface {
	CreatePayment(ctx context.Context, payment domain.Payment) error
	GetBalanceByUserID(ctx context.Context, userID int, currency string) (*domain.Balance, error)
	TransferBetweenUsers(ctx context.Context, transfer domain.Transfer) (*domain.Transfer, error)
	GetPaymentHistoryByUserID(ctx context.Context, userID int, paging, filter filters.FilterContract) ([]domain.Payment, int, error)
}

// api structure api.
type api struct {
	app application
}

// New build and return new grpc.Server.
func New(log *zap.SugaredLogger, app application) *grpc.Server {
	srv := grpc_helper.NewServer(log, apiError, []grpc.UnaryServerInterceptor{})

	paymentpb.RegisterPaymentAPIServer(srv, &api{app: app})

	return srv
}

// apiError convert err in status code.
func apiError(err error) *status.Status {
	if err == nil {
		return nil
	}

	code := codes.Internal
	switch {
	case errors.Is(err, errUncorrectedPaging):
		code = codes.InvalidArgument
	case errors.Is(err, errUncorrectedFilter):
		code = codes.InvalidArgument
	case errors.Is(err, context.DeadlineExceeded):
		code = codes.DeadlineExceeded
	case errors.Is(err, context.Canceled):
		code = codes.Canceled
	}

	return status.New(code, err.Error())
}
