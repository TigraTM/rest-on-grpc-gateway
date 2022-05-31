package api_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	decimalpb "google.golang.org/genproto/googleapis/type/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	paymentpb "rest-on-grpc-gateway/api/proto/payment/v1"
	"rest-on-grpc-gateway/modules/payment/internal/api"
	"rest-on-grpc-gateway/modules/payment/internal/app"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
)

func TestAPI_CreatePayment(t *testing.T) {
	t.Parallel()

	var (
		errInternal              = status.Error(codes.Internal, fmt.Sprintf("a.app.CreateUser: %s", errAny))
		errNotFound              = status.Error(codes.NotFound, api.ErrNotFound.Error())
		errNotEnoughMoney        = status.Error(codes.InvalidArgument, api.ErrNotEnoughMoney.Error())
		errValidateCompanyName   = status.Error(codes.InvalidArgument, "invalid CreatePaymentRequest.CompanyName: value length must be between 1 and 50 runes, inclusive")
		errValidateCategory      = status.Error(codes.InvalidArgument, "invalid CreatePaymentRequest.Category: value length must be between 1 and 50 runes, inclusive")
		errValidateAccountNumber = status.Error(codes.InvalidArgument, "invalid CreatePaymentRequest.AccountNumber: value length must be between 1 and 20 runes, inclusive")
	)

	req := &paymentpb.CreatePaymentRequest{
		Amount: &decimalpb.Decimal{
			Value: payment.Amount.String(),
		},
		CompanyName:   payment.CompanyName,
		Category:      string(payment.Category),
		AccountNumber: payment.AccountNumber,
		UserId:        int64(userID),
	}

	testCases := map[string]struct {
		req     *paymentpb.CreatePaymentRequest
		resp    *paymentpb.CreatePaymentResponse
		appErr  error
		wantErr error
	}{
		"success": {req, &paymentpb.CreatePaymentResponse{}, nil, nil},
		"err_validate_min_company_name": {&paymentpb.CreatePaymentRequest{
			CompanyName: "",
		},
			nil, nil, errValidateCompanyName,
		},
		"err_validate_max_company_name": {&paymentpb.CreatePaymentRequest{
			CompanyName: strings.Repeat("s", 51),
		},
			nil, nil, errValidateCompanyName,
		},
		"err_validate_min_category": {&paymentpb.CreatePaymentRequest{
			CompanyName: req.CompanyName,
			Category:    "",
		},
			nil, nil, errValidateCategory,
		},
		"err_validate_max_category": {&paymentpb.CreatePaymentRequest{
			CompanyName: req.CompanyName,
			Category:    strings.Repeat("s", 51),
		},
			nil, nil, errValidateCategory,
		},
		"err_validate_min_account_number": {&paymentpb.CreatePaymentRequest{
			CompanyName:   req.CompanyName,
			Category:      req.Category,
			AccountNumber: "",
		},
			nil, nil, errValidateAccountNumber,
		},
		"err_validate_max_account_number": {&paymentpb.CreatePaymentRequest{
			CompanyName:   req.CompanyName,
			Category:      req.Category,
			AccountNumber: strings.Repeat("s", 21),
		},
			nil, nil, errValidateAccountNumber,
		},
		"err_not_found":        {req, nil, app.ErrNotFound, errNotFound},
		"err_not_enough_money": {req, nil, app.ErrNotEnoughMoney, errNotEnoughMoney},
		"err_any":              {req, nil, errAny, errInternal},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			client, app, assert := setup(t)

			if tc.resp != nil || tc.appErr != nil {
				amount, err := decimal.NewFromString(tc.req.Amount.Value)
				assert.NoError(err)

				payment := domain.Payment{
					AccountNumber: tc.req.AccountNumber,
					Amount:        amount,
					CompanyName:   tc.req.CompanyName,
					Category:      domain.PaymentCategory(tc.req.Category),
				}

				app.EXPECT().CreatePayment(gomock.Any(), int(tc.req.UserId), payment).Return(tc.appErr)
			}

			resp, err := client.CreatePayment(ctx, tc.req)
			assert.ErrorIs(err, tc.wantErr)
			assert.True(proto.Equal(resp, tc.resp))
		})
	}
}
