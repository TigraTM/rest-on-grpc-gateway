package repo

import (
	"time"

	"github.com/shopspring/decimal"
)

// Account model for work with database.
type Account struct {
	ID            int             `db:"id"`
	CreateAt      time.Time       `db:"create_at"`
	UpdateAt      time.Time       `db:"update_at"`
	AccountNumber string          `db:"account_number"`
	Balance       decimal.Decimal `db:"balance"`
	Currency      string          `db:"currency"`
	UserID        int             `db:"user_id"`
}

// Payment model for work with database.
type Payment struct {
	ID            int             `db:"id"`
	CreateAt      time.Time       `db:"create_at"`
	UpdateAt      time.Time       `db:"update_at"`
	Amount        decimal.Decimal `db:"amount"`
	CompanyName   string          `db:"company_name"`
	Category      string          `db:"category"`
	AccountNumber string          `db:"account_number"`
}
