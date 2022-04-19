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
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

func TestApi_CreateUser(t *testing.T) {
	var (
		//errInternal                = status.Error(codes.Internal, errAny.Error())
		errInvalidArgumentUsername = status.Error(codes.InvalidArgument, "invalid CreateUserRequest.Name: value length must be between 2 and 40 runes, inclusive")
		errInvalidArgumentEmail    = status.Error(codes.InvalidArgument, "invalid CreateUserRequest.Email: value must be a valid email address | caused by: mail: missing '@' or angle-addr")
		errInvalidArgumentEmailLen = status.Error(codes.InvalidArgument, "invalid CreateUserRequest.Email: value length must be at most 50 runes")
		errInvalidArgumentPassword = status.Error(codes.InvalidArgument, "invalid CreateUserRequest.Password: value length must be between 8 and 100 runes, inclusive")
	)

	req := &userpb.CreateUserRequest{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "12345678",
	}
	resp := &userpb.CreateUserResponse{
		Id:       int64(userID),
		Name:     "user",
		Email:    "user@mail.com",
		Password: "12345678",
	}
	newUser := &domain.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "12345678",
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
				m.EXPECT().CreateUser(gomock.Any(), newUser).Return(user, nil).Times(1)
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
		//{
		//	name:    "err any",
		//	req:     req,
		//	resp:    &userpb.CreateUserResponse{},
		//	wantErr: errInternal,
		//	prepare: func(m *Mockapplication) {
		//		m.EXPECT().CreateUser(gomock.Any(), newUser).Return(nil, errAny).Times(1)
		//	},
		//},
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
