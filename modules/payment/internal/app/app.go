// Package app contains all business logic for work with payment.
package app

import (
	"context"
	"errors"

	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"

	"github.com/shopspring/decimal"
)

// Errors.
var (
	ErrNotEnoughMoney           = errors.New("not enough money")
	ErrNotFound                 = errors.New("not found")
	ErrSameAccountNumber        = errors.New("same account number")
	ErrTransferAmountNotCorrect = errors.New("transfer amount is not correct")
	ErrExchangeClient           = errors.New("exchange client")
)

//go:generate mockgen -source=app.go -destination mock.app.contracts_test.go -package app_test

type (
	// Repo interface for payment database.
	Repo interface {
		GetUserAccountByAccountNumber(ctx context.Context, userID int, accountNumber string) (*domain.Account, error)
		GetAccountsByUserID(ctx context.Context, userID int) ([]domain.Account, error)
		GetPaymentHistoryByAccountNumber(ctx context.Context, userID int, accountNumber string, paging, filters filters.FilterContract) (_ []domain.Payment, total int, err error)
		CreateOrUpdateAccount(ctx context.Context, userID int, accountNumber string, sum decimal.Decimal) error
		CreatePayment(ctx context.Context, payment domain.Payment) error
	}
	// TxRepo interface for payment database with transaction.
	TxRepo interface {
		DoTx(ctx context.Context, f func(repo Repo) error) error
		Repo
	}
	// ExchangeClient interface to convert balance to other currencies.
	ExchangeClient interface {
		ConvertAmount(ctx context.Context, fromCurrency, toCurrency string, amount decimal.Decimal) (decimal.Decimal, error)
		GetSymbols(ctx context.Context) (map[string]string, error)
	}
	// UserClient interface for work with user service.
	UserClient interface {
		ExistUserByID(ctx context.Context, userID int) error
	}
)

// App domain structure business logic.
type App struct {
	repo     Repo
	txRepo   TxRepo
	exchange ExchangeClient
	user     UserClient
}

// New build and return new App.
func New(repo Repo, txRepo TxRepo, exchange ExchangeClient, user UserClient) *App {
	return &App{
		repo:     repo,
		txRepo:   txRepo,
		exchange: exchange,
		user:     user,
	}
}
