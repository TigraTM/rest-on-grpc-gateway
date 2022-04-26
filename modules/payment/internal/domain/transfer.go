package domain

import (
	"github.com/shopspring/decimal"
)

// Transfer domain model for transferring money between users.
type Transfer struct {
	Sum           decimal.Decimal
	SenderID      int
	RecipientID   int
	RecipientName string
}
