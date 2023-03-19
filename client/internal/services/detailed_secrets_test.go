package services

import (
	"context"
	"fmt"
	mock_services "github.com/Nymfeparakit/gophkeeper/client/internal/services/mocks"
	"github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	mock_secrets "github.com/Nymfeparakit/gophkeeper/server/proto/secrets/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	authService.EXPECT().AddAuthMetadata(gomock.Any(), credentials.Token).Return(context.Background(), nil).AnyTimes()
	userStorage.EXPECT().GetCredentials().Return(&credentials, nil).AnyTimes()
	cryptoService.EXPECT().DecryptSecret(gomock.Any()).Return(nil).AnyTimes()
	cryptoService.EXPECT().EncryptSecret(gomock.Any()).Return(nil).AnyTimes()
}

func TestUpdateRetrieveDeleteSecretService_GetPasswordByID(t *testing.T) {
	remoteLogin := "remote login"
	localLogin := "local login"
	now := time.Now().UTC()
	nowHourLater := now.Add(time.Hour * 1)
	tests := []struct {
		name           string
		remotePwd      *secrets.Password
		localSecret    *dto.BaseSecret
		expLogin       string
		setupTestMocks func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient)
	}{
		{
			name:        "Return pwd from server",
			localSecret: &dto.BaseSecret{UpdatedAt: now},
			remotePwd:   &secrets.Password{Login: remoteLogin, UpdatedAt: timestamppb.New(nowHourLater)},
			expLogin:    remoteLogin,
			setupTestMocks: func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient) {
				localStorageMock.EXPECT().UpdateCredentials(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:        "Return pwd from local storage",
			localSecret: &dto.BaseSecret{UpdatedAt: nowHourLater},
			remotePwd:   &secrets.Password{Login: remoteLogin, UpdatedAt: timestamppb.New(now)},
			expLogin:    localLogin,
			setupTestMocks: func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient) {
				clientMock.EXPECT().UpdateCredentials(gomock.Any(), gomock.Any()).Return(&secrets.EmptyResponse{}, nil)
			},
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
			tt.setupTestMocks(localStorageMock, secretsClientMock)

			pwdInstanceService := NewUpdateRetrieveDeletePasswordService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			result, err := pwdInstanceService.GetSecretByID(context.Background(), id)

			require.NoError(t, err)
			assert.Equal(t, tt.expLogin, result.Login)
		})
	}
}

func TestUpdateRetrieveDeleteSecretService_GetCardByID(t *testing.T) {
	remoteName := "remote name"
	localName := "local name"
	now := time.Now().UTC()
	nowHourLater := now.Add(time.Hour * 1)
	tests := []struct {
		name           string
		remoteSecret   *secrets.CardInfo
		localSecret    *dto.BaseSecret
		expName        string
		setupTestMocks func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient)
	}{
		{
			name:         "Return secret from server",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: now},
			remoteSecret: &secrets.CardInfo{Name: remoteName, UpdatedAt: timestamppb.New(nowHourLater)},
			expName:      remoteName,
			setupTestMocks: func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient) {
				localStorageMock.EXPECT().UpdateCardInfo(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:         "Return secret from local storage",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: nowHourLater},
			remoteSecret: &secrets.CardInfo{Name: remoteName, UpdatedAt: timestamppb.New(now)},
			expName:      localName,
			setupTestMocks: func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient) {
				clientMock.EXPECT().UpdateCardInfo(gomock.Any(), gomock.Any()).Return(&secrets.EmptyResponse{}, nil)
			},
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
			tt.setupTestMocks(localStorageMock, secretsClientMock)

			instanceService := NewUpdateRetrieveDeleteCardService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			result, err := instanceService.GetSecretByID(context.Background(), id)

			require.NoError(t, err)
			assert.Equal(t, tt.expName, result.Name)
		})
	}
}

func TestUpdateRetrieveDeleteSecretService_GetTextByID(t *testing.T) {
	remoteName := "remote name"
	localName := "local name"
	now := time.Now().UTC()
	nowHourLater := now.Add(time.Hour * 1)
	tests := []struct {
		name           string
		remoteSecret   *secrets.TextInfo
		localSecret    *dto.BaseSecret
		expName        string
		setupTestMocks func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient)
	}{
		{
			name:         "Return secret from server",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: now},
			remoteSecret: &secrets.TextInfo{Name: remoteName, UpdatedAt: timestamppb.New(nowHourLater)},
			expName:      remoteName,
			setupTestMocks: func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient) {
				localStorageMock.EXPECT().UpdateTextInfo(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:         "Return secret from local storage",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: nowHourLater},
			remoteSecret: &secrets.TextInfo{Name: remoteName, UpdatedAt: timestamppb.New(now)},
			expName:      localName,
			setupTestMocks: func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient) {
				clientMock.EXPECT().UpdateTextInfo(gomock.Any(), gomock.Any()).Return(&secrets.EmptyResponse{}, nil)
			},
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
			tt.setupTestMocks(localStorageMock, secretsClientMock)

			instanceService := NewUpdateRetrieveDeleteTextService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			result, err := instanceService.GetSecretByID(context.Background(), id)

			require.NoError(t, err)
			assert.Equal(t, tt.expName, result.Name)
		})
	}
}

func TestUpdateRetrieveDeleteSecretService_GetBinaryByID(t *testing.T) {
	remoteName := "remote name"
	localName := "local name"
	tests := []struct {
		name           string
		remoteSecret   *secrets.BinaryInfo
		localSecret    *dto.BaseSecret
		expName        string
		setupTestMocks func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient)
	}{
		{
			name:         "Return secret from server",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: time.Now()},
			remoteSecret: &secrets.BinaryInfo{Name: remoteName, UpdatedAt: timestamppb.New(time.Now().Add(time.Hour * 1))},
			expName:      remoteName,
			setupTestMocks: func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient) {
				localStorageMock.EXPECT().UpdateBinaryInfo(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:         "Return secret from local storage",
			localSecret:  &dto.BaseSecret{Name: localName, UpdatedAt: time.Now().Add(time.Hour * 1)},
			remoteSecret: &secrets.BinaryInfo{Name: remoteName, UpdatedAt: timestamppb.New(time.Now())},
			expName:      localName,
			setupTestMocks: func(localStorageMock *mock_services.MockLocalSecretsStorage, clientMock *mock_secrets.MockSecretsManagementClient) {
				clientMock.EXPECT().UpdateBinaryInfo(gomock.Any(), gomock.Any()).Return(&secrets.EmptyResponse{}, nil)
			},
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
			tt.setupTestMocks(localStorageMock, secretsClientMock)

			instanceService := NewUpdateRetrieveDeleteBinaryService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			result, err := instanceService.GetSecretByID(context.Background(), id)

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
	err := pwdInstanceService.UpdateSecret(context.Background(), pwd)

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
	err := instanceService.UpdateSecret(context.Background(), secret)

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
	err := instanceService.UpdateSecret(context.Background(), secret)

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
	err := instanceService.UpdateSecret(context.Background(), secret)

	require.NoError(t, err)
}

func deleteSetupMocks(
	authService *mock_services.MockAuthMetadataService,
	userStorage *mock_services.MockUserCredentialsStorage,
) {
	credentials := dto.UserCredentials{Email: "email", Token: "token"}
	authService.EXPECT().AddAuthMetadata(gomock.Any(), credentials.Token).Return(context.Background(), nil)
	userStorage.EXPECT().GetCredentials().Return(&credentials, nil).AnyTimes()
}

type deleteTestCase struct {
	name      string
	localErr  error
	remoteErr error
	expErr    error
}

func getDeleteTestCases() []deleteTestCase {
	tests := []deleteTestCase{
		{
			name: "Simple test",
		},
		{
			name:     "Secret was already deleted locally",
			localErr: storage.ErrSecretDoesNotExistOrWasDeleted,
		},
		{
			name:      "Secret was already deleted on remote",
			remoteErr: status.Error(codes.FailedPrecondition, ""),
		},
		{
			name:      "Secret was already deleted on remote and locally",
			localErr:  storage.ErrSecretDoesNotExistOrWasDeleted,
			remoteErr: status.Error(codes.FailedPrecondition, ""),
			expErr:    fmt.Errorf("secret with provided id doesn't exist"),
		},
	}

	return tests
}

func TestUpdateRetrieveDeleteSecretService_DeletePassword(t *testing.T) {
	tests := getDeleteTestCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			id := "123"

			authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
			userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
			deleteSetupMocks(authServiceMock, userCredsMock)

			secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
			expRequest := secrets.DeleteSecretRequest{Id: id}
			response := &secrets.ResponseWithError{}
			secretsClientMock.EXPECT().DeleteCredentials(gomock.Any(), &expRequest).Return(response, tt.remoteErr)

			itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)

			localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
			localStorageMock.EXPECT().DeleteCredentials(gomock.Any(), id).Return(tt.localErr)

			pwdInstanceService := NewUpdateRetrieveDeletePasswordService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			err := pwdInstanceService.DeleteSecret(context.Background(), id)

			if tt.expErr != nil {
				assert.Equal(t, tt.expErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUpdateRetrieveDeleteSecretService_DeleteCardInfo(t *testing.T) {
	tests := getDeleteTestCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			id := "123"

			authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
			userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
			deleteSetupMocks(authServiceMock, userCredsMock)

			secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
			expRequest := secrets.DeleteSecretRequest{Id: id}
			response := &secrets.ResponseWithError{}
			secretsClientMock.EXPECT().DeleteCardInfo(gomock.Any(), &expRequest).Return(response, tt.remoteErr)

			itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)

			localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
			localStorageMock.EXPECT().DeleteCardInfo(gomock.Any(), id).Return(tt.localErr)

			instanceService := NewUpdateRetrieveDeleteCardService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			err := instanceService.DeleteSecret(context.Background(), id)

			if tt.expErr != nil {
				assert.Equal(t, tt.expErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUpdateRetrieveDeleteSecretService_DeleteBinaryInfo(t *testing.T) {
	tests := getDeleteTestCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			id := "123"

			authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
			userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
			deleteSetupMocks(authServiceMock, userCredsMock)

			secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
			expRequest := secrets.DeleteSecretRequest{Id: id}
			response := &secrets.ResponseWithError{}
			secretsClientMock.EXPECT().DeleteBinaryInfo(gomock.Any(), &expRequest).Return(response, tt.remoteErr)

			itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)

			localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
			localStorageMock.EXPECT().DeleteBinaryInfo(gomock.Any(), id).Return(tt.localErr)

			instanceService := NewUpdateRetrieveDeleteBinaryService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			err := instanceService.DeleteSecret(context.Background(), id)

			if tt.expErr != nil {
				assert.Equal(t, tt.expErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUpdateRetrieveDeleteSecretService_DeleteTextInfo(t *testing.T) {
	tests := getDeleteTestCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			id := "123"

			authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
			userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
			deleteSetupMocks(authServiceMock, userCredsMock)

			secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
			expRequest := secrets.DeleteSecretRequest{Id: id}
			response := &secrets.ResponseWithError{}
			secretsClientMock.EXPECT().DeleteTextInfo(gomock.Any(), &expRequest).Return(response, tt.remoteErr)

			itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)

			localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
			localStorageMock.EXPECT().DeleteTextInfo(gomock.Any(), id).Return(tt.localErr)

			instanceService := NewUpdateRetrieveDeleteTextService(
				authServiceMock,
				itemCryptoMock,
				userCredsMock,
				localStorageMock,
				secretsClientMock,
			)
			err := instanceService.DeleteSecret(context.Background(), id)

			if tt.expErr != nil {
				assert.Equal(t, tt.expErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
