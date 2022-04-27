// Package domain contains domain models.
package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

// Payment domain model payment.
type Payment struct {
	ID            int
	CreateAt      time.Time
	AccountNumber string
	Sum           decimal.Decimal
	CompanyName   string
	Category      string
	AccountID     int
}
