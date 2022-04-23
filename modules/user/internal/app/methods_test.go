package app_test

import (
	"testing"

	"rest-on-grpc-gateway/modules/user/internal/app"
	"rest-on-grpc-gateway/modules/user/internal/domain"
)

func TestApp_CreateUser(t *testing.T) {
	ctx := setupCtx(t)
	newUser := domain.User{
		Name:         userName,
		Email:        email,
		PasswordHash: []byte(password),
	}

	tests := []struct {
		name                  string
		userName, email, pass string
		want                  *domain.User
		wantErr               error
		prepare               func(m *mocks)
	}{
		{
			name:     "success",
			userName: userName,
			email:    email,
			pass:     password,
			want:     user,
			wantErr:  nil,
			prepare: func(m *mocks) {
				m.hash.EXPECT().Hashing(password).Return([]byte(password), nil).Times(1)
				m.repo.EXPECT().CreateUser(ctx, &newUser).Return(user, nil).Times(1)
			},
		},
		{
			name:     "err any hashing password",
			userName: userName,
			email:    email,
			pass:     password,
			want:     nil,
			wantErr:  errAny,
			prepare: func(m *mocks) {
				m.hash.EXPECT().Hashing(password).Return(nil, errAny).Times(1)
			},
		},
		{
			name:     "err email exist",
			userName: userName,
			email:    email,
			pass:     password,
			want:     nil,
			wantErr:  app.ErrEmailExist,
			prepare: func(m *mocks) {
				m.hash.EXPECT().Hashing(password).Return([]byte(password), nil).Times(1)
				m.repo.EXPECT().CreateUser(ctx, newUser).Return(nil, app.ErrEmailExist).Times(1)
			},
		},
		{
			name:     "err any",
			userName: userName,
			email:    email,
			pass:     password,
			want:     nil,
			wantErr:  errAny,
			prepare: func(m *mocks) {
				m.hash.EXPECT().Hashing(password).Return([]byte(password), nil).Times(1)
				m.repo.EXPECT().CreateUser(ctx, newUser).Return(nil, errAny).Times(1)
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

			u, err := app.CreateUser(ctx, tt.userName, tt.email, tt.pass)
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
			name:    "err not found",
			req:     userID,
			want:    nil,
			wantErr: app.ErrNotFound,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(nil, app.ErrNotFound).Times(1)
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

	newName, newEmail := "user name", "email@email.com"
	updateUser := &domain.User{
		ID:    userID,
		Name:  "user",
		Email: "user@mail.com",
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
			userName: newName,
			email:    newEmail,
			want: &domain.User{
				ID:    userID,
				Name:  newName,
				Email: newEmail,
			},
			wantErr: nil,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(user, nil).Times(1)

				updateUser.Name = newName
				updateUser.Email = newEmail
				m.repo.EXPECT().UpdateUserByID(ctx, userID, newName, newEmail).Return(updateUser, nil).Times(1)
			},
		},
		{
			name:     "err not found",
			userID:   userID,
			userName: newName,
			email:    newEmail,
			want:     nil,
			wantErr:  app.ErrNotFound,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(nil, app.ErrNotFound).Times(1)
			},
		},
		{
			name:     "err any on GetUserByID",
			userID:   userID,
			userName: newName,
			email:    newEmail,
			want:     nil,
			wantErr:  errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(nil, errAny).Times(1)
			},
		},
		{
			name:     "err any on UpdateUserByID",
			userID:   userID,
			userName: newName,
			email:    newEmail,
			want:     nil,
			wantErr:  errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(user, nil).Times(1)
				m.repo.EXPECT().UpdateUserByID(ctx, userID, newName, newEmail).Return(nil, errAny).Times(1)
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

	var (
		oldPass    = "12345678"
		newPass    = "87654321"
		errOldPass = "12344321"
	)
	updateUser := &domain.User{
		ID:           userID,
		Name:         "user",
		Email:        "user@mail.com",
		PasswordHash: []byte("12345678"),
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
				m.hash.EXPECT().Compare(updateUser.PasswordHash, []byte(oldPass)).Return(true).Times(1)
				m.hash.EXPECT().Compare(updateUser.PasswordHash, []byte(newPass)).Return(false).Times(1)
				m.hash.EXPECT().Hashing(newPass).Return([]byte(newPass), nil).Times(1)
				m.repo.EXPECT().UpdateUserPasswordByID(ctx, userID, []byte(newPass)).Return(nil).Times(1)
			},
		},
		{
			name:    "err not found",
			userID:  userID,
			oldPass: oldPass,
			newPass: newPass,
			wantErr: app.ErrNotFound,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(nil, app.ErrNotFound).Times(1)
			},
		},
		{
			name:    "err invalid password",
			userID:  userID,
			oldPass: errOldPass,
			newPass: newPass,
			wantErr: app.ErrInvalidPassword,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(updateUser, nil).Times(1)
				m.hash.EXPECT().Compare(updateUser.PasswordHash, []byte(errOldPass)).Return(false).Times(1)
			},
		},
		{
			name:    "err values must be different",
			userID:  userID,
			oldPass: oldPass,
			newPass: oldPass,
			wantErr: app.ErrMustDifferent,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(updateUser, nil).Times(1)
				m.hash.EXPECT().Compare(updateUser.PasswordHash, []byte(oldPass)).Return(true).Times(1)
				m.hash.EXPECT().Compare(updateUser.PasswordHash, []byte(oldPass)).Return(true).Times(1)
			},
		},
		{
			name:    "err any hashing new pass",
			userID:  userID,
			oldPass: oldPass,
			newPass: newPass,
			wantErr: errAny,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(updateUser, nil).Times(1)
				m.hash.EXPECT().Compare(updateUser.PasswordHash, []byte(oldPass)).Return(true).Times(1)
				m.hash.EXPECT().Compare(updateUser.PasswordHash, []byte(newPass)).Return(false).Times(1)
				m.hash.EXPECT().Hashing(newPass).Return(nil, errAny).Times(1)
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
				m.hash.EXPECT().Compare(updateUser.PasswordHash, []byte(oldPass)).Return(true).Times(1)
				m.hash.EXPECT().Compare(updateUser.PasswordHash, []byte(newPass)).Return(false).Times(1)
				m.hash.EXPECT().Hashing(newPass).Return([]byte(newPass), nil).Times(1)
				m.repo.EXPECT().UpdateUserPasswordByID(ctx, userID, []byte(newPass)).Return(errAny).Times(1)
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
			name:    "err not found",
			req:     userID,
			wantErr: app.ErrNotFound,
			prepare: func(m *mocks) {
				m.repo.EXPECT().GetUserByID(ctx, userID).Return(nil, app.ErrNotFound).Times(1)
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
