package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

// Balance domain model of user balance.
type Balance struct {
	ID       int
	CreateAt time.Time
	UpdateAt time.Time
	Sum      decimal.Decimal
	Currency string
	UserID   int
}
