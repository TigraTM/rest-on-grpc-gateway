package repo

import (
	"context"
	"database/sql"
	"fmt"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"

	"github.com/Masterminds/squirrel"
	"github.com/shopspring/decimal"
)

// WrapperOnSqlx wrapper on method sqlx for work with transaction and without transaction.
type WrapperOnSqlx interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func getAccountsByUserID(ctx context.Context, sqlx WrapperOnSqlx, userID int) ([]domain.Account, error) {
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
	err := sqlx.SelectContext(ctx, &accounts, query, userID)
	if err != nil {
		return nil, fmt.Errorf("r.DB.SelectContext: %w", convertErr(err))
	}

	return toDomainAccounts(accounts), nil
}

func getUserAccountByAccountNumber(ctx context.Context, sqlx WrapperOnSqlx, userID int, accountNumber string) (*domain.Account, error) {
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
	err := sqlx.GetContext(ctx, &account, query, accountNumber, userID)
	if err != nil {
		return nil, fmt.Errorf("r.DB.GetContext: %w", convertErr(err))
	}

	return toDomainAccount(account), nil
}

func getPaymentHistoryByAccountNumber(ctx context.Context, sqlx WrapperOnSqlx, userID int, accountNumber string, paging, filters filters.FilterContract) (
	_ []domain.Payment, total int, err error,
) {
	query := squirrel.Select("payment_history.id",
		"payment_history.create_at",
		"payment_history.update_at",
		"payment_history.amount",
		"payment_history.company_name",
		"payment_history.category",
		"payment_history.account_number").
		From(`"payment".payment_history`).
		LeftJoin(`payment.accounts a on a.account_number = payment_history.account_number`).
		Where(`payment_history.account_number = $1`, accountNumber).
		Where(`a.user_id = $2`, userID)

	query = paging.ApplyTo(query)
	query = filters.ApplyTo(query)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("query.ToSql: %w", err)
	}

	var payments []Payment
	err = sqlx.SelectContext(ctx, &payments, sqlQuery, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("r.DB.SelectContext: %w", convertErr(err))
	}

	total, err = getTotal(ctx, sqlx, userID, accountNumber)
	if err != nil {
		return nil, 0, fmt.Errorf("r.getTotal: %w", err)
	}

	return toDomainPayments(payments), total, nil
}

func createOrUpdateAccount(ctx context.Context, sqlx WrapperOnSqlx, userID int, accountNumber string, sum decimal.Decimal) error {
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

	_, err := sqlx.ExecContext(ctx, query, userID, accountNumber, sum, sum)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", convertErr(err))
	}

	return nil
}

func createPayment(ctx context.Context, sqlx WrapperOnSqlx, payment domain.Payment) error {
	const query = `INSERT INTO
						"payment".payment_history
							( amount,
							company_name,
							category,
							account_number )
					VALUES ($1, $2, $3, $4) `

	_, err := sqlx.ExecContext(ctx, query, payment.Amount, payment.CompanyName, payment.Category,
		payment.AccountNumber)
	if err != nil {
		return fmt.Errorf("r.DB.ExecContext: %w", convertErr(err))
	}

	return nil
}

func getTotal(ctx context.Context, sqlx WrapperOnSqlx, userID int, accountNumber string) (total int, err error) {
	const getTotal = ` SELECT 
							count(*) OVER() AS total 
						FROM  
							"payment".payment_history 
						LEFT JOIN 
							payment.accounts a on a.account_number = payment_history.account_number
						WHERE 
							payment_history.account_number = $1
						AND a.user_id = $2`

	err = sqlx.GetContext(ctx, &total, getTotal, accountNumber, userID)
	if err != nil {
		return 0, fmt.Errorf("db.GetContext: %w", convertErr(err))
	}

	return total, nil
}
