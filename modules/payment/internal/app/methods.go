package app

import (
	"context"
	"fmt"

	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"
)

// CreatePayment create payment and update balance.
// Payment can be positive or negative. Payment create if:
// - recipient exist
// - balance value will not become negative after the payment
func (a *App) CreatePayment(ctx context.Context, payment domain.Payment) error {
	// TODO: add check exist user
	// TODO: add create payment history
	return a.repo.CreateOrUpdatePayment(ctx, payment)
}

// GetBalanceByUserID get balance by userID.
func (a *App) GetBalanceByUserID(ctx context.Context, userID int, currency string) (*domain.Balance, error) {
	// TODO: add convert in currency
	return a.repo.GetBalanceByUserID(ctx, userID)
}

// TransferBetweenUsers transferring money between users.
func (a *App) TransferBetweenUsers(ctx context.Context, transfer domain.Transfer) (*domain.Transfer, error) {
	senderBalance, err := a.repo.GetBalanceByUserID(ctx, transfer.SenderID)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetBalanceByUserID: %w", err)
	}

	if senderBalance.Sum.Sub(transfer.Sum).IsNegative() {
		return nil, ErrNotEnoughMoney
	}
	// TODO: add check exist recipient
	// TODO: add payment history
	// TODO: add transactions
	if err = a.repo.SubBalanceByUserID(ctx, transfer.SenderID, transfer.Sum); err != nil {
		return nil, fmt.Errorf("a.repo.SubBalanceByUserID: %w", err)
	}

	if err = a.repo.AddBalanceByUserID(ctx, transfer.RecipientID, transfer.Sum); err != nil {
		return nil, fmt.Errorf("a.repo.AddBalanceByUserID: %w", err)
	}

	return nil, nil
}

// GetPaymentHistoryByUserID get payment history by userID.
func (a *App) GetPaymentHistoryByUserID(ctx context.Context, userID int, paging, filter filters.FilterContract) ([]domain.Payment, int, error) {
	return a.repo.GetPaymentHistoryByUserID(ctx, userID, paging, filter)
}
