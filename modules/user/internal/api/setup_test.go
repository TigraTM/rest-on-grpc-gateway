package api_test

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
	"rest-on-grpc-gateway/modules/user/internal/api"
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

// maxTimeout for tests.
const maxTimeout = time.Second * 60

// api errors.
var (
	errUserNotFound    = status.Error(codes.NotFound, "user not found")
	errInvalidPassword = status.Error(codes.InvalidArgument, "invalid password")
	errEmailExist      = status.Error(codes.InvalidArgument, "user with this email already exists")
	errMustDifferent   = status.Error(codes.InvalidArgument, "value of the existing password and the new password must be different")
)

var (
	ctx      = context.Background()
	errAny   = errors.New("err any")
	userID   = 1
	userName = "user"
	password = "12345678"
	email    = "user@mail.com"
	user     = &domain.User{
		ID:    userID,
		Name:  userName,
		Email: email,
	}
)

func setup(t *testing.T) (userpb.UserAPIClient, *Mockapplication, *require.Assertions) {
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

	// TODO: fix work with ctx in api tests.
	ctx, cancelFunc := context.WithTimeout(context.Background(), maxTimeout)
	t.Cleanup(cancelFunc)

	server := api.New(log.Sugar(), mockApp)

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

	return userpb.NewUserAPIClient(conn), mockApp, assert
}
