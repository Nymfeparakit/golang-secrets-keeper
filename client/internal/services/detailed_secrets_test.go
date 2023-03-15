package services

import (
	"context"
	mock_services "github.com/Nymfeparakit/gophkeeper/client/internal/services/mocks"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	mock_secrets "github.com/Nymfeparakit/gophkeeper/server/proto/secrets/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func getSetupMocks(
	authService *mock_services.MockAuthMetadataService,
	userStorage *mock_services.MockUserCredentialsStorage,
	cryptoService *mock_services.MockSecretCryptoService,
) {
	credentials := dto.UserCredentials{Email: "email", Token: "token"}
	authService.EXPECT().AddAuthMetadata(gomock.Any(), credentials.Token).Return(context.Background(), nil)
	userStorage.EXPECT().GetCredentials().Return(&credentials, nil).AnyTimes()
	cryptoService.EXPECT().DecryptSecret(gomock.Any()).Return(nil).AnyTimes()
}

func TestUpdateRetrieveDeleteSecretService_GetPasswordByID(t *testing.T) {
	remoteLogin := "remote login"
	localLogin := "local login"
	tests := []struct {
		name        string
		remotePwd   *secrets.Password
		localSecret *dto.BaseSecret
		expLogin    string
	}{
		{
			name:        "Return pwd from server",
			localSecret: &dto.BaseSecret{UpdatedAt: time.Now()},
			remotePwd:   &secrets.Password{Login: remoteLogin, UpdatedAt: timestamppb.New(time.Now().Add(time.Hour * 1))},
			expLogin:    remoteLogin,
		},
		{
			name:        "Return pwd from local storage",
			localSecret: &dto.BaseSecret{UpdatedAt: time.Now().Add(time.Hour * 1)},
			remotePwd:   &secrets.Password{Login: remoteLogin, UpdatedAt: timestamppb.New(time.Now())},
			expLogin:    localLogin,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			localPwd := dto.LoginPassword{Login: localLogin, BaseSecret: *tt.localSecret}
			id := "123"

			authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
			userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
			secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
			expRequest := secrets.GetSecretRequest{Id: id}
			response := &secrets.GetCredentialsResponse{Password: tt.remotePwd}
			secretsClientMock.EXPECT().GetCredentialsByID(gomock.Any(), &expRequest).Return(response, nil)
			itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
			localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
			localStorageMock.EXPECT().GetCredentialsById(gomock.Any(), id, "email").Return(localPwd, nil)
			getSetupMocks(authServiceMock, userCredsMock, itemCryptoMock)

			pwdInstanceService := NewUpdateRetrieveDeletePasswordService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			result, err := pwdInstanceService.GetSecretByID(id)

			require.NoError(t, err)
			assert.Equal(t, tt.expLogin, result.Login)
		})
	}
}

func TestUpdateRetrieveDeleteSecretService_GetCardByID(t *testing.T) {
	remoteName := "remote name"
	localName := "local name"
	tests := []struct {
		name         string
		remoteSecret *secrets.CardInfo
		localSecret  *dto.BaseSecret
		expName      string
	}{
		{
			name:         "Return secret from server",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: time.Now()},
			remoteSecret: &secrets.CardInfo{Name: remoteName, UpdatedAt: timestamppb.New(time.Now().Add(time.Hour * 1))},
			expName:      remoteName,
		},
		{
			name:         "Return secret from local storage",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: time.Now().Add(time.Hour * 1)},
			remoteSecret: &secrets.CardInfo{Name: remoteName, UpdatedAt: timestamppb.New(time.Now())},
			expName:      localName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			localCard := dto.CardInfo{BaseSecret: *tt.localSecret}
			id := "123"

			authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
			userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
			secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
			expRequest := secrets.GetSecretRequest{Id: id}
			response := &secrets.GetCardResponse{Card: tt.remoteSecret}
			secretsClientMock.EXPECT().GetCardByID(gomock.Any(), &expRequest).Return(response, nil)
			itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
			localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
			localStorageMock.EXPECT().GetCardById(gomock.Any(), id, "email").Return(localCard, nil)
			getSetupMocks(authServiceMock, userCredsMock, itemCryptoMock)

			instanceService := NewUpdateRetrieveDeleteCardService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			result, err := instanceService.GetSecretByID(id)

			require.NoError(t, err)
			assert.Equal(t, tt.expName, result.Name)
		})
	}
}

func TestUpdateRetrieveDeleteSecretService_GetTextByID(t *testing.T) {
	remoteName := "remote name"
	localName := "local name"
	tests := []struct {
		name         string
		remoteSecret *secrets.TextInfo
		localSecret  *dto.BaseSecret
		expName      string
	}{
		{
			name:         "Return secret from server",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: time.Now()},
			remoteSecret: &secrets.TextInfo{Name: remoteName, UpdatedAt: timestamppb.New(time.Now().Add(time.Hour * 1))},
			expName:      remoteName,
		},
		{
			name:         "Return secret from local storage",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: time.Now().Add(time.Hour * 1)},
			remoteSecret: &secrets.TextInfo{Name: remoteName, UpdatedAt: timestamppb.New(time.Now())},
			expName:      localName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			localCard := dto.TextInfo{BaseSecret: *tt.localSecret}
			id := "123"

			authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
			userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
			secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
			expRequest := secrets.GetSecretRequest{Id: id}
			response := &secrets.GetTextResponse{Text: tt.remoteSecret}
			secretsClientMock.EXPECT().GetTextByID(gomock.Any(), &expRequest).Return(response, nil)
			itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
			localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
			localStorageMock.EXPECT().GetTextById(gomock.Any(), id, "email").Return(localCard, nil)
			getSetupMocks(authServiceMock, userCredsMock, itemCryptoMock)

			instanceService := NewUpdateRetrieveDeleteTextService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			result, err := instanceService.GetSecretByID(id)

			require.NoError(t, err)
			assert.Equal(t, tt.expName, result.Name)
		})
	}
}

func TestUpdateRetrieveDeleteSecretService_GetBinaryByID(t *testing.T) {
	remoteName := "remote name"
	localName := "local name"
	tests := []struct {
		name         string
		remoteSecret *secrets.BinaryInfo
		localSecret  *dto.BaseSecret
		expName      string
	}{
		{
			name:         "Return secret from server",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: time.Now()},
			remoteSecret: &secrets.BinaryInfo{Name: remoteName, UpdatedAt: timestamppb.New(time.Now().Add(time.Hour * 1))},
			expName:      remoteName,
		},
		{
			name:         "Return secret from local storage",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: time.Now().Add(time.Hour * 1)},
			remoteSecret: &secrets.BinaryInfo{Name: remoteName, UpdatedAt: timestamppb.New(time.Now())},
			expName:      localName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			localBin := dto.BinaryInfo{BaseSecret: *tt.localSecret}
			id := "123"

			authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
			userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
			secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
			expRequest := secrets.GetSecretRequest{Id: id}
			response := &secrets.GetBinaryResponse{Bin: tt.remoteSecret}
			secretsClientMock.EXPECT().GetBinaryByID(gomock.Any(), &expRequest).Return(response, nil)
			itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
			localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
			localStorageMock.EXPECT().GetBinaryById(gomock.Any(), id, "email").Return(localBin, nil)
			getSetupMocks(authServiceMock, userCredsMock, itemCryptoMock)

			instanceService := NewUpdateRetrieveDeleteBinaryService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			result, err := instanceService.GetSecretByID(id)

			require.NoError(t, err)
			assert.Equal(t, tt.expName, result.Name)
		})
	}
}

func updateSetupMocks(
	authService *mock_services.MockAuthMetadataService,
	userStorage *mock_services.MockUserCredentialsStorage,
	cryptoService *mock_services.MockSecretCryptoService,
) {
	credentials := dto.UserCredentials{Email: "email", Token: "token"}
	authService.EXPECT().AddAuthMetadata(gomock.Any(), credentials.Token).Return(context.Background(), nil)
	userStorage.EXPECT().GetCredentials().Return(&credentials, nil).AnyTimes()
	cryptoService.EXPECT().EncryptSecret(gomock.Any()).Return(nil).AnyTimes()
}

func TestUpdateRetrieveDeleteSecretService_UpdatePassword(t *testing.T) {
	pwd := dto.LoginPassword{Login: "login", Password: "password"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
	itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
	secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	secretsClientMock.EXPECT().UpdateCredentials(gomock.Any(), gomock.Any()).Return(&secrets.EmptyResponse{}, nil)
	localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
	localStorageMock.EXPECT().UpdateCredentials(gomock.Any(), gomock.Any())
	updateSetupMocks(authServiceMock, userCredsMock, itemCryptoMock)

	pwdInstanceService := NewUpdateRetrieveDeletePasswordService(
		authServiceMock,
		itemCryptoMock,
		userCredsMock,
		localStorageMock,
		secretsClientMock,
	)
	err := pwdInstanceService.UpdateSecret(pwd)

	require.NoError(t, err)
}

func TestUpdateRetrieveDeleteSecretService_UpdateText(t *testing.T) {
	secret := dto.TextInfo{Text: "some text"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
	itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
	secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	secretsClientMock.EXPECT().UpdateTextInfo(gomock.Any(), gomock.Any()).Return(&secrets.EmptyResponse{}, nil)
	localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
	localStorageMock.EXPECT().UpdateTextInfo(gomock.Any(), gomock.Any())
	updateSetupMocks(authServiceMock, userCredsMock, itemCryptoMock)

	instanceService := NewUpdateRetrieveDeleteTextService(
		authServiceMock,
		itemCryptoMock,
		userCredsMock,
		localStorageMock,
		secretsClientMock,
	)
	err := instanceService.UpdateSecret(secret)

	require.NoError(t, err)
}

func TestUpdateRetrieveDeleteSecretService_UpdateCard(t *testing.T) {
	secret := dto.CardInfo{Number: "123456789"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
	itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
	secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	secretsClientMock.EXPECT().UpdateCardInfo(gomock.Any(), gomock.Any()).Return(&secrets.EmptyResponse{}, nil)
	localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
	localStorageMock.EXPECT().UpdateCardInfo(gomock.Any(), gomock.Any())
	updateSetupMocks(authServiceMock, userCredsMock, itemCryptoMock)

	instanceService := NewUpdateRetrieveDeleteCardService(
		authServiceMock,
		itemCryptoMock,
		userCredsMock,
		localStorageMock,
		secretsClientMock,
	)
	err := instanceService.UpdateSecret(secret)

	require.NoError(t, err)
}

func TestUpdateRetrieveDeleteSecretService_UpdateBinary(t *testing.T) {
	secret := dto.BinaryInfo{Data: "some data"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
	itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
	secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	secretsClientMock.EXPECT().UpdateBinaryInfo(gomock.Any(), gomock.Any()).Return(&secrets.EmptyResponse{}, nil)
	localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
	localStorageMock.EXPECT().UpdateBinaryInfo(gomock.Any(), gomock.Any())
	updateSetupMocks(authServiceMock, userCredsMock, itemCryptoMock)

	instanceService := NewUpdateRetrieveDeleteBinaryService(
		authServiceMock,
		itemCryptoMock,
		userCredsMock,
		localStorageMock,
		secretsClientMock,
	)
	err := instanceService.UpdateSecret(secret)

	require.NoError(t, err)
}
