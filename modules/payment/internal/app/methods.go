package app

import (
	"context"
	"fmt"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"

	"github.com/shopspring/decimal"
)

// CreatePayment create payment and update account balance.
// Payment can be positive or negative. Payment create if:
// - recipient exist
// - balance value will not become negative after the payment.
func (a *App) CreatePayment(ctx context.Context, userID int, payment domain.Payment) (err error) {
	// TODO: add check exist user, maybe helped session

	// if sum is negative, need to check value account balance after the payment has been charged
	if payment.Amount.IsNegative() {
		if err = a.checkAccountBalanceByID(ctx, userID, payment.AccountNumber, payment.Amount); err != nil {
			return fmt.Errorf("a.checkingBalanceByUserID: %w", err)
		}
	}

	err = a.txRepo.DoTx(ctx, func(repo Repo) error {
		if err = repo.CreateOrUpdateAccount(ctx, userID, payment.AccountNumber, payment.Amount); err != nil {
			return fmt.Errorf("a.repo.CreateOrUpdateAccount: %w", err)
		}

		return repo.CreatePayment(ctx, payment)
	})

	return nil
}

// GetAccountByAccountNumber get account by account number.
func (a *App) GetAccountByAccountNumber(ctx context.Context, userID int, accountNumber, needCurrency string) (*domain.Account, error) {
	account, err := a.repo.GetUserAccountByAccountNumber(ctx, userID, accountNumber)
	if err != nil {
		return nil, fmt.Errorf("a.repo.GetUserAccountByAccountNumber: %w", err)
	}

	convertBalance, err := a.exchange.ConvertAmount(ctx, account.Currency, needCurrency, account.Balance)
	if err != nil {
		return nil, fmt.Errorf("a.exchange.Convert: %w", err)
	}
	account.Balance = convertBalance
	account.Currency = needCurrency

	return account, nil
}

// TransferBetweenUsers transferring money between users.
func (a *App) TransferBetweenUsers(ctx context.Context, transfer domain.Transfer) (_ *domain.Transfer, err error) {
	if transfer.SenderAccountNumber == transfer.RecipientAccountNumber {
		return nil, ErrSameAccountNumber
	}

	if transfer.Amount.IsNegative() {
		return nil, ErrTransferAmountNotCorrect
	}
	// TODO: add check exist recipient

	if err = a.checkAccountBalanceByID(ctx, transfer.SenderID, transfer.SenderAccountNumber, transfer.Amount); err != nil {
		return nil, fmt.Errorf("a.checkingBalanceByUserID: %w", err)
	}

	err = a.txRepo.DoTx(ctx, func(repo Repo) error {
		if err = repo.CreateOrUpdateAccount(ctx, transfer.SenderID, transfer.SenderAccountNumber, transfer.Amount.Neg()); err != nil {
			return fmt.Errorf("a.repo.SubBalanceByUserID: %w", err)
		}

		if err = repo.CreatePayment(ctx, domain.Payment{
			AccountNumber: transfer.SenderAccountNumber,
			Amount:        transfer.Amount.Neg(),
			CompanyName:   transfer.RecipientName,
			Category:      domain.PaymentCategoryTransfer,
		}); err != nil {
			return fmt.Errorf("a.repo.CreatePayment: %w", err)
		}

		if err = repo.CreateOrUpdateAccount(ctx, transfer.RecipientID, transfer.RecipientAccountNumber, transfer.Amount); err != nil {
			return fmt.Errorf("a.repo.AddBalanceByUserID: %w", err)
		}

		if err = repo.CreatePayment(ctx, domain.Payment{
			AccountNumber: transfer.RecipientAccountNumber,
			Amount:        transfer.Amount,
			CompanyName:   "senderName", // TODO: add sender name in session
			Category:      domain.PaymentCategoryTransfer,
		}); err != nil {
			return fmt.Errorf("a.repo.CreatePayment: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("a.txRepo.DoTx: %w", err)
	}

	return &domain.Transfer{
		Amount:                 transfer.Amount,
		RecipientID:            transfer.RecipientID,
		RecipientAccountNumber: transfer.RecipientAccountNumber,
		RecipientName:          transfer.RecipientName,
	}, nil
}

// GetPaymentHistoryByAccountID get payment history by accountID.
func (a *App) GetPaymentHistoryByAccountID(ctx context.Context, _ int, accountNumber string, paging, filter filters.FilterContract) ([]domain.Payment, int, error) {
	// TODO: add check if the account belongs to the user who made the request
	return a.repo.GetPaymentHistoryByAccountNumber(ctx, accountNumber, paging, filter)
}

// GetAccountsByUserID get all accounts for user by user id.
func (a *App) GetAccountsByUserID(ctx context.Context, userID int) ([]domain.Account, error) {
	return a.repo.GetAccountsByUserID(ctx, userID)
}

// checkAccountBalanceByID checks the balance in the client's account when deducting money.
func (a *App) checkAccountBalanceByID(ctx context.Context, userID int, accountNumber string, sum decimal.Decimal) error {
	senderBalance, err := a.repo.GetUserAccountByAccountNumber(ctx, userID, accountNumber)
	if err != nil {
		return fmt.Errorf("a.repo.GetBalanceByUserID: %w", err)
	}

	if senderBalance.Balance.Sub(sum).IsNegative() {
		return ErrNotEnoughMoney
	}

	return nil
}

// GetAllCurrencies return all currencies with platform APILayer.
func (a *App) GetAllCurrencies(ctx context.Context) (map[string]string, error) {
	return a.exchange.GetSymbols(ctx)
}
