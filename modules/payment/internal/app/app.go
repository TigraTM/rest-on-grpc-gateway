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
	ErrNotEnoughMoney = errors.New("not enough money")
	ErrNotFound       = errors.New("not found")
)

//go:generate mockgen -source=app.go -destination mock.app.contracts_test.go -package app_test

type (
	// Repo interface for payment database.
	Repo interface {
		GetAccountByID(ctx context.Context, accountID int) (*domain.Account, error)
		GetAccountsByUserID(ctx context.Context, userID int) ([]domain.Account, error)
		GetPaymentHistoryByAccountID(ctx context.Context, accountID int, paging, filters filters.FilterContract) ([]domain.Payment, int, error)
		CreateOrUpdateAccount(ctx context.Context, userID int, accountNumber string, sum decimal.Decimal) error
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
