package app_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"rest-on-grpc-gateway/modules/user/internal/app"
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

// maxTimeout for tests.
const maxTimeout = time.Second * 60

var (
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

type mocks struct {
	repo *MockRepo
	hash *MockPasswordHash
}

func setup(t *testing.T) (*app.App, *mocks, *require.Assertions) {
	t.Helper()

	ctrl := gomock.NewController(t)

	mockRepo := NewMockRepo(ctrl)
	mockHash := NewMockPasswordHash(ctrl)

	appl := app.New(mockRepo, mockHash)

	mocks := &mocks{
		repo: mockRepo,
		hash: mockHash,
	}

	return appl, mocks, require.New(t)
}

func setupCtx(t *testing.T) context.Context {
	ctx, cancelFunc := context.WithTimeout(context.Background(), maxTimeout)
	t.Cleanup(cancelFunc)

	return ctx
}
