package app_test

import (
	"testing"

	"rest-on-grpc-gateway/modules/user/internal/app"
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

func TestApp_CreateUser(t *testing.T) {
	ctx := setupCtx(t)

	req := &domain.User{
		Name:     "user",
		Email:    "user@mail.com",
		Password: "12345678",
	}

	tests := []struct {
		name    string
		req     *domain.User
		want    *domain.User
		wantErr error
		prepare func(m *mocks)
	}{
		{
			name:    "success",
			req:     req,
			want:    user,
			wantErr: nil,
			prepare: func(m *mocks) {
				m.repo.EXPECT().CreateUser(ctx, req).Return(user, nil).Times(1)
			},
		},
		{
			name:    "err any",
			req:     req,
			want:    nil,
			wantErr: errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().CreateUser(ctx, req).Return(nil, errAny).Times(1)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			app, mocks, assert := setup(t)
			if tt.prepare != nil {
				tt.prepare(mocks)
			}

			u, err := app.CreateUser(ctx, tt.req)
			assert.ErrorIs(err, tt.wantErr)
			assert.Equal(tt.want, u)
		})
	}
}

func TestApp_GetUserByID(t *testing.T) {
	ctx := setupCtx(t)

	tests := []struct {
		name    string
		req     int
		want    *domain.User
		wantErr error
		prepare func(m *mocks)
	}{
		{
			name:    "success",
			req:     userID,
			want:    user,
			wantErr: nil,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(user, nil).Times(1)
			},
		},
		{
			name:    "err any",
			req:     userID,
			want:    nil,
			wantErr: errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(nil, errAny).Times(1)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			app, mocks, assert := setup(t)
			if tt.prepare != nil {
				tt.prepare(mocks)
			}

			u, err := app.GetUserByID(ctx, tt.req)
			assert.ErrorIs(err, tt.wantErr)
			assert.Equal(tt.want, u)
		})
	}
}

func TestApp_UpdateUserByID(t *testing.T) {
	ctx := setupCtx(t)

	userName, email := "user name", "email@email.com"
	updateUser := &domain.User{
		ID:       userID,
		Name:     "user",
		Email:    "user@mail.com",
		Password: "",
	}

	tests := []struct {
		name            string
		userID          int
		userName, email string
		want            *domain.User
		wantErr         error
		prepare         func(m *mocks)
	}{
		{
			name:     "success",
			userID:   userID,
			userName: userName,
			email:    email,
			want: &domain.User{
				ID:       userID,
				Name:     userName,
				Email:    email,
				Password: "",
			},
			wantErr: nil,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(user, nil).Times(1)

				updateUser.Name = userName
				updateUser.Email = email
				m.repo.EXPECT().UpdateUserByID(ctx, userID, userName, email).Return(updateUser, nil).Times(1)
			},
		},
		{
			name:     "err any on GetUserByID",
			userID:   userID,
			userName: userName,
			email:    email,
			want:     nil,
			wantErr:  errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(nil, errAny).Times(1)
			},
		},
		{
			name:     "err any on UpdateUserByID",
			userID:   userID,
			userName: userName,
			email:    email,
			want:     nil,
			wantErr:  errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(user, nil).Times(1)
				m.repo.EXPECT().UpdateUserByID(ctx, userID, userName, email).Return(nil, errAny).Times(1)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			app, mocks, assert := setup(t)
			if tt.prepare != nil {
				tt.prepare(mocks)
			}

			newUser, err := app.UpdateUserByID(ctx, tt.userID, tt.userName, tt.email)
			assert.ErrorIs(err, tt.wantErr)
			assert.Equal(tt.want, newUser)
		})
	}
}

func TestApp_UpdateUserPasswordByID(t *testing.T) {
	ctx := setupCtx(t)

	oldPass, newPass := "12345678", "87654321"
	updateUser := &domain.User{
		ID:       userID,
		Name:     "user",
		Email:    "user@mail.com",
		Password: "12345678",
	}

	tests := []struct {
		name             string
		userID           int
		oldPass, newPass string
		wantErr          error
		prepare          func(m *mocks)
	}{
		{
			name:    "success",
			userID:  userID,
			oldPass: oldPass,
			newPass: newPass,
			wantErr: nil,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(updateUser, nil).Times(1)
				m.repo.EXPECT().UpdateUserPasswordByID(ctx, userID, newPass).Return(nil).Times(1)
			},
		},
		{
			name:    "err invalid password",
			userID:  userID,
			oldPass: "12344321",
			newPass: newPass,
			wantErr: app.ErrInvalidPassword,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(user, nil).Times(1)
			},
		},
		{
			name:    "err any on GetUserByID",
			userID:  userID,
			oldPass: oldPass,
			newPass: newPass,
			wantErr: errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(nil, errAny).Times(1)
			},
		},
		{
			name:    "err any on UpdateUserPasswordByID",
			userID:  userID,
			oldPass: oldPass,
			newPass: newPass,
			wantErr: errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(updateUser, nil).Times(1)
				updateUser.Password = oldPass
				m.repo.EXPECT().UpdateUserPasswordByID(ctx, userID, newPass).Return(errAny).Times(1)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			app, mocks, assert := setup(t)
			if tt.prepare != nil {
				tt.prepare(mocks)
			}

			err := app.UpdateUserPasswordByID(ctx, tt.userID, tt.oldPass, tt.newPass)
			assert.ErrorIs(err, tt.wantErr)
		})
	}
}

func TestApp_DeleteUserByID(t *testing.T) {
	ctx := setupCtx(t)

	tests := []struct {
		name    string
		req     int
		wantErr error
		prepare func(m *mocks)
	}{
		{
			name:    "success",
			req:     userID,
			wantErr: nil,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(user, nil).Times(1)
				m.repo.EXPECT().DeleteUserByID(ctx, userID).Return(nil).Times(1)
			},
		},
		{
			name:    "err any on GetUserByID",
			req:     userID,
			wantErr: errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(nil, errAny).Times(1)
			},
		},
		{
			name:    "err any on DeleteUserByID",
			req:     userID,
			wantErr: errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(user, nil).Times(1)
				m.repo.EXPECT().DeleteUserByID(ctx, userID).Return(errAny).Times(1)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			app, mocks, assert := setup(t)
			if tt.prepare != nil {
				tt.prepare(mocks)
			}

			err := app.DeleteUserByID(ctx, tt.req)
			assert.ErrorIs(err, tt.wantErr)
		})
	}
}
