package repo

import (
	"context"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"

	"github.com/shopspring/decimal"
)

// GetAccountsByUserID get accounts by user id in db with transaction.
func (w *WrapperTx) GetAccountsByUserID(ctx context.Context, userID int) ([]domain.Account, error) {
	return getAccountsByUserID(ctx, w.tx, userID)
}

// GetUserAccountByAccountNumber get user account by account number in db with transaction.
func (w *WrapperTx) GetUserAccountByAccountNumber(ctx context.Context, userID int, accountNumber string) (*domain.Account, error) {
	return getUserAccountByAccountNumber(ctx, w.tx, userID, accountNumber)
}

// GetPaymentHistoryByAccountNumber get payment history by account number in db with transaction.
func (w *WrapperTx) GetPaymentHistoryByAccountNumber(ctx context.Context, accountNumber string, paging, filters filters.FilterContract) (
	_ []domain.Payment, total int, err error,
) {
	return getPaymentHistoryByAccountNumber(ctx, w.tx, accountNumber, paging, filters)
}

// CreateOrUpdateAccount create or update account balance in db with transaction.
func (w *WrapperTx) CreateOrUpdateAccount(ctx context.Context, userID int, accountNumber string, sum decimal.Decimal) error {
	return createOrUpdateAccount(ctx, w.tx, userID, accountNumber, sum)
}

// CreatePayment create payment in db with transaction.
func (w *WrapperTx) CreatePayment(ctx context.Context, payment domain.Payment) error {
	return createPayment(ctx, w.tx, payment)
}
