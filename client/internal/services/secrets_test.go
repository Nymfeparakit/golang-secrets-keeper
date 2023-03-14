package services

import (
	"context"
	mock_services "github.com/Nymfeparakit/gophkeeper/client/internal/services/mocks"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	mock_secrets "github.com/Nymfeparakit/gophkeeper/server/proto/secrets/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestItemsService_AddPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	authServiceMock.EXPECT().AddAuthMetadata(gomock.Any()).Return(context.Background(), nil)
	item := dto.BaseSecret{Name: "pwd"}
	pwd := dto.LoginPassword{
		BaseSecret: item,
		Login:      "login",
		Password:   "pwd",
	}
	request := secrets.Password{
		Name:     pwd.Name,
		Login:    pwd.Login,
		Password: pwd.Password,
	}
	authClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	response := secrets.Response{}
	authClientMock.EXPECT().AddCredentials(gomock.Any(), &request).Return(&response, nil)
	itemCryptoMock := mock_services.NewMockItemCryptoService(ctrl)
	itemCryptoMock.EXPECT().EncryptItem(&pwd).Return(nil)

	itemsService := NewSecretsService(authClientMock, authServiceMock, itemCryptoMock)
	err := itemsService.AddCredentials(&pwd)

	require.NoError(t, err)
}

func TestItemsService_AddTextInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	authServiceMock.EXPECT().AddAuthMetadata(gomock.Any()).Return(context.Background(), nil)
	item := dto.BaseSecret{Name: "textinfo", Metadata: "metadata"}
	textInfo := dto.TextInfo{
		BaseSecret: item,
		Text:       "test text",
	}
	expectedRequest := secrets.TextInfo{
		Name:     textInfo.Name,
		Text:     textInfo.Text,
		Metadata: textInfo.Metadata,
	}
	authClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	response := secrets.Response{}
	authClientMock.EXPECT().AddTextInfo(gomock.Any(), &expectedRequest).Return(&response, nil)
	itemCryptoMock := mock_services.NewMockItemCryptoService(ctrl)
	itemCryptoMock.EXPECT().EncryptItem(&textInfo).Return(nil)

	itemsService := NewSecretsService(authClientMock, authServiceMock, itemCryptoMock)
	err := itemsService.AddTextInfo(&textInfo)

	require.NoError(t, err)
}

func TestItemsService_AddCardInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	authServiceMock.EXPECT().AddAuthMetadata(gomock.Any()).Return(context.Background(), nil)
	item := dto.BaseSecret{Name: "cardinfo", Metadata: "metadata"}
	cardInfo := dto.CardInfo{
		BaseSecret: item,
		Number:     "123123",
	}
	expectedRequest := secrets.CardInfo{
		Name:     cardInfo.Name,
		Number:   cardInfo.Number,
		Metadata: cardInfo.Metadata,
	}
	authClientMock := mock_secrets.NewMockSecretsManagementClient(ctrl)
	response := secrets.Response{}
	authClientMock.EXPECT().AddCardInfo(gomock.Any(), &expectedRequest).Return(&response, nil)
	itemCryptoMock := mock_services.NewMockItemCryptoService(ctrl)
	itemCryptoMock.EXPECT().EncryptItem(&cardInfo).Return(nil)

	itemsService := NewSecretsService(authClientMock, authServiceMock, itemCryptoMock)
	err := itemsService.AddCardInfo(&cardInfo)

	require.NoError(t, err)
}
