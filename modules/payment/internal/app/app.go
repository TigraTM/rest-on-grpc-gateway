// Package app contains all business logic for work with payment.
package app

import (
	"context"
	"errors"

	"github.com/shopspring/decimal"

	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"
)

// Errors.
var (
	ErrNotEnoughMoney = errors.New("not enough money")
)

//go:generate mockgen -source=app.go -destination mock.app.contracts_test.go -package app_test

type (
	// Repo interface for payment database.
	Repo interface {
		CreateOrUpdatePayment(ctx context.Context, payment domain.Payment) error
		GetBalanceByUserID(ctx context.Context, userID int) (*domain.Balance, error)
		GetPaymentHistoryByUserID(ctx context.Context, userID int, paging, filter filters.FilterContract) ([]domain.Payment, int, error)
		AddBalanceByUserID(ctx context.Context, userID int, sum decimal.Decimal) error
		SubBalanceByUserID(ctx context.Context, userID int, sum decimal.Decimal) error
	}
	// ExchangeClient interface to convert balance to other currencies.
	ExchangeClient interface {
	}
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
