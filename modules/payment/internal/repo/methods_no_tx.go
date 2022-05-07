package repo

import (
	"context"
	"rest-on-grpc-gateway/modules/payment/internal/app"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"

	"github.com/jmoiron/sqlx"

	"github.com/shopspring/decimal"
)

// GetAccountsByUserID get accounts by user id in db without transaction.
func (r *Repo) GetAccountsByUserID(ctx context.Context, userID int) ([]domain.Account, error) {
	return getAccountsByUserID(ctx, r.DB, userID)
}

// GetUserAccountByAccountNumber get user account by account number in db without transaction.
func (r *Repo) GetUserAccountByAccountNumber(ctx context.Context, userID int, accountNumber string) (*domain.Account, error) {
	return getUserAccountByAccountNumber(ctx, r.DB, userID, accountNumber)
}

// GetPaymentHistoryByAccountNumber get payment history by account number in db without transaction.
func (r *Repo) GetPaymentHistoryByAccountNumber(ctx context.Context, userID int, accountNumber string, paging, filters filters.FilterContract) (
	_ []domain.Payment, total int, err error,
) {
	return getPaymentHistoryByAccountNumber(ctx, r.DB, userID, accountNumber, paging, filters)
}

// CreateOrUpdateAccount create or update account balance in db without transaction.
func (r *Repo) CreateOrUpdateAccount(ctx context.Context, userID int, accountNumber string, sum decimal.Decimal) error {
	return createOrUpdateAccount(ctx, r.DB, userID, accountNumber, sum)
}

// CreatePayment create payment in db without transaction.
func (r *Repo) CreatePayment(ctx context.Context, payment domain.Payment) error {
	return createPayment(ctx, r.DB, payment)
}

// DoTx wrapper over a transaction, used to apply the transaction at the business logic level.
func (r *Repo) DoTx(ctx context.Context, f func(repo app.Repo) error) error {
	return r.Repo.Tx(ctx, nil, func(tx *sqlx.Tx) error {
		return f(&WrapperTx{tx: tx})
	})
}
