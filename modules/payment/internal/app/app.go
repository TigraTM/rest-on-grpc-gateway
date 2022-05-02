// Package app contains all business logic for work with payment.
package app

import (
	"context"
	"database/sql"
	"errors"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"

	"github.com/jmoiron/sqlx"

	"github.com/shopspring/decimal"
)

// Errors.
var (
	ErrNotEnoughMoney           = errors.New("not enough money")
	ErrNotFound                 = errors.New("not found")
	ErrSameAccountNumber        = errors.New("same account number")
	ErrTransferAmountNotCorrect = errors.New("transfer amount is not correct")
)

//go:generate mockgen -source=app.go -destination mock.app.contracts_test.go -package app_test

type (
	// Repo interface for payment database.
	Repo interface {
		GetUserAccountByAccountNumber(ctx context.Context, userID int, accountNumber string) (*domain.Account, error)
		GetAccountsByUserID(ctx context.Context, userID int) ([]domain.Account, error)
		GetPaymentHistoryByAccountNumber(ctx context.Context, accountNumber string, paging, filters filters.FilterContract) (_ []domain.Payment, total int, err error)
		CreateOrUpdateAccount(ctx context.Context, userID int, accountNumber string, sum decimal.Decimal) error
		CreatePayment(ctx context.Context, payment domain.Payment) error
		Tx(ctx context.Context, opts *sql.TxOptions, f func(*sqlx.Tx) error) (err error)
	}
	// ExchangeClient interface to convert balance to other currencies.
	ExchangeClient interface{}
)

type App struct {
	repo     Repo
	exchange ExchangeClient
}

// New build and return new App.
func New(repo Repo, exchange ExchangeClient) *App {
	return &App{
		repo:     repo,
		exchange: exchange,
	}
}
