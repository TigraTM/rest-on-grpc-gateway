package app_test

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"rest-on-grpc-gateway/modules/payment/internal/app"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
	"rest-on-grpc-gateway/modules/payment/internal/filters"
)

//func TestApp_TransferBetweenUsers(t *testing.T) {
//	t.Parallel()
//
//	successTransfer := domain.Transfer{
//		Amount:                 decimal.New(50, 0),
//		SenderID:               1,
//		SenderAccountNumber:    "123",
//		RecipientID:            2,
//		RecipientAccountNumber: "321",
//		RecipientName:          "Elon",
//	}
//
//	testCases := []struct {
//		name     string
//		transfer domain.Transfer
//		want     *domain.Transfer
//		wantErr  error
//		prepare  func(m *mocks)
//	}{
//		{
//			name:     "success",
//			transfer: successTransfer,
//			want:     &successTransfer,
//			wantErr:  nil,
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//		t.Run(tc.name, func(t *testing.T) {
//			t.Parallel()
//
//			ctx, module, mocks, assert := setup(t)
//			if tc.prepare != nil {
//				tc.prepare(mocks)
//			}
//
//			mocks.user.EXPECT().ExistUserByID(ctx, tc.transfer.RecipientID).Return(nil)
//
//			mocks.repo.EXPECT().GetUserAccountByAccountNumber(ctx, tc.transfer.SenderID, tc.transfer.SenderAccountNumber).
//				Return(&account, nil)
//
//			mocks.txRepo.EXPECT().DoTx(ctx, func(app.Repo) error {
//				mocks.repo.EXPECT().CreateOrUpdateAccount(ctx, tc.transfer.SenderID, tc.transfer.SenderAccountNumber,
//					tc.transfer.Amount.Neg()).Return(nil)
//
//				mocks.repo.EXPECT().CreatePayment(ctx, domain.Payment{
//					AccountNumber: tc.transfer.SenderAccountNumber,
//					Amount:        tc.transfer.Amount.Neg(),
//					CompanyName:   tc.transfer.RecipientName,
//					Category:      domain.PaymentCategoryTransfer,
//				}).Return(nil)
//
//				mocks.repo.EXPECT().CreateOrUpdateAccount(ctx, tc.transfer.RecipientID, tc.transfer.RecipientAccountNumber,
//					tc.transfer.Amount).Return(nil)
//
//				mocks.repo.EXPECT().CreatePayment(ctx, domain.Payment{
//					AccountNumber: tc.transfer.RecipientAccountNumber,
//					Amount:        tc.transfer.Amount,
//					CompanyName:   "senderName",
//					Category:      domain.PaymentCategoryTransfer,
//				}).Return(nil)
//
//				return nil
//			})
//
//			transfer, err := module.TransferBetweenUsers(ctx, tc.transfer)
//			assert.ErrorIs(err, tc.wantErr)
//			assert.Equal(tc.want, transfer)
//		})
//	}
//}

func TestApp_GetPaymentHistoryByAccountID(t *testing.T) {
	t.Parallel()

	payment := domain.Payment{
		ID:            1,
		CreateAt:      time.Now(),
		AccountNumber: account.AccountNumber,
		Amount:        decimal.New(100, 0),
		CompanyName:   "xxx",
		Category:      domain.PaymentCategoryTransfer,
	}

	testCases := []struct {
		name           string
		userID         int
		accountNumber  string
		paging, filter filters.FilterContract
		want           []domain.Payment
		wantTotal      int
		wantErr        error
	}{
		{
			name:          "success",
			userID:        account.UserID,
			accountNumber: account.AccountNumber,
			paging:        nil,
			filter:        nil,
			want:          []domain.Payment{payment},
			wantTotal:     1,
			wantErr:       nil,
		},
		{
			name:          "err_any",
			userID:        account.UserID,
			accountNumber: account.AccountNumber,
			paging:        nil,
			filter:        nil,
			want:          nil,
			wantTotal:     0,
			wantErr:       errAny,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx, module, mocks, assert := setup(t)

			mocks.repo.EXPECT().GetPaymentHistoryByAccountNumber(ctx, tc.userID, tc.accountNumber, tc.paging, tc.filter).
				Return(tc.want, tc.wantTotal, tc.wantErr)

			payments, total, err := module.GetPaymentHistoryByAccountID(ctx, tc.userID, tc.accountNumber, tc.paging, tc.filter)
			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantTotal, total)
			assert.Equal(tc.want, payments)
		})
	}
}

func TestApp_GetAccountsByUserID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		userID  int
		want    []domain.Account
		wantErr error
	}{
		{
			name:    "success",
			userID:  userID,
			want:    []domain.Account{account},
			wantErr: nil,
		},
		{
			name:    "err_not_found",
			userID:  userID,
			want:    nil,
			wantErr: app.ErrNotFound,
		},
		{
			name:    "err_any",
			userID:  userID,
			want:    nil,
			wantErr: errAny,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx, module, mocks, assert := setup(t)

			mocks.repo.EXPECT().GetAccountsByUserID(ctx, tc.userID).Return(tc.want, tc.wantErr)

			accounts, err := module.GetAccountsByUserID(ctx, tc.userID)
			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.want, accounts)
		})
	}
}

func TestApp_GetAllCurrencies(t *testing.T) {
	t.Parallel()

	want := map[string]string{
		"USD": "United States Dollar",
	}

	testCases := []struct {
		name    string
		want    map[string]string
		wantErr error
	}{
		{
			name:    "success",
			want:    want,
			wantErr: nil,
		},
		{
			name:    "err_exchange_client",
			want:    nil,
			wantErr: app.ErrExchangeClient,
		},
		{
			name:    "err_any",
			want:    nil,
			wantErr: errAny,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctx, module, mocks, assert := setup(t)

			mocks.exchange.EXPECT().GetSymbols(ctx).Return(tc.want, tc.wantErr)

			currencies, err := module.GetAllCurrencies(ctx)
			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.want, currencies)
		})
	}
}
