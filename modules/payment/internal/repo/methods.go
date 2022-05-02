package repo

import (
	"context"
	"fmt"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"

	"github.com/Masterminds/squirrel"
	"github.com/shopspring/decimal"
)

// GetAccountsByUserID get accounts by user id in db.
func (r *Repo) GetAccountsByUserID(ctx context.Context, userID int) ([]domain.Account, error) {
	const query = `SELECT 
						id,
						create_at,
						update_at,
						balance,
						currency,
						user_id,
						account_number
					FROM
						"payment".accounts
					WHERE
						user_id = $1 `

	var accounts []Account
	err := r.DB.SelectContext(ctx, &accounts, query, userID)
	if err != nil {
		return nil, fmt.Errorf("r.DB.SelectContext: %w", convertErr(err))
	}

	return toDomainAccounts(accounts), nil
}

// GetUserAccountByAccountNumber get user account by account number in db.
func (r *Repo) GetUserAccountByAccountNumber(ctx context.Context, userID int, accountNumber string) (*domain.Account, error) {
	const query = `SELECT 
						id,
						create_at,
						update_at,
						balance,
						currency,
						user_id,
						account_number
					FROM
					    "payment".accounts
					WHERE
						account_number = $1
					AND
						user_id = $2`

	account := Account{}
	err := r.DB.GetContext(ctx, &account, query, accountNumber, userID)
	if err != nil {
		return nil, fmt.Errorf("r.DB.GetContext: %w", convertErr(err))
	}

	return toDomainAccount(account), nil
}

// GetPaymentHistoryByAccountNumber get payment history by account number in db.
func (r *Repo) GetPaymentHistoryByAccountNumber(ctx context.Context, accountNumber string, paging, filters filters.FilterContract) (
	_ []domain.Payment, total int, err error,
) {
	query := squirrel.Select("id",
		"create_at",
		"update_at",
		"amount",
		"company_name",
		"category",
		"account_number").
		From(`"payment".payment_history`).
		Where("account_number = $1", accountNumber)

	query = paging.ApplyTo(query)
	query = filters.ApplyTo(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("query.ToSql: %w", err)
	}

	var payments []Payment
	err = r.DB.SelectContext(ctx, &payments, sqlQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("r.DB.SelectContext: %w", convertErr(err))
	}

	total, err = r.getTotal(ctx, accountNumber)
	if err != nil {
		return nil, 0, fmt.Errorf("r.getTotal: %w", err)
	}

	return toDomainPayments(payments), total, nil
}

// CreateOrUpdateAccount create or update account balance in db.
func (r *Repo) CreateOrUpdateAccount(ctx context.Context, userID int, accountNumber string, sum decimal.Decimal) error {
	const query = `INSERT INTO
						"payment".accounts
							( user_id, 
							 account_number,
							balance )
					VALUES ($1, $2, $3)
					ON CONFLICT
							(user_id, account_number)
					DO UPDATE
					SET
						balance = accounts.balance+$4 `

	_, err := r.DB.ExecContext(ctx, query, userID, accountNumber, sum, sum)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", convertErr(err))
	}

	return nil
}

// CreatePayment create payment in db.
func (r *Repo) CreatePayment(ctx context.Context, payment domain.Payment) error {
	const query = `INSERT INTO
						"payment".payment_history
							( amount,
							company_name,
							category,
							account_number )
					VALUES ($1, $2, $3, $4) `

	_, err := r.DB.ExecContext(ctx, query, payment.Amount, payment.CompanyName, payment.Category,
		payment.AccountNumber)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", convertErr(err))
	}

	return nil
}

// getTotal helper method for get total count payment history.
func (r *Repo) getTotal(ctx context.Context, accountNumber string) (total int, err error) {
	const getTotal = `SELECT count(*) OVER() AS total FROM  "payment".payment_history WHERE account_number = $1`

	err = r.DB.GetContext(ctx, &total, getTotal, accountNumber)
	if err != nil {
		return 0, fmt.Errorf("db.GetContext: %w", convertErr(err))
	}

	return total, nil
}
