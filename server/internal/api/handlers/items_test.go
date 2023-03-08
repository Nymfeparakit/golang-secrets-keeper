package handlers

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	mock_handlers "github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers/mocks"
	"github.com/Nymfeparakit/gophkeeper/server/proto/items"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemsServer_AddPassword(t *testing.T) {
	userEmail := "test@email.com"
	item := dto.Item{
		Name:     "name",
		User:     userEmail,
		Metadata: "metadata",
	}
	password := dto.LoginPassword{
		Login:    "login",
		Password: "pwd",
		Item:     item,
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockItemsService)
		request     *items.Password
		expResponse *items.Response
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockItemsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().AddPassword(gomock.Any(), &password)
			},
			request: &items.Password{
				Name:     password.Name,
				Login:    password.Login,
				Password: password.Password,
				Metadata: password.Metadata,
			},
			expResponse: &items.Response{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authServiceMock := mock_handlers.NewMockAuthService(ctrl)
			itemsServiceMock := mock_handlers.NewMockItemsService(ctrl)
			tt.setupMocks(authServiceMock, itemsServiceMock)
			itemsServer := NewItemsServer(itemsServiceMock, authServiceMock)
			response, err := itemsServer.AddPassword(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}

func TestItemsServer_AddTextInfo(t *testing.T) {
	userEmail := "test@email.com"
	item := dto.Item{
		Name:     "name",
		User:     userEmail,
		Metadata: "metadata",
	}
	textInfo := dto.TextInfo{
		Text: "text",
		Item: item,
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockItemsService)
		request     *items.TextInfo
		expResponse *items.Response
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockItemsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().AddTextInfo(gomock.Any(), &textInfo)
			},
			request: &items.TextInfo{
				Name:     textInfo.Name,
				Text:     textInfo.Text,
				Metadata: textInfo.Metadata,
			},
			expResponse: &items.Response{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authServiceMock := mock_handlers.NewMockAuthService(ctrl)
			itemsServiceMock := mock_handlers.NewMockItemsService(ctrl)
			tt.setupMocks(authServiceMock, itemsServiceMock)
			itemsServer := NewItemsServer(itemsServiceMock, authServiceMock)
			response, err := itemsServer.AddTextInfo(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}

func TestItemsServer_AddCardInfo(t *testing.T) {
	userEmail := "test@email.com"
	item := dto.Item{
		Name:     "name",
		User:     userEmail,
		Metadata: "metadata",
	}
	cardInfo := dto.CardInfo{
		Item:   item,
		Number: "123123",
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockItemsService)
		request     *items.CardInfo
		expResponse *items.Response
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockItemsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().AddCardInfo(gomock.Any(), &cardInfo)
			},
			request: &items.CardInfo{
				Name:     cardInfo.Name,
				Number:   cardInfo.Number,
				Metadata: cardInfo.Metadata,
			},
			expResponse: &items.Response{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authServiceMock := mock_handlers.NewMockAuthService(ctrl)
			itemsServiceMock := mock_handlers.NewMockItemsService(ctrl)
			tt.setupMocks(authServiceMock, itemsServiceMock)
			itemsServer := NewItemsServer(itemsServiceMock, authServiceMock)
			response, err := itemsServer.AddCardInfo(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}

func TestItemsServer_ListItems(t *testing.T) {
	userEmail := "test@email.com"
	itemsList := dto.ItemsList{
		Passwords: []*dto.LoginPassword{{Password: "pwd1"}, {Password: "pwd2"}},
		Texts:     []*dto.TextInfo{{Text: "text1"}, {Text: "text2"}},
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockItemsService)
		request     *items.EmptyRequest
		expResponse *items.ListItemResponse
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockItemsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().ListItems(gomock.Any(), userEmail).Return(itemsList, nil)
			},
			request: &items.EmptyRequest{},
			expResponse: &items.ListItemResponse{
				Passwords: []*items.Password{
					{Password: itemsList.Passwords[0].Password},
					{Password: itemsList.Passwords[1].Password},
				},
				Texts: []*items.TextInfo{
					{Text: itemsList.Texts[0].Text},
					{Text: itemsList.Texts[1].Text},
				},
				Error: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authServiceMock := mock_handlers.NewMockAuthService(ctrl)
			itemsServiceMock := mock_handlers.NewMockItemsService(ctrl)
			tt.setupMocks(authServiceMock, itemsServiceMock)
			itemsServer := NewItemsServer(itemsServiceMock, authServiceMock)
			response, err := itemsServer.ListItems(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}
