package domain

import (
	"github.com/shopspring/decimal"
)

// Transfer domain model for transferring money between users.
type Transfer struct {
	Sum                    decimal.Decimal
	SenderID               int
	SenderAccountID        int
	SenderAccountNumber    string
	RecipientID            int
	RecipientAccountID     int
	RecipientAccountNumber string
	RecipientName          string
}