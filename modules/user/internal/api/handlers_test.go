package api_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	userpb "rest-on-grpc-gateway/api/proto/user/v1"
	"rest-on-grpc-gateway/modules/user/internal/app"
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

func TestAPI_CreateUser(t *testing.T) {
	var (
		errInternal                = status.Error(codes.Internal, fmt.Sprintf("a.app.CreateUser: %s", errAny))
		errInvalidArgumentUsername = status.Error(codes.InvalidArgument, "invalid CreateUserRequest.Name: value length must be between 2 and 40 runes, inclusive")
		errInvalidArgumentEmail    = status.Error(codes.InvalidArgument, "invalid CreateUserRequest.Email: value must be a valid email address | caused by: mail: missing '@' or angle-addr")
		errInvalidArgumentEmailLen = status.Error(codes.InvalidArgument, "invalid CreateUserRequest.Email: value length must be at most 50 runes")
		errInvalidArgumentPassword = status.Error(codes.InvalidArgument, "invalid CreateUserRequest.Password: value length must be between 8 and 100 runes, inclusive")
	)

	req := &userpb.CreateUserRequest{
		Name:     userName,
		Email:    email,
		Password: password,
	}
	resp := &userpb.CreateUserResponse{
		Id:    int64(userID),
		Name:  userName,
		Email: email,
	}

	tests := []struct {
		name    string
		req     *userpb.CreateUserRequest
		resp    *userpb.CreateUserResponse
		wantErr error
		prepare func(m *Mockapplication)
	}{
		{
			name:    "success",
			req:     req,
			resp:    resp,
			wantErr: nil,
			prepare: func(m *Mockapplication) {
				m.EXPECT().CreateUser(gomock.Any(), req.Name, req.Email, req.Password).Return(user, nil).Times(1)
			},
		},
		{
			name: "err name is too short",
			req: &userpb.CreateUserRequest{
				Name:     "a",
				Email:    "user@mail.com",
				Password: "12345678",
			},
			resp:    nil,
			wantErr: errInvalidArgumentUsername,
			prepare: func(m *Mockapplication) {},
		},
		{
			name: "err invalid name is too long",
			req: &userpb.CreateUserRequest{
				Name:     strings.Repeat("a", 41),
				Email:    "user@mail.com",
				Password: "12345678",
			},
			resp:    nil,
			wantErr: errInvalidArgumentUsername,
			prepare: func(m *Mockapplication) {},
		},
		{
			name: "err invalid email",
			req: &userpb.CreateUserRequest{
				Name:     strings.Repeat("a", 40),
				Email:    "usermail.com",
				Password: "12345678",
			},
			resp:    nil,
			wantErr: errInvalidArgumentEmail,
			prepare: func(m *Mockapplication) {},
		},
		{
			name: "err invalid email",
			req: &userpb.CreateUserRequest{
				Name:     "user",
				Email:    fmt.Sprintf("%s%s", strings.Repeat("a", 30), "mail.com"),
				Password: "12345678",
			},
			resp:    nil,
			wantErr: errInvalidArgumentEmail,
			prepare: func(m *Mockapplication) {},
		},
		{
			name: "err invalid email to long",
			req: &userpb.CreateUserRequest{
				Name:     "user",
				Email:    fmt.Sprintf("%s%s", strings.Repeat("a", 42), "@mail.com"),
				Password: "12345678",
			},
			resp:    nil,
			wantErr: errInvalidArgumentEmailLen,
			prepare: func(m *Mockapplication) {},
		},
		{
			name: "err invalid password to short",
			req: &userpb.CreateUserRequest{
				Name:     "user",
				Email:    "user@mail.com",
				Password: "1234",
			},
			resp:    nil,
			wantErr: errInvalidArgumentPassword,
			prepare: func(m *Mockapplication) {},
		},
		{
			name: "err invalid password to short",
			req: &userpb.CreateUserRequest{
				Name:     "user",
				Email:    "user@mail.com",
				Password: strings.Repeat("1", 101),
			},
			resp:    nil,
			wantErr: errInvalidArgumentPassword,
			prepare: func(m *Mockapplication) {},
		},
		{
			name:    "err email exist",
			req:     req,
			resp:    nil,
			wantErr: errEmailExist,
			prepare: func(m *Mockapplication) {
				m.EXPECT().CreateUser(gomock.Any(), req.Name, req.Email, req.Password).Return(nil, app.ErrEmailExist).Times(1)
			},
		},
		{
			name:    "err any",
			req:     req,
			resp:    nil,
			wantErr: errInternal,
			prepare: func(m *Mockapplication) {
				m.EXPECT().CreateUser(gomock.Any(), req.Name, req.Email, req.Password).Return(nil, errAny).Times(1)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, app, assert := setup(t)
			if tt.prepare != nil {
				tt.prepare(app)
			}

			resp, err := client.CreateUser(ctx, tt.req)
			assert.ErrorIs(err, tt.wantErr)
			assert.True(proto.Equal(resp, tt.resp))
		})
	}
}

func TestApi_GetUserByID(t *testing.T) {
	var errInternal = status.Error(codes.Internal, fmt.Sprintf("a.app.GetUserByID: %s", errAny))

	req := &userpb.GetUserByIDRequest{
		Id: int64(userID),
	}
	resp := &userpb.GetUserByIDResponse{
		Id:    int64(userID),
		Name:  "user",
		Email: "user@mail.com",
	}

	tests := []struct {
		name    string
		req     *userpb.GetUserByIDRequest
		resp    *userpb.GetUserByIDResponse
		wantErr error
		prepare func(m *Mockapplication)
	}{
		{
			name:    "success",
			req:     req,
			resp:    resp,
			wantErr: nil,
			prepare: func(m *Mockapplication) {
				m.EXPECT().GetUserByID(gomock.Any(), int(req.Id)).Return(user, nil).Times(1)
			},
		},
		{
			name:    "err user not found",
			req:     req,
			resp:    nil,
			wantErr: errUserNotFound,
			prepare: func(m *Mockapplication) {
				m.EXPECT().GetUserByID(gomock.Any(), int(req.Id)).Return(nil, app.ErrNotFound).Times(1)
			},
		},
		{
			name:    "err any",
			req:     req,
			resp:    nil,
			wantErr: errInternal,
			prepare: func(m *Mockapplication) {
				m.EXPECT().GetUserByID(gomock.Any(), int(req.Id)).Return(nil, errAny).Times(1)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, app, assert := setup(t)
			if tt.prepare != nil {
				tt.prepare(app)
			}

			resp, err := client.GetUserByID(ctx, tt.req)
			assert.ErrorIs(err, tt.wantErr)
			assert.True(proto.Equal(resp, tt.resp))
		})
	}
}

func TestApi_UpdateUserByID(t *testing.T) {
	var errInternal = status.Error(codes.Internal, fmt.Sprintf("a.app.UpdateUserByID: %s", errAny))

	req := &userpb.UpdateUserByIDRequest{
		Name:  "user",
		Email: "user@mail.com",
	}
	resp := &userpb.UpdateUserByIDResponse{
		Id:    int64(userID),
		Name:  "user",
		Email: "user@mail.com",
	}
	updateUser := &domain.User{
		ID:    userID,
		Name:  "user",
		Email: "user@mail.com",
	}

	tests := []struct {
		name    string
		req     *userpb.UpdateUserByIDRequest
		resp    *userpb.UpdateUserByIDResponse
		wantErr error
		prepare func(m *Mockapplication)
	}{
		{
			name:    "success",
			req:     req,
			resp:    resp,
			wantErr: nil,
			prepare: func(m *Mockapplication) {
				m.EXPECT().UpdateUserByID(gomock.Any(), int(req.Id), req.Name, req.Email).Return(updateUser, nil).Times(1)
			},
		},
		{
			name:    "err user not found",
			req:     req,
			resp:    nil,
			wantErr: errUserNotFound,
			prepare: func(m *Mockapplication) {
				m.EXPECT().UpdateUserByID(gomock.Any(), int(req.Id), req.Name, req.Email).Return(nil, app.ErrNotFound).Times(1)
			},
		},
		{
			name:    "err any",
			req:     req,
			resp:    nil,
			wantErr: errInternal,
			prepare: func(m *Mockapplication) {
				m.EXPECT().UpdateUserByID(gomock.Any(), int(req.Id), req.Name, req.Email).Return(nil, errAny).Times(1)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, app, assert := setup(t)
			if tt.prepare != nil {
				tt.prepare(app)
			}

			resp, err := client.UpdateUserByID(ctx, tt.req)
			assert.ErrorIs(err, tt.wantErr)
			assert.True(proto.Equal(resp, tt.resp))
		})
	}
}

func TestApi_UpdateUserPasswordByID(t *testing.T) {
	var errInternal = status.Error(codes.Internal, fmt.Sprintf("a.app.UpdateUserPasswordByID: %s", errAny))

	req := &userpb.UpdateUserPasswordByIDRequest{
		Id:          int64(userID),
		OldPassword: "12345678",
		NewPassword: "87654321",
	}

	tests := []struct {
		name    string
		req     *userpb.UpdateUserPasswordByIDRequest
		resp    *userpb.UpdateUserPasswordByIDResponse
		wantErr error
		prepare func(m *Mockapplication)
	}{
		{
			name:    "success",
			req:     req,
			resp:    &userpb.UpdateUserPasswordByIDResponse{},
			wantErr: nil,
			prepare: func(m *Mockapplication) {
				m.EXPECT().UpdateUserPasswordByID(gomock.Any(), int(req.Id), req.OldPassword, req.NewPassword).Return(nil).Times(1)
			},
		},
		{
			name:    "err user not found",
			req:     req,
			resp:    nil,
			wantErr: errUserNotFound,
			prepare: func(m *Mockapplication) {
				m.EXPECT().UpdateUserPasswordByID(gomock.Any(), int(req.Id), req.OldPassword, req.NewPassword).Return(app.ErrNotFound).Times(1)
			},
		},
		{
			name:    "err invalid password",
			req:     req,
			resp:    nil,
			wantErr: errInvalidPassword,
			prepare: func(m *Mockapplication) {
				m.EXPECT().UpdateUserPasswordByID(gomock.Any(), int(req.Id), req.OldPassword, req.NewPassword).Return(app.ErrInvalidPassword).Times(1)
			},
		},
		{
			name:    "err value must different",
			req:     req,
			resp:    nil,
			wantErr: errMustDifferent,
			prepare: func(m *Mockapplication) {
				m.EXPECT().UpdateUserPasswordByID(gomock.Any(), int(req.Id), req.OldPassword, req.NewPassword).Return(app.ErrMustDifferent).Times(1)
			},
		},
		{
			name:    "err any",
			req:     req,
			resp:    nil,
			wantErr: errInternal,
			prepare: func(m *Mockapplication) {
				m.EXPECT().UpdateUserPasswordByID(gomock.Any(), int(req.Id), req.OldPassword, req.NewPassword).Return(errAny).Times(1)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, app, assert := setup(t)
			if tt.prepare != nil {
				tt.prepare(app)
			}

			resp, err := client.UpdateUserPasswordByID(ctx, tt.req)
			assert.ErrorIs(err, tt.wantErr)
			assert.True(proto.Equal(resp, tt.resp))
		})
	}
}

func TestApi_DeleteUserByID(t *testing.T) {
	var errInternal = status.Error(codes.Internal, fmt.Sprintf("a.app.DeleteUserByID: %s", errAny))

	req := &userpb.DeleteUserByIDRequest{
		Id: int64(userID),
	}

	tests := []struct {
		name    string
		req     *userpb.DeleteUserByIDRequest
		resp    *userpb.DeleteUserByIDResponse
		wantErr error
		prepare func(m *Mockapplication)
	}{
		{
			name:    "success",
			req:     req,
			resp:    &userpb.DeleteUserByIDResponse{},
			wantErr: nil,
			prepare: func(m *Mockapplication) {
				m.EXPECT().DeleteUserByID(gomock.Any(), int(req.Id)).Return(nil).Times(1)
			},
		},
		{
			name:    "err user not found",
			req:     req,
			resp:    nil,
			wantErr: errUserNotFound,
			prepare: func(m *Mockapplication) {
				m.EXPECT().DeleteUserByID(gomock.Any(), int(req.Id)).Return(app.ErrNotFound).Times(1)
			},
		},
		{
			name:    "err any",
			req:     req,
			resp:    nil,
			wantErr: errInternal,
			prepare: func(m *Mockapplication) {
				m.EXPECT().DeleteUserByID(gomock.Any(), int(req.Id)).Return(errAny).Times(1)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, app, assert := setup(t)
			if tt.prepare != nil {
				tt.prepare(app)
			}

			resp, err := client.DeleteUserByID(ctx, tt.req)
			assert.ErrorIs(err, tt.wantErr)
			assert.True(proto.Equal(resp, tt.resp))
		})
	}
}
