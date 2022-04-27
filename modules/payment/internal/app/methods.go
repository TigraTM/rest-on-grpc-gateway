package app

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"

	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"
)

// CreatePayment create payment and update account balance.
// Payment can be positive or negative. Payment create if:
// - recipient exist
// - balance value will not become negative after the payment
func (a *App) CreatePayment(ctx context.Context, userID int, payment domain.Payment) (err error) {
	// TODO: add check exist user

	// If sum is negative(example -1), then we have to check the balance after the subtraction,
	// save the payment history and update balance.
	if !payment.Sum.IsNegative() {
		if err = a.checkAccountBalanceByID(ctx, userID, payment.Sum); err != nil {
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

// GetAccountByUserID get account by accountID.
func (a *App) GetAccountByUserID(ctx context.Context, accountID int, currency string) (*domain.Account, error) {
	//TODO: add convert in currency
	return a.repo.GetAccountByID(ctx, accountID)
}

// TransferBetweenUsers transferring money between users.
func (a *App) TransferBetweenUsers(ctx context.Context, transfer domain.Transfer) (_ *domain.Transfer, err error) {
	if err = a.checkAccountBalanceByID(ctx, transfer.SenderAccountID, transfer.Sum); err != nil {
		return nil, fmt.Errorf("a.checkingBalanceByUserID: %w", err)
	}
	//TODO: add check exist recipient
	//TODO: add payment history|2 methods for sender and recipient
	//TODO: add transactions
	if err = a.repo.CreateOrUpdateAccount(ctx, transfer.SenderID, transfer.SenderAccountNumber, transfer.Sum); err != nil {
		return nil, fmt.Errorf("a.repo.SubBalanceByUserID: %w", err)
	}

	if err = a.repo.CreateOrUpdateAccount(ctx, transfer.RecipientID, transfer.RecipientAccountNumber, transfer.Sum); err != nil {
		return nil, fmt.Errorf("a.repo.AddBalanceByUserID: %w", err)
	}

	return nil, nil
}

// GetPaymentHistoryByAccountID get payment history by accountID.
func (a *App) GetPaymentHistoryByAccountID(ctx context.Context, userID, accountID int, paging, filter filters.FilterContract) ([]domain.Payment, int, error) {
	// TODO: add check if the account belongs to the user who made the request
	return a.repo.GetPaymentHistoryByAccountID(ctx, accountID, paging, filter)
}

// GetAccountsByUserID get all accounts for user by user id.
func (a *App) GetAccountsByUserID(ctx context.Context, userID int) ([]domain.Account, error) {
	return a.repo.GetAccountsByUserID(ctx, userID)
}

// checkAccountBalanceByID checks the balance in the client's account when deducting money.
func (a *App) checkAccountBalanceByID(ctx context.Context, accountID int, sum decimal.Decimal) error {
	senderBalance, err := a.repo.GetAccountByID(ctx, accountID)
	if err != nil {
		return fmt.Errorf("a.repo.GetBalanceByUserID: %w", err)
	}

	if senderBalance.Balance.Sub(sum).IsNegative() {
		return ErrNotEnoughMoney
	}

	return nil
}
