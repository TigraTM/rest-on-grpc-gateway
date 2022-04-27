package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"rest-on-grpc-gateway/modules/payment/internal/app"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
)

func convertErr(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return app.ErrNotFound
	default:
		return err
	}
}

func toDomainAccount(account Account) *domain.Account {
	return &domain.Account{
		ID:            account.ID,
		CreateAt:      account.CreateAt,
		UpdateAt:      account.UpdateAt,
		AccountNumber: account.AccountNumber,
		Balance:       account.Balance,
		Currency:      account.Currency,
		UserID:        account.UserID,
	}
}

func toDomainPayment(payment Payment) *domain.Payment {
	return &domain.Payment{
		ID:          payment.ID,
		CreateAt:    payment.CreateAt,
		Sum:         payment.Sum,
		CompanyName: payment.CompanyName,
		Category:    payment.Category,
		AccountID:   payment.AccountID,
	}
}

func toDomainAccounts(accounts []Account) []domain.Account {
	results := make([]domain.Account, 0, len(accounts))
	for i, account := range accounts {
		results[i] = *toDomainAccount(account)
	}
	//nolint:forbidigo,gosimple // ...
	fmt.Println(fmt.Sprintf("cap: %d, len: %d", cap(results), len(results)))

	return results
}

func toDomainPayments(payments []Payment) []domain.Payment {
	results := make([]domain.Payment, 0, len(payments))

	for i, payment := range payments {
		results[i] = *toDomainPayment(payment)
	}

	return results
}
