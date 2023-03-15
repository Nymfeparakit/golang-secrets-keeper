package services

import (
	"context"
	mock_services "github.com/Nymfeparakit/gophkeeper/client/internal/services/mocks"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	mock_secrets "github.com/Nymfeparakit/gophkeeper/server/proto/secrets/mocks"
	"github.com/golang/mock/gomock"
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
	localStorageMock.EXPECT().AddCredentials(gomock.Any(), &localPwd).Return(nil)
	updatePwdMock := mock_services.NewMockUpdatePasswordsService(ctrl)

	itemsService := NewSecretsService(
		secretsClientMock,
		authServiceMock,
		itemCryptoMock,
		localStorageMock,
		userCredsMock,
		updatePwdMock,
	)
	err := itemsService.AddCredentials(&pwd)

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
	localStorageMock.EXPECT().AddTextInfo(gomock.Any(), &localTextInfo).Return(nil)
	updatePwdMock := mock_services.NewMockUpdatePasswordsService(ctrl)

	itemsService := NewSecretsService(
		secretsClientMock,
		authServiceMock,
		itemCryptoMock,
		localStorageMock,
		userCredsMock,
		updatePwdMock,
	)
	err := itemsService.AddTextInfo(&textInfo)

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
	localStorageMock.EXPECT().AddCardInfo(gomock.Any(), &localCardInfo).Return(nil)
	updatePwdMock := mock_services.NewMockUpdatePasswordsService(ctrl)

	itemsService := NewSecretsService(
		secretsClientMock,
		authServiceMock,
		itemCryptoMock,
		localStorageMock,
		userCredsMock,
		updatePwdMock,
	)
	err := itemsService.AddCardInfo(&cardInfo)

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
	localStorageMock.EXPECT().AddBinaryInfo(gomock.Any(), &localBinInfo).Return(nil)
	updatePwdMock := mock_services.NewMockUpdatePasswordsService(ctrl)

	itemsService := NewSecretsService(
		secretsClientMock,
		authServiceMock,
		itemCryptoMock,
		localStorageMock,
		userCredsMock,
		updatePwdMock,
	)
	err := itemsService.AddBinaryInfo(&binInfo)

	require.NoError(t, err)
}
