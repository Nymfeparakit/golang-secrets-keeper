package services

import (
	"context"
	mock_services "github.com/Nymfeparakit/gophkeeper/client/internal/services/mocks"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/items"
	mock_items "github.com/Nymfeparakit/gophkeeper/server/proto/items/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestItemsService_AddPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	authServiceMock.EXPECT().AddAuthMetadata(gomock.Any()).Return(context.Background(), nil)
	item := dto.Item{Name: "pwd"}
	pwd := dto.LoginPassword{
		Item:     item,
		Login:    "login",
		Password: "pwd",
	}
	request := items.Password{
		Name:     pwd.Name,
		Login:    pwd.Login,
		Password: pwd.Password,
	}
	authClientMock := mock_items.NewMockItemsManagementClient(ctrl)
	response := items.Response{}
	authClientMock.EXPECT().AddPassword(gomock.Any(), &request).Return(&response, nil)
	itemCryptoMock := mock_services.NewMockItemCryptoService(ctrl)
	itemCryptoMock.EXPECT().EncryptItem(&pwd).Return(nil)

	itemsService := NewItemsService(authClientMock, authServiceMock, itemCryptoMock)
	err := itemsService.AddPassword(&pwd)

	require.NoError(t, err)
}

func TestItemsService_AddTextInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	authServiceMock.EXPECT().AddAuthMetadata(gomock.Any()).Return(context.Background(), nil)
	item := dto.Item{Name: "textinfo", Metadata: "metadata"}
	textInfo := dto.TextInfo{
		Item: item,
		Text: "test text",
	}
	expectedRequest := items.TextInfo{
		Name:     textInfo.Name,
		Text:     textInfo.Text,
		Metadata: textInfo.Metadata,
	}
	authClientMock := mock_items.NewMockItemsManagementClient(ctrl)
	response := items.Response{}
	authClientMock.EXPECT().AddTextInfo(gomock.Any(), &expectedRequest).Return(&response, nil)
	itemCryptoMock := mock_services.NewMockItemCryptoService(ctrl)
	itemCryptoMock.EXPECT().EncryptItem(&textInfo).Return(nil)

	itemsService := NewItemsService(authClientMock, authServiceMock, itemCryptoMock)
	err := itemsService.AddTextInfo(&textInfo)

	require.NoError(t, err)
}

func TestItemsService_AddCardInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authServiceMock := mock_services.NewMockAuthMetadataService(ctrl)
	authServiceMock.EXPECT().AddAuthMetadata(gomock.Any()).Return(context.Background(), nil)
	item := dto.Item{Name: "cardinfo", Metadata: "metadata"}
	cardInfo := dto.CardInfo{
		Item:   item,
		Number: "123123",
	}
	expectedRequest := items.CardInfo{
		Name:     cardInfo.Name,
		Number:   cardInfo.Number,
		Metadata: cardInfo.Metadata,
	}
	authClientMock := mock_items.NewMockItemsManagementClient(ctrl)
	response := items.Response{}
	authClientMock.EXPECT().AddCardInfo(gomock.Any(), &expectedRequest).Return(&response, nil)
	itemCryptoMock := mock_services.NewMockItemCryptoService(ctrl)
	itemCryptoMock.EXPECT().EncryptItem(&cardInfo).Return(nil)

	itemsService := NewItemsService(authClientMock, authServiceMock, itemCryptoMock)
	err := itemsService.AddCardInfo(&cardInfo)

	require.NoError(t, err)
}
