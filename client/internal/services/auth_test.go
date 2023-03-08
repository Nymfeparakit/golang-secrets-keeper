package services

import (
	mock_services "github.com/Nymfeparakit/gophkeeper/client/internal/services/mocks"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	mock_auth "github.com/Nymfeparakit/gophkeeper/server/proto/auth/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuthService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClientMock := mock_auth.NewMockAuthManagementClient(ctrl)
	user := &dto.User{
		Email:    "test@mail.com",
		Password: "123123",
	}
	expectedRequest := auth.SignUpRequest{
		Login:    user.Email,
		Password: user.Password,
	}
	response := auth.SignUpResponse{}
	authClientMock.EXPECT().SignUp(gomock.Any(), &expectedRequest).Return(&response, nil)

	tokenStorageMock := mock_services.NewMockTokenStorage(ctrl)
	userCryptoMock := mock_services.NewMockUserCryptoService(ctrl)
	authService := NewAuthService(authClientMock, tokenStorageMock, userCryptoMock)

	err := authService.Register(user)

	require.NoError(t, err)
}

func TestAuthService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authClientMock := mock_auth.NewMockAuthManagementClient(ctrl)
	email := "test@mail.com"
	pwd := "123"
	token := "token"
	expectedRequest := auth.LoginRequest{
		Login:    email,
		Password: pwd,
	}
	resp := auth.LoginResponse{Token: token}
	authClientMock.EXPECT().Login(gomock.Any(), &expectedRequest).Return(&resp, nil)

	tokenStorageMock := mock_services.NewMockTokenStorage(ctrl)
	tokenStorageMock.EXPECT().SaveToken(token).Return(nil)
	userCryptoMock := mock_services.NewMockUserCryptoService(ctrl)
	userCryptoMock.EXPECT().CreateUserKey(pwd).Return(nil)

	authService := NewAuthService(authClientMock, tokenStorageMock, userCryptoMock)

	err := authService.Login(email, pwd)

	require.NoError(t, err)
}

func TestAuthService_getUserToken(t *testing.T) {
	tests := []struct {
		name              string
		setupMocks        func(tokenMock *mock_services.MockTokenStorage)
		initialTokenValue string
	}{
		{
			name: "token is empty",
			setupMocks: func(tokenMock *mock_services.MockTokenStorage) {
				tokenMock.EXPECT().GetToken().Return("token", nil)
			},
			initialTokenValue: "",
		},
		{
			name:              "token is not empty",
			initialTokenValue: "token",
			setupMocks:        func(tokenMock *mock_services.MockTokenStorage) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authClientMock := mock_auth.NewMockAuthManagementClient(ctrl)
			tokenStorageMock := mock_services.NewMockTokenStorage(ctrl)
			tt.setupMocks(tokenStorageMock)
			userCryptoMock := mock_services.NewMockUserCryptoService(ctrl)
			authService := NewAuthService(authClientMock, tokenStorageMock, userCryptoMock)
			authService.userToken = tt.initialTokenValue

			actualResult, err := authService.getUserToken()

			require.NoError(t, err)
			assert.Equal(t, "token", actualResult)
		})
	}
}
