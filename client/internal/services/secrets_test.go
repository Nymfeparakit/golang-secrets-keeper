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

func TestSecretsService_AddCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	item := dto.BaseSecret{Name: "pwd", UpdatedAt: time.Now().UTC()}
	pwd := dto.LoginPassword{
		BaseSecret: item,
		Login:      "login",
		Password:   "pwd",
	}
	request := secrets.Password{
		Name:      pwd.Name,
		Login:     pwd.Login,
		Password:  pwd.Password,
		UpdatedAt: timestamppb.New(item.UpdatedAt),
	}
	credentials := dto.UserCredentials{Email: "email", Token: "token"}

	userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
	userCredsMock.EXPECT().GetCredentials().Return(&credentials, nil)
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	authServiceMock.EXPECT().AddAuthMetadata(gomock.Any(), credentials.Token).Return(context.Background(), nil)
	itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
	itemCryptoMock.EXPECT().EncryptSecret(&pwd).Return(nil)
	secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	response := secrets.AddResponse{Id: "123"}
	secretsClientMock.EXPECT().AddCredentials(gomock.Any(), &request).Return(&response, nil)
	localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
	localPwd := pwd
	localPwd.ID = "123"
	localPwd.User = credentials.Email
	localStorageMock.EXPECT().AddCredentials(gomock.Any(), &localPwd).Return("123", nil)
	updatePwdMock := mock_services.NewMockUpdateDeletePasswordService(ctrl)
	updateCrdMock := mock_services.NewMockUpdateDeleteCardService(ctrl)
	updateTxtMock := mock_services.NewMockUpdateDeleteTextService(ctrl)
	updateBinMock := mock_services.NewMockUpdateDeleteBinaryService(ctrl)

	itemsService := NewSecretsService(
		secretsClientMock,
		authServiceMock,
		itemCryptoMock,
		localStorageMock,
		userCredsMock,
		updatePwdMock,
		updateCrdMock,
		updateTxtMock,
		updateBinMock,
	)
	err := itemsService.AddCredentials(context.Background(), &pwd)

	require.NoError(t, err)
}

func TestSecretsService_AddTextInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	item := dto.BaseSecret{Name: "textinfo", Metadata: "metadata", UpdatedAt: time.Now().UTC()}
	textInfo := dto.TextInfo{
		BaseSecret: item,
		Text:       "test text",
	}
	expectedRequest := secrets.TextInfo{
		Name:      textInfo.Name,
		Text:      textInfo.Text,
		Metadata:  textInfo.Metadata,
		UpdatedAt: timestamppb.New(item.UpdatedAt),
	}
	credentials := dto.UserCredentials{Email: "email", Token: "token"}

	userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
	userCredsMock.EXPECT().GetCredentials().Return(&credentials, nil)
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	authServiceMock.EXPECT().AddAuthMetadata(gomock.Any(), credentials.Token).Return(context.Background(), nil)
	itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
	itemCryptoMock.EXPECT().EncryptSecret(&textInfo).Return(nil)
	secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	response := secrets.AddResponse{Id: "123"}
	secretsClientMock.EXPECT().AddTextInfo(gomock.Any(), &expectedRequest).Return(&response, nil)
	localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
	localTextInfo := textInfo
	localTextInfo.ID = "123"
	localTextInfo.User = credentials.Email
	localStorageMock.EXPECT().AddTextInfo(gomock.Any(), &localTextInfo).Return("123", nil)
	updatePwdMock := mock_services.NewMockUpdateDeletePasswordService(ctrl)
	updateCrdMock := mock_services.NewMockUpdateDeleteCardService(ctrl)
	updateTxtMock := mock_services.NewMockUpdateDeleteTextService(ctrl)
	updateBinMock := mock_services.NewMockUpdateDeleteBinaryService(ctrl)

	itemsService := NewSecretsService(
		secretsClientMock,
		authServiceMock,
		itemCryptoMock,
		localStorageMock,
		userCredsMock,
		updatePwdMock,
		updateCrdMock,
		updateTxtMock,
		updateBinMock,
	)
	err := itemsService.AddTextInfo(context.Background(), &textInfo)

	require.NoError(t, err)
}

func TestSecretsService_AddCardInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	item := dto.BaseSecret{Name: "cardinfo", Metadata: "metadata", UpdatedAt: time.Now().UTC()}
	cardInfo := dto.CardInfo{
		BaseSecret: item,
		Number:     "123123",
	}
	expectedRequest := secrets.CardInfo{
		Name:      cardInfo.Name,
		Number:    cardInfo.Number,
		Metadata:  cardInfo.Metadata,
		UpdatedAt: timestamppb.New(item.UpdatedAt),
	}
	credentials := dto.UserCredentials{Email: "email", Token: "token"}

	userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
	userCredsMock.EXPECT().GetCredentials().Return(&credentials, nil)
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	authServiceMock.EXPECT().AddAuthMetadata(gomock.Any(), credentials.Token).Return(context.Background(), nil)
	itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
	itemCryptoMock.EXPECT().EncryptSecret(&cardInfo).Return(nil)
	secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	response := secrets.AddResponse{Id: "123"}
	secretsClientMock.EXPECT().AddCardInfo(gomock.Any(), &expectedRequest).Return(&response, nil)
	localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
	localCardInfo := cardInfo
	localCardInfo.ID = "123"
	localCardInfo.User = credentials.Email
	localStorageMock.EXPECT().AddCardInfo(gomock.Any(), &localCardInfo).Return("123", nil)
	updatePwdMock := mock_services.NewMockUpdateDeletePasswordService(ctrl)
	updateCrdMock := mock_services.NewMockUpdateDeleteCardService(ctrl)
	updateTxtMock := mock_services.NewMockUpdateDeleteTextService(ctrl)
	updateBinMock := mock_services.NewMockUpdateDeleteBinaryService(ctrl)

	itemsService := NewSecretsService(
		secretsClientMock,
		authServiceMock,
		itemCryptoMock,
		localStorageMock,
		userCredsMock,
		updatePwdMock,
		updateCrdMock,
		updateTxtMock,
		updateBinMock,
	)
	err := itemsService.AddCardInfo(context.Background(), &cardInfo)

	require.NoError(t, err)
}

func TestSecretsService_AddBinaryInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	item := dto.BaseSecret{Name: "cardinfo", Metadata: "metadata", UpdatedAt: time.Now().UTC()}
	binInfo := dto.BinaryInfo{
		BaseSecret: item,
		Data:       "123123",
	}
	expectedRequest := secrets.BinaryInfo{
		Name:      binInfo.Name,
		Data:      binInfo.Data,
		Metadata:  binInfo.Metadata,
		UpdatedAt: timestamppb.New(item.UpdatedAt),
	}
	credentials := dto.UserCredentials{Email: "email", Token: "token"}

	userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
	userCredsMock.EXPECT().GetCredentials().Return(&credentials, nil)
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	authServiceMock.EXPECT().AddAuthMetadata(gomock.Any(), credentials.Token).Return(context.Background(), nil)
	itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
	itemCryptoMock.EXPECT().EncryptSecret(&binInfo).Return(nil)
	secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	response := secrets.AddResponse{Id: "123"}
	secretsClientMock.EXPECT().AddBinaryInfo(gomock.Any(), &expectedRequest).Return(&response, nil)
	localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
	localBinInfo := binInfo
	localBinInfo.ID = "123"
	localBinInfo.User = credentials.Email
	localStorageMock.EXPECT().AddBinaryInfo(gomock.Any(), &localBinInfo).Return("123", nil)
	updatePwdMock := mock_services.NewMockUpdateDeletePasswordService(ctrl)
	updateCrdMock := mock_services.NewMockUpdateDeleteCardService(ctrl)
	updateTxtMock := mock_services.NewMockUpdateDeleteTextService(ctrl)
	updateBinMock := mock_services.NewMockUpdateDeleteBinaryService(ctrl)

	itemsService := NewSecretsService(
		secretsClientMock,
		authServiceMock,
		itemCryptoMock,
		localStorageMock,
		userCredsMock,
		updatePwdMock,
		updateCrdMock,
		updateTxtMock,
		updateBinMock,
	)
	err := itemsService.AddBinaryInfo(context.Background(), &binInfo)

	require.NoError(t, err)
}

func TestSecretsService_ListSecrets(t *testing.T) {
	now := time.Now().UTC()
	nowHourLater := now.Add(time.Hour * 1)
	tests := []struct {
		name          string
		remoteSecrets *secrets.ListSecretResponse
		localSecrets  dto.SecretsList
		expResult     dto.SecretsList
		setupMocks    func(
			remoteStorage *mock_secrets.MockSecretsManagementClient,
			localStorage *mock_services.MockLocalSecretsStorage,
			updatePwdService *mock_services.MockUpdateDeletePasswordService,
			updateCrdService *mock_services.MockUpdateDeleteCardService,
			updateTxtService *mock_services.MockUpdateDeleteTextService,
			updateBinService *mock_services.MockUpdateDeleteBinaryService,
		)
	}{
		{
			name: "all ids are in sync",
			remoteSecrets: &secrets.ListSecretResponse{
				Passwords: []*secrets.Password{
					{Id: "pwd1", Login: "login1", UpdatedAt: timestamppb.New(now)},
					{Id: "pwd2", Login: "login2", UpdatedAt: timestamppb.New(now)},
				},
				Texts: []*secrets.TextInfo{
					{Id: "txt1", Text: "text1", UpdatedAt: timestamppb.New(now)},
					{Id: "txt2", Text: "text2", UpdatedAt: timestamppb.New(now)},
				},
				Cards: []*secrets.CardInfo{
					{Id: "crd1", Number: "123", UpdatedAt: timestamppb.New(now)},
					{Id: "crd2", Number: "456", UpdatedAt: timestamppb.New(now)},
				},
				Bins: []*secrets.BinaryInfo{
					{Id: "bin1", Data: "data1", UpdatedAt: timestamppb.New(now)},
					{Id: "bin2", Data: "data2", UpdatedAt: timestamppb.New(now)},
				},
			},
			localSecrets: dto.SecretsList{
				Passwords: []*dto.LoginPassword{
					{BaseSecret: dto.BaseSecret{ID: "pwd1", UpdatedAt: now}, Login: "login1"},
					{BaseSecret: dto.BaseSecret{ID: "pwd2", UpdatedAt: now}, Login: "login2"},
				},
				Texts: []*dto.TextInfo{
					{BaseSecret: dto.BaseSecret{ID: "txt1", UpdatedAt: now}, Text: "text1"},
					{BaseSecret: dto.BaseSecret{ID: "txt2", UpdatedAt: now}, Text: "text2"},
				},
				Cards: []*dto.CardInfo{
					{BaseSecret: dto.BaseSecret{ID: "crd1", UpdatedAt: now}, Number: "123"},
					{BaseSecret: dto.BaseSecret{ID: "crd2", UpdatedAt: now}, Number: "456"},
				},
				Bins: []*dto.BinaryInfo{
					{BaseSecret: dto.BaseSecret{ID: "bin1", UpdatedAt: now}, Data: "data1"},
					{BaseSecret: dto.BaseSecret{ID: "bin2", UpdatedAt: now}, Data: "data2"},
				},
			},
			expResult: dto.SecretsList{
				Passwords: []*dto.LoginPassword{
					{BaseSecret: dto.BaseSecret{ID: "pwd1", UpdatedAt: now}, Login: "login1"},
					{BaseSecret: dto.BaseSecret{ID: "pwd2", UpdatedAt: now}, Login: "login2"},
				},
				Texts: []*dto.TextInfo{
					{BaseSecret: dto.BaseSecret{ID: "txt1", UpdatedAt: now}, Text: "text1"},
					{BaseSecret: dto.BaseSecret{ID: "txt2", UpdatedAt: now}, Text: "text2"},
				},
				Cards: []*dto.CardInfo{
					{BaseSecret: dto.BaseSecret{ID: "crd1", UpdatedAt: now}, Number: "123"},
					{BaseSecret: dto.BaseSecret{ID: "crd2", UpdatedAt: now}, Number: "456"},
				},
				Bins: []*dto.BinaryInfo{
					{BaseSecret: dto.BaseSecret{ID: "bin1", UpdatedAt: now}, Data: "data1"},
					{BaseSecret: dto.BaseSecret{ID: "bin2", UpdatedAt: now}, Data: "data2"},
				},
			},
			setupMocks: func(
				remoteStorage *mock_secrets.MockSecretsManagementClient,
				localStorage *mock_services.MockLocalSecretsStorage,
				updatePwdService *mock_services.MockUpdateDeletePasswordService,
				updateCrdService *mock_services.MockUpdateDeleteCardService,
				updateTxtService *mock_services.MockUpdateDeleteTextService,
				updateBinService *mock_services.MockUpdateDeleteBinaryService,
			) {
			},
		},
		{
			name: "all ids are in sync, remote secrets were updated later",
			remoteSecrets: &secrets.ListSecretResponse{
				Passwords: []*secrets.Password{
					{Id: "pwd1", Login: "remote login1", UpdatedAt: timestamppb.New(nowHourLater)},
				},
				Texts: []*secrets.TextInfo{
					{Id: "txt1", Text: "remote text1", UpdatedAt: timestamppb.New(nowHourLater)},
				},
				Cards: []*secrets.CardInfo{
					{Id: "crd1", Number: "remote 123", UpdatedAt: timestamppb.New(nowHourLater)},
				},
				Bins: []*secrets.BinaryInfo{
					{Id: "bin1", Data: "remote data1", UpdatedAt: timestamppb.New(nowHourLater)},
				},
			},
			localSecrets: dto.SecretsList{
				Passwords: []*dto.LoginPassword{
					{BaseSecret: dto.BaseSecret{ID: "pwd1", UpdatedAt: now}, Login: "login1"},
				},
				Texts: []*dto.TextInfo{
					{BaseSecret: dto.BaseSecret{ID: "txt1", UpdatedAt: now}, Text: "text1"},
				},
				Cards: []*dto.CardInfo{
					{BaseSecret: dto.BaseSecret{ID: "crd1", UpdatedAt: now}, Number: "123"},
				},
				Bins: []*dto.BinaryInfo{
					{BaseSecret: dto.BaseSecret{ID: "bin1", UpdatedAt: now}, Data: "data1"},
				},
			},
			expResult: dto.SecretsList{
				Passwords: []*dto.LoginPassword{
					{BaseSecret: dto.BaseSecret{ID: "pwd1", UpdatedAt: nowHourLater}, Login: "remote login1"},
				},
				Texts: []*dto.TextInfo{
					{BaseSecret: dto.BaseSecret{ID: "txt1", UpdatedAt: nowHourLater}, Text: "remote text1"},
				},
				Cards: []*dto.CardInfo{
					{BaseSecret: dto.BaseSecret{ID: "crd1", UpdatedAt: nowHourLater}, Number: "remote 123"},
				},
				Bins: []*dto.BinaryInfo{
					{BaseSecret: dto.BaseSecret{ID: "bin1", UpdatedAt: nowHourLater}, Data: "remote data1"},
				},
			},
			setupMocks: func(
				remoteStorage *mock_secrets.MockSecretsManagementClient,
				localStorage *mock_services.MockLocalSecretsStorage,
				updatePwdService *mock_services.MockUpdateDeletePasswordService,
				updateCrdService *mock_services.MockUpdateDeleteCardService,
				updateTxtService *mock_services.MockUpdateDeleteTextService,
				updateBinService *mock_services.MockUpdateDeleteBinaryService,
			) {
				updatePwdService.EXPECT().UpdateLocalSecret(gomock.Any(), gomock.Any()).Return(nil)
				updateCrdService.EXPECT().UpdateLocalSecret(gomock.Any(), gomock.Any()).Return(nil)
				updateTxtService.EXPECT().UpdateLocalSecret(gomock.Any(), gomock.Any()).Return(nil)
				updateBinService.EXPECT().UpdateLocalSecret(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "all ids are in sync, local secrets were updated later",
			remoteSecrets: &secrets.ListSecretResponse{
				Passwords: []*secrets.Password{
					{Id: "pwd1", Login: "remote login1", UpdatedAt: timestamppb.New(now)},
				},
				Texts: []*secrets.TextInfo{
					{Id: "txt1", Text: "remote text1", UpdatedAt: timestamppb.New(now)},
				},
				Cards: []*secrets.CardInfo{
					{Id: "crd1", Number: "remote 123", UpdatedAt: timestamppb.New(now)},
				},
				Bins: []*secrets.BinaryInfo{
					{Id: "bin1", Data: "remote data1", UpdatedAt: timestamppb.New(now)},
				},
			},
			localSecrets: dto.SecretsList{
				Passwords: []*dto.LoginPassword{
					{BaseSecret: dto.BaseSecret{ID: "pwd1", UpdatedAt: nowHourLater}, Login: "local login1"},
				},
				Texts: []*dto.TextInfo{
					{BaseSecret: dto.BaseSecret{ID: "txt1", UpdatedAt: nowHourLater}, Text: "local text1"},
				},
				Cards: []*dto.CardInfo{
					{BaseSecret: dto.BaseSecret{ID: "crd1", UpdatedAt: nowHourLater}, Number: "local 123"},
				},
				Bins: []*dto.BinaryInfo{
					{BaseSecret: dto.BaseSecret{ID: "bin1", UpdatedAt: nowHourLater}, Data: "local data1"},
				},
			},
			expResult: dto.SecretsList{
				Passwords: []*dto.LoginPassword{
					{BaseSecret: dto.BaseSecret{ID: "pwd1", UpdatedAt: nowHourLater}, Login: "local login1"},
				},
				Texts: []*dto.TextInfo{
					{BaseSecret: dto.BaseSecret{ID: "txt1", UpdatedAt: nowHourLater}, Text: "local text1"},
				},
				Cards: []*dto.CardInfo{
					{BaseSecret: dto.BaseSecret{ID: "crd1", UpdatedAt: nowHourLater}, Number: "local 123"},
				},
				Bins: []*dto.BinaryInfo{
					{BaseSecret: dto.BaseSecret{ID: "bin1", UpdatedAt: nowHourLater}, Data: "local data1"},
				},
			},
			setupMocks: func(
				remoteStorage *mock_secrets.MockSecretsManagementClient,
				localStorage *mock_services.MockLocalSecretsStorage,
				updatePwdService *mock_services.MockUpdateDeletePasswordService,
				updateCrdService *mock_services.MockUpdateDeleteCardService,
				updateTxtService *mock_services.MockUpdateDeleteTextService,
				updateBinService *mock_services.MockUpdateDeleteBinaryService,
			) {
				updatePwdService.EXPECT().UpdateRemoteSecret(gomock.Any(), gomock.Any()).Return(nil)
				updateCrdService.EXPECT().UpdateRemoteSecret(gomock.Any(), gomock.Any()).Return(nil)
				updateTxtService.EXPECT().UpdateRemoteSecret(gomock.Any(), gomock.Any()).Return(nil)
				updateBinService.EXPECT().UpdateRemoteSecret(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "all ids are not in sync",
			remoteSecrets: &secrets.ListSecretResponse{
				Passwords: []*secrets.Password{
					{Id: "pwd1", Login: "login1", UpdatedAt: timestamppb.New(now)},
				},
				Texts: []*secrets.TextInfo{
					{Id: "txt1", Text: "text1", UpdatedAt: timestamppb.New(now)},
				},
				Cards: []*secrets.CardInfo{
					{Id: "crd1", Number: "123", UpdatedAt: timestamppb.New(now)},
				},
				Bins: []*secrets.BinaryInfo{
					{Id: "bin1", Data: "data1", UpdatedAt: timestamppb.New(now)},
				},
			},
			localSecrets: dto.SecretsList{
				Passwords: []*dto.LoginPassword{
					{BaseSecret: dto.BaseSecret{ID: "pwd2", UpdatedAt: now}, Login: "login2"},
				},
				Texts: []*dto.TextInfo{
					{BaseSecret: dto.BaseSecret{ID: "txt2", UpdatedAt: now}, Text: "text2"},
				},
				Cards: []*dto.CardInfo{
					{BaseSecret: dto.BaseSecret{ID: "crd2", UpdatedAt: now}, Number: "456"},
				},
				Bins: []*dto.BinaryInfo{
					{BaseSecret: dto.BaseSecret{ID: "bin2", UpdatedAt: now}, Data: "data2"},
				},
			},
			expResult: dto.SecretsList{
				Passwords: []*dto.LoginPassword{
					{BaseSecret: dto.BaseSecret{ID: "pwd1", UpdatedAt: now}, Login: "login1"},
					{BaseSecret: dto.BaseSecret{ID: "pwd2", UpdatedAt: now}, Login: "login2"},
				},
				Texts: []*dto.TextInfo{
					{BaseSecret: dto.BaseSecret{ID: "txt1", UpdatedAt: now}, Text: "text1"},
					{BaseSecret: dto.BaseSecret{ID: "txt2", UpdatedAt: now}, Text: "text2"},
				},
				Cards: []*dto.CardInfo{
					{BaseSecret: dto.BaseSecret{ID: "crd1", UpdatedAt: now}, Number: "123"},
					{BaseSecret: dto.BaseSecret{ID: "crd2", UpdatedAt: now}, Number: "456"},
				},
				Bins: []*dto.BinaryInfo{
					{BaseSecret: dto.BaseSecret{ID: "bin1", UpdatedAt: now}, Data: "data1"},
					{BaseSecret: dto.BaseSecret{ID: "bin2", UpdatedAt: now}, Data: "data2"},
				},
			},
			setupMocks: func(
				remoteStorage *mock_secrets.MockSecretsManagementClient,
				localStorage *mock_services.MockLocalSecretsStorage,
				updatePwdService *mock_services.MockUpdateDeletePasswordService,
				updateCrdService *mock_services.MockUpdateDeleteCardService,
				updateTxtService *mock_services.MockUpdateDeleteTextService,
				updateBinService *mock_services.MockUpdateDeleteBinaryService,
			) {
				response := &secrets.AddResponse{Id: "123"}
				remoteStorage.EXPECT().AddBinaryInfo(gomock.Any(), gomock.Any()).Return(response, nil)
				remoteStorage.EXPECT().AddCredentials(gomock.Any(), gomock.Any()).Return(response, nil)
				remoteStorage.EXPECT().AddCardInfo(gomock.Any(), gomock.Any()).Return(response, nil)
				remoteStorage.EXPECT().AddTextInfo(gomock.Any(), gomock.Any()).Return(response, nil)

				localStorage.EXPECT().AddBinaryInfo(gomock.Any(), gomock.Any()).Return("123", nil)
				localStorage.EXPECT().AddCredentials(gomock.Any(), gomock.Any()).Return("123", nil)
				localStorage.EXPECT().AddCardInfo(gomock.Any(), gomock.Any()).Return("123", nil)
				localStorage.EXPECT().AddTextInfo(gomock.Any(), gomock.Any()).Return("123", nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			credentials := dto.UserCredentials{Email: "email", Token: "token"}
			userCredsMock := mock_services.NewMockUserCredentialsStorage(ctrl)
			userCredsMock.EXPECT().GetCredentials().Return(&credentials, nil).AnyTimes()
			authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
			authServiceMock.EXPECT().AddAuthMetadata(gomock.Any(), credentials.Token).Return(context.Background(), nil).AnyTimes()
			secretsClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
			secretsClientMock.EXPECT().ListSecrets(gomock.Any(), &secrets.EmptyRequest{}).Return(tt.remoteSecrets, nil)
			itemCryptoMock := mock_services.NewMockSecretCryptoService(ctrl)
			itemCryptoMock.EXPECT().DecryptSecret(gomock.Any()).Return(nil).AnyTimes()
			localStorageMock := mock_services.NewMockLocalSecretsStorage(ctrl)
			localStorageMock.EXPECT().ListSecrets(gomock.Any(), "email").Return(tt.localSecrets, nil)
			updatePwdMock := mock_services.NewMockUpdateDeletePasswordService(ctrl)
			updateCrdMock := mock_services.NewMockUpdateDeleteCardService(ctrl)
			updateTxtMock := mock_services.NewMockUpdateDeleteTextService(ctrl)
			updateBinMock := mock_services.NewMockUpdateDeleteBinaryService(ctrl)
			tt.setupMocks(secretsClientMock, localStorageMock, updatePwdMock, updateCrdMock, updateTxtMock, updateBinMock)

			itemsService := NewSecretsService(
				secretsClientMock,
				authServiceMock,
				itemCryptoMock,
				localStorageMock,
				userCredsMock,
				updatePwdMock,
				updateCrdMock,
				updateTxtMock,
				updateBinMock,
			)
			result, err := itemsService.ListSecrets(context.Background())

			require.NoError(t, err)
			for _, expSecret := range tt.expResult.Passwords {
				foundSecret := false
				for _, secret := range result.Passwords {
					if secret.ID == expSecret.ID {
						assert.Equal(t, expSecret, secret)
						foundSecret = true
						break
					}
				}
				assert.True(t, foundSecret)
			}
			for _, expSecret := range tt.expResult.Cards {
				foundSecret := false
				for _, secret := range result.Cards {
					if secret.ID == expSecret.ID {
						assert.Equal(t, expSecret, secret)
						foundSecret = true
						break
					}
				}
				assert.True(t, foundSecret)
			}
			for _, expSecret := range tt.expResult.Texts {
				foundSecret := false
				for _, secret := range result.Texts {
					if secret.ID == expSecret.ID {
						assert.Equal(t, expSecret, secret)
						foundSecret = true
						break
					}
				}
				assert.True(t, foundSecret)
			}
			for _, expSecret := range tt.expResult.Bins {
				foundSecret := false
				for _, secret := range result.Bins {
					if secret.ID == expSecret.ID {
						assert.Equal(t, expSecret, secret)
						foundSecret = true
						break
					}
				}
				assert.True(t, foundSecret)
			}
		})
	}
}
