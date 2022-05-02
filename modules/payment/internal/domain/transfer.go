package domain

import (
	"github.com/shopspring/decimal"
)

// Transfer domain model for transferring money between users.
type Transfer struct {
	Amount                 decimal.Decimal
	SenderID               int
	SenderAccountNumber    string
	RecipientID            int
	RecipientAccountNumber string
	RecipientName          string
}
