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

	// If sum is negative(example -1), then we have to check the balance after the subtraction,
	// save the payment history and update balance.
	if payment.Sum.IsNegative() {
		if err = a.checkAccountBalanceByID(ctx, payment.AccountNumber, payment.Sum); err != nil {
			return fmt.Errorf("a.checkingBalanceByUserID: %w", err)
		}
		// TODO: add create payment history|1 method for 1 user
		if err = a.repo.CreateOrUpdateAccount(ctx, userID, payment.AccountNumber, payment.Sum); err != nil {
			return fmt.Errorf("a.repo.SubBalanceByUserID: %w", err)
		}
	}
	// if sim is positive(example +1), we don't check the balance and can add sum to balance.

	// TODO: add create payment history|1 method for 1 user
	return a.repo.CreateOrUpdateAccount(ctx, userID, payment.AccountNumber, payment.Sum)
}

// GetAccountByAccountNumber get account by account number.
func (a *App) GetAccountByAccountNumber(ctx context.Context, accountNumber, _ string) (*domain.Account, error) {
	// TODO: add check if the account belongs to the user who made the request
	return a.repo.GetAccountByAccountNumber(ctx, accountNumber)
	// TODO: add convert in currency
}

// TransferBetweenUsers transferring money between users.
func (a *App) TransferBetweenUsers(ctx context.Context, transfer domain.Transfer) (_ *domain.Transfer, err error) {
	// TODO: add check exist recipient
	// TODO: add check exist recipient account

	if err = a.checkAccountBalanceByID(ctx, transfer.SenderAccountNumber, transfer.Sum); err != nil {
		return nil, fmt.Errorf("a.checkingBalanceByUserID: %w", err)
	}

	// TODO: add payment history|2 methods for sender and recipient
	// TODO: add transactions
	if err = a.repo.CreateOrUpdateAccount(ctx, transfer.SenderID, transfer.SenderAccountNumber, transfer.Sum); err != nil {
		return nil, fmt.Errorf("a.repo.SubBalanceByUserID: %w", err)
	}

	if err = a.repo.CreateOrUpdateAccount(ctx, transfer.RecipientID, transfer.RecipientAccountNumber, transfer.Sum); err != nil {
		return nil, fmt.Errorf("a.repo.AddBalanceByUserID: %w", err)
	}

	return &domain.Transfer{
		Sum:                    transfer.Sum,
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
func (a *App) checkAccountBalanceByID(ctx context.Context, accountNumber string, sum decimal.Decimal) error {
	senderBalance, err := a.repo.GetAccountByAccountNumber(ctx, accountNumber)
	if err != nil {
		return fmt.Errorf("a.repo.GetBalanceByUserID: %w", err)
	}

	if senderBalance.Balance.Sub(sum).IsNegative() {
		return ErrNotEnoughMoney
	}

	return nil
}
