package api

import (
	"context"
	"errors"
	"fmt"
	"rest-on-grpc-gateway/modules/payment/internal/app"
	"rest-on-grpc-gateway/modules/payment/internal/domain"

	"github.com/shopspring/decimal"

	decimalpb "google.golang.org/genproto/googleapis/type/decimal"

	"google.golang.org/protobuf/types/known/timestamppb"

	paymentpb "rest-on-grpc-gateway/api/proto/payment/v1"
)

// CreatePayment implements userpb.UserAPIServer.
func (a *api) CreatePayment(ctx context.Context, in *paymentpb.CreatePaymentRequest) (*paymentpb.CreatePaymentResponse, error) {
	amount, err := decimal.NewFromString(in.Amount.Value)
	if err != nil {
		return nil, fmt.Errorf("decimal.NewFromString: %w", err)
	}

	newPayment := domain.Payment{
		AccountNumber: in.AccountNumber,
		Amount:        amount,
		CompanyName:   in.CompanyName,
		Category:      domain.PaymentCategory(in.Category),
	}

	err = a.app.CreatePayment(ctx, int(in.UserId), newPayment)
	switch {
	case err == nil:
		return &paymentpb.CreatePaymentResponse{}, nil
	case errors.Is(err, app.ErrNotEnoughMoney):
		return nil, errNotEnoughMoney
	default:
		return nil, fmt.Errorf("a.app.CreateUser: %w", err)
	}
}

// GetAccountByAccountNumber implements userpb.UserAPIServer.
func (a *api) GetAccountByAccountNumber(ctx context.Context, in *paymentpb.GetAccountByUserIDRequest) (*paymentpb.GetAccountByUserIDResponse, error) {
	account, err := a.app.GetAccountByAccountNumber(ctx, int(in.UserId), in.AccountNumber, in.Currency)
	switch {
	case err == nil:
		return &paymentpb.GetAccountByUserIDResponse{
			Balance: &decimalpb.Decimal{
				Value: account.Balance.String(),
			},
			Currency:      account.Currency,
			AccountNumber: account.AccountNumber,
		}, nil
	case errors.Is(err, app.ErrNotFound):
		return nil, errNotFound
	default:
		return nil, fmt.Errorf("a.app.GetAccountByUserID: %w", err)
	}
}

// TransferBetweenUsers implements userpb.UserAPIServer.
func (a *api) TransferBetweenUsers(ctx context.Context, in *paymentpb.TransferBetweenUsersRequest) (*paymentpb.TransferBetweenUsersResponse, error) {
	amount, err := decimal.NewFromString(in.Amount.Value)
	if err != nil {
		return nil, fmt.Errorf("decimal.NewFromString: %w", err)
	}

	newTransfer := domain.Transfer{
		Amount:                 amount,
		SenderID:               int(in.SenderId),
		SenderAccountNumber:    in.SenderAccountNumber,
		RecipientID:            int(in.RecipientId),
		RecipientAccountNumber: in.RecipientAccountNumber,
		RecipientName:          in.RecipientName,
	}

	transfer, err := a.app.TransferBetweenUsers(ctx, newTransfer)
	switch {
	case err == nil:
		return &paymentpb.TransferBetweenUsersResponse{
			Amount: &decimalpb.Decimal{
				Value: transfer.Amount.String(),
			},
			RecipientId:            int64(transfer.RecipientID),
			RecipientName:          transfer.RecipientName,
			RecipientAccountNumber: transfer.RecipientAccountNumber,
		}, nil
	case errors.Is(err, app.ErrNotFound):
		return nil, errNotFound
	case errors.Is(err, app.ErrNotEnoughMoney):
		return nil, errNotEnoughMoney
	case errors.Is(err, app.ErrSameAccountNumber):
		return nil, errSameAccountNumber
	case errors.Is(err, app.ErrTransferAmountNotCorrect):
		return nil, errTransferAmountNotCorrect
	default:
		return nil, fmt.Errorf("a.app.TransferBetweenUsers: %w", err)
	}
}

// GetPaymentsHistoryByAccountNumber implements userpb.UserAPIServer.
func (a *api) GetPaymentsHistoryByAccountNumber(ctx context.Context, in *paymentpb.GetPaymentsHistoryByAccountIDRequest) (*paymentpb.GetPaymentsHistoryByAccountIDResponse, error) {
	paging, err := getPaging(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errUncorrectedPaging, err.Error())
	}

	sort, err := getSortForPaymentHistory(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errUncorrectedSort, err.Error())
	}

	payments, total, err := a.app.GetPaymentHistoryByAccountID(ctx, int(in.UserId), in.AccountNumber, &paging, &sort)
	switch {
	case errors.Is(err, app.ErrNotFound):
		return nil, errNotFound
	case err != nil:
		return nil, fmt.Errorf("a.app.GetPaymentHistoryByAccountID: %w", err)
	}

	pbPayments := make([]*paymentpb.Payment, len(payments))
	for i := range payments {
		pbPayments[i] = &paymentpb.Payment{
			Id:            int64(payments[i].ID),
			CreateAt:      timestamppb.New(payments[i].CreateAt),
			AccountNumber: payments[i].AccountNumber,
			Amount: &decimalpb.Decimal{
				Value: payments[i].Amount.String(),
			},
			CompanyName: payments[i].CompanyName,
			Category:    string(payments[i].Category),
		}
	}

	return &paymentpb.GetPaymentsHistoryByAccountIDResponse{
		Payments: pbPayments,
		Total:    int64(total),
	}, nil
}

// GetAccountsByUserID implements userpb.UserAPIServer.
func (a *api) GetAccountsByUserID(ctx context.Context, in *paymentpb.GetAccountsByUserIDRequest) (*paymentpb.GetAccountsByUserIDResponse, error) {
	accounts, err := a.app.GetAccountsByUserID(ctx, int(in.UserId))
	if err != nil {
		return nil, fmt.Errorf("a.app.GetAccountsByUserID: %w", err)
	}

	pbAccounts := make([]*paymentpb.Account, len(accounts))
	for i := range accounts {
		pbAccounts[i] = &paymentpb.Account{
			Balance: &decimalpb.Decimal{
				Value: accounts[i].Balance.String(),
			},
			Currency:      accounts[i].Currency,
			AccountNumber: accounts[i].AccountNumber,
		}
	}

	return &paymentpb.GetAccountsByUserIDResponse{Accounts: pbAccounts}, nil
}
