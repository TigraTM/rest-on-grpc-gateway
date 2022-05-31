package api_test

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	paymentpb "rest-on-grpc-gateway/api/proto/payment/v1"
	"rest-on-grpc-gateway/modules/payment/internal/api"
	"rest-on-grpc-gateway/modules/payment/internal/domain"
)

// maxTimeout for tests.
const maxTimeout = time.Second * 60

var (
	ctx     = context.Background()
	errAny  = errors.New("err any")
	userID  = 1
	payment = domain.Payment{
		ID:            1,
		CreateAt:      time.Now(),
		AccountNumber: "123",
		Amount:        decimal.New(100, 0),
		CompanyName:   "Test Company",
		Category:      "Test Category",
	}
)

func setup(t *testing.T) (paymentpb.PaymentExternalAPIClient, *Mockapplication, *require.Assertions) {
	t.Helper()

	assert := require.New(t)

	ctrl := gomock.NewController(t)

	mockApp := NewMockapplication(ctrl)

	logCfg := zap.NewProductionConfig()
	logCfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	log, err := logCfg.Build(
		zap.WithClock(zapcore.DefaultClock),
		zap.AddCaller(),
	)

	// TODO: fix work with ctx in apiExternal tests.
	ctx, cancelFunc := context.WithTimeout(context.Background(), maxTimeout)
	t.Cleanup(cancelFunc)

	server := api.New(log, mockApp)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	assert.NoError(err)

	go func() {
		err := server.Serve(ln)
		assert.NoError(err)
	}()

	conn, err := grpc.DialContext(ctx, ln.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()), // TODO Add TLS and remove this.
		grpc.WithBlock(),
	)
	assert.NoError(err)

	t.Cleanup(func() {
		err := conn.Close()
		assert.NoError(err)
		server.GracefulStop()
	})

	return paymentpb.NewPaymentExternalAPIClient(conn), mockApp, assert
}
