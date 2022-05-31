package app_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"rest-on-grpc-gateway/modules/payment/internal/app"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
)

// maxTimeout for tests.
const maxTimeout = time.Second * 60

var (
	userID  = 1
	account = domain.Account{
		ID:            1,
		CreateAt:      time.Now(),
		UpdateAt:      time.Now(),
		AccountNumber: "111",
		Balance:       decimal.New(100, 1),
		Currency:      "USD",
		UserID:        userID,
	}
	errAny = errors.New("err any")
)

type mocks struct {
	repo     *MockRepo
	txRepo   *MockTxRepo
	exchange *MockExchangeClient
	user     *MockUserClient
}

func setup(t *testing.T) (context.Context, *app.App, *mocks, *require.Assertions) {
	t.Helper()

	ctrl := gomock.NewController(t)

	mockRepo := NewMockRepo(ctrl)
	mockTxRepo := NewMockTxRepo(ctrl)
	mockExchange := NewMockExchangeClient(ctrl)
	mockUser := NewMockUserClient(ctrl)

	appl := app.New(mockRepo, mockTxRepo, mockExchange, mockUser)

	mocks := &mocks{
		repo:     mockRepo,
		txRepo:   mockTxRepo,
		exchange: mockExchange,
		user:     mockUser,
	}

	return setupCtx(t), appl, mocks, require.New(t)
}

func setupCtx(t *testing.T) context.Context {
	ctx, cancelFunc := context.WithTimeout(context.Background(), maxTimeout)
	t.Cleanup(cancelFunc)

	return ctx
}
