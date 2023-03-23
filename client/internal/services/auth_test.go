package services

import (
	"context"
	mock_services "github.com/Nymfeparakit/gophkeeper/client/internal/services/mocks"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	mock_auth "github.com/Nymfeparakit/gophkeeper/server/proto/auth/mocks"
	"github.com/golang/mock/gomock"
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

	credsStorageMock := mock_services.NewMockCredentialsStorage(ctrl)
	userCryptoMock := mock_services.NewMockUserCryptoService(ctrl)
	userStorageMock := mock_services.NewMockUsersStorage(ctrl)
	secretLoadMock := mock_services.NewMockSecretsLoadingService(ctrl)
	authService := NewAuthService(authClientMock, credsStorageMock, userCryptoMock, userStorageMock, secretLoadMock)

	err := authService.Register(context.Background(), user)

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

	credsStorageMock := mock_services.NewMockCredentialsStorage(ctrl)
	userCryptoMock := mock_services.NewMockUserCryptoService(ctrl)
	userStorageMock := mock_services.NewMockUsersStorage(ctrl)
	secretLoadMock := mock_services.NewMockSecretsLoadingService(ctrl)
	authService := NewAuthService(authClientMock, credsStorageMock, userCryptoMock, userStorageMock, secretLoadMock)

	credsStorageMock.EXPECT().SaveCredentials(email, token).Return(nil)
	userCryptoMock.EXPECT().CreateUserKey(pwd).Return(nil)
	userStorageMock.EXPECT().CreateUser(gomock.Any(), email).Return(nil)
	secretLoadMock.EXPECT().LoadSecrets(gomock.Any()).Return(nil)

	err := authService.Login(context.Background(), email, pwd)

	require.NoError(t, err)
}
