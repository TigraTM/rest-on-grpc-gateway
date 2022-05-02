// Package api contains payment handlers for work gRPC-gateway.
package api

import (
	"context"
	"errors"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"
	"rest-on-grpc-gateway/pkg/grpc_helper"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	paymentpb "rest-on-grpc-gateway/api/proto/payment/v1"
)

var (
	errUncorrectedSort          = errors.New("uncorrected sort")
	errUncorrectedPaging        = errors.New("uncorrected paging")
	errNotFound                 = errors.New("not found")
	errNotEnoughMoney           = errors.New("not enough money")
	errSameAccountNumber        = errors.New("sender's and receiver's accounts are the same")
	errTransferAmountNotCorrect = errors.New("amount of the transfer must not be negative")
)

// application for easy test.
type application interface {
	CreatePayment(ctx context.Context, userID int, payment domain.Payment) (err error)
	GetAccountByAccountNumber(ctx context.Context, userID int, accountNumber, currency string) (*domain.Account, error)
	TransferBetweenUsers(ctx context.Context, transfer domain.Transfer) (*domain.Transfer, error)
	GetPaymentHistoryByAccountID(ctx context.Context, userID int, accountNumber string, paging, filter filters.FilterContract) ([]domain.Payment, int, error)
	GetAccountsByUserID(ctx context.Context, userID int) ([]domain.Account, error)
}

// api structure api.
type api struct {
	app application
}

// New build and return new grpc.Server.
func New(log *zap.SugaredLogger, app application) *grpc.Server {
	srv := grpc_helper.NewServer(log, apiError, []grpc.UnaryServerInterceptor{})

	paymentpb.RegisterPaymentExternalAPIServer(srv, &api{app: app})

	return srv
}

// apiError convert err in status code.
func apiError(err error) *status.Status {
	if err == nil {
		return nil
	}

	code := codes.Internal
	switch {
	case errors.Is(err, errSameAccountNumber):
		code = codes.InvalidArgument
	case errors.Is(err, errNotFound):
		code = codes.NotFound
	case errors.Is(err, errNotEnoughMoney):
		code = codes.InvalidArgument
	case errors.Is(err, errTransferAmountNotCorrect):
		code = codes.InvalidArgument
	case errors.Is(err, errUncorrectedPaging):
		code = codes.InvalidArgument
	case errors.Is(err, errUncorrectedSort):
		code = codes.InvalidArgument
	case errors.Is(err, context.DeadlineExceeded):
		code = codes.DeadlineExceeded
	case errors.Is(err, context.Canceled):
		code = codes.Canceled
	}

	return status.New(code, err.Error())
}
