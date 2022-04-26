package api

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	paymentpb "rest-on-grpc-gateway/api/proto/payment/v1"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
)

// CreatePayment implements userpb.UserAPIServer.
func (a *api) CreatePayment(ctx context.Context, in *paymentpb.CreatePaymentRequest) (*paymentpb.CreatePaymentResponse, error) {
	sum, err := decimal.NewFromString(in.Sum.Value)
	if err != nil {
		return nil, fmt.Errorf("decimal.NewFromString: %w", err)
	}

	newPayment := domain.Payment{
		Sum:         sum,
		CompanyName: in.CompanyName,
		Category:    in.Category,
		UserID:      int(in.UserId),
	}

	err = a.app.CreatePayment(ctx, newPayment)
	switch {
	default:
		return nil, fmt.Errorf("a.app.CreateUser: %w", err)
	}
}

// GetBalanceByUserID implements userpb.UserAPIServer.
func (a *api) GetBalanceByUserID(ctx context.Context, in *paymentpb.GetBalanceByUserIDRequest) (*paymentpb.GetBalanceByUserIDResponse, error) {
	return nil, nil
}

// TransferBetweenUsers implements userpb.UserAPIServer.
func (a *api) TransferBetweenUsers(ctx context.Context, in *paymentpb.TransferBetweenUsersRequest) (*paymentpb.TransferBetweenUsersResponse, error) {
	return nil, nil
}

// GetPaymentsHistoryByUserID implements userpb.UserAPIServer.
func (a *api) GetPaymentsHistoryByUserID(ctx context.Context, in *paymentpb.GetPaymentsHistoryByUserIDRequest) (*paymentpb.GetPaymentsHistoryByUserIDResponse, error) {
	return nil, nil
}
