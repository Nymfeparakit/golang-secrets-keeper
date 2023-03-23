package handlers

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	mock_handlers "github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers/mocks"
	"github.com/Nymfeparakit/gophkeeper/server/internal/services"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestAuthServer_SignUp(t *testing.T) {
	user := dto.User{
		Email:    "testemail@mail.com",
		Password: "123123",
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService)
		request     *auth.SignUpRequest
		expResponse *auth.SignUpResponse
	}{
		{
			name: "positive test",
			request: &auth.SignUpRequest{
				Login:    user.Email,
				Password: user.Password,
			},
			setupMocks: func(authMock *mock_handlers.MockAuthService) {
				authMock.EXPECT().Register(gomock.Any(), &user).Return(nil)
			},
			expResponse: &auth.SignUpResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authServiceMock := mock_handlers.NewMockAuthService(ctrl)
			tt.setupMocks(authServiceMock)
			authServer := NewAuthServer(authServiceMock)
			response, err := authServer.SignUp(context.Background(), tt.request)

			require.NoError(t, err)
			assert.Equal(t, tt.expResponse, response)
		})
	}
}

func TestAuthServer_Login(t *testing.T) {
	email := "testemail@mail.com"
	pwd := "123123"
	token := "token"

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService)
		request     *auth.LoginRequest
		expResponse *auth.LoginResponse
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService) {
				authMock.EXPECT().Login(context.Background(), email, pwd).Return(token, nil)
			},
			request: &auth.LoginRequest{
				Login:    email,
				Password: pwd,
			},
			expResponse: &auth.LoginResponse{
				Token: token,
			},
		},
		{
			name: "invalid credentials",
			setupMocks: func(authMock *mock_handlers.MockAuthService) {
				authMock.EXPECT().Login(context.Background(), email, pwd).Return("", services.ErrInvalidCredentials)
			},
			request: &auth.LoginRequest{
				Login:    email,
				Password: pwd,
			},
			expError: status.Error(codes.Unauthenticated, "Invalid login or password"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authServiceMock := mock_handlers.NewMockAuthService(ctrl)
			tt.setupMocks(authServiceMock)
			authServer := NewAuthServer(authServiceMock)
			response, err := authServer.Login(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}
