package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

// Account domain model of user account.
type Account struct {
	ID            int
	CreateAt      time.Time
	UpdateAt      time.Time
	AccountNumber string
	Balance       decimal.Decimal
	Currency      string
	UserID        int
}
