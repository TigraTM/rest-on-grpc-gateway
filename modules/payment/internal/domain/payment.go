// Package domain contains domain models.
package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

// PaymentCategory type payment category.
type PaymentCategory string

const (
	// PaymentCategoryTransfer transferring money between users (+/- money).
	PaymentCategoryTransfer PaymentCategory = "transfer"
	// PaymentCategoryReplenishment cash deposit (+ money).
	PaymentCategoryReplenishment PaymentCategory = "replenishment"
	// PaymentCategoryWriteOff charge money (- money).
	PaymentCategoryWriteOff PaymentCategory = "write-off"
)

// Payment domain model payment.
type Payment struct {
	ID            int
	CreateAt      time.Time
	AccountNumber string
	Amount        decimal.Decimal
	CompanyName   string
	Category      PaymentCategory
}
