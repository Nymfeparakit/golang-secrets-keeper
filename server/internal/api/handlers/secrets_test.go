package handlers

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	mock_handlers "github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers/mocks"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemsServer_AddPassword(t *testing.T) {
	userEmail := "test@email.com"
	item := dto.BaseSecret{
		Name:     "name",
		User:     userEmail,
		Metadata: "metadata",
	}
	password := dto.LoginPassword{
		Login:      "login",
		Password:   "pwd",
		BaseSecret: item,
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request     *secrets.Password
		expResponse *secrets.AddResponse
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().AddCredentials(gomock.Any(), &password).Return("123", nil)
			},
			request: &secrets.Password{
				Name:     password.Name,
				Login:    password.Login,
				Password: password.Password,
				Metadata: password.Metadata,
			},
			expResponse: &secrets.AddResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authServiceMock := mock_handlers.NewMockAuthService(ctrl)
			itemsServiceMock := mock_handlers.NewMockSecretsService(ctrl)
			tt.setupMocks(authServiceMock, itemsServiceMock)
			itemsServer := NewSecretsServer(itemsServiceMock, authServiceMock)
			response, err := itemsServer.AddCredentials(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}

func TestItemsServer_AddTextInfo(t *testing.T) {
	userEmail := "test@email.com"
	item := dto.BaseSecret{
		Name:     "name",
		User:     userEmail,
		Metadata: "metadata",
	}
	textInfo := dto.TextInfo{
		Text:       "text",
		BaseSecret: item,
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request     *secrets.TextInfo
		expResponse *secrets.AddResponse
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().AddTextInfo(gomock.Any(), &textInfo)
			},
			request: &secrets.TextInfo{
				Name:     textInfo.Name,
				Text:     textInfo.Text,
				Metadata: textInfo.Metadata,
			},
			expResponse: &secrets.AddResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authServiceMock := mock_handlers.NewMockAuthService(ctrl)
			itemsServiceMock := mock_handlers.NewMockSecretsService(ctrl)
			tt.setupMocks(authServiceMock, itemsServiceMock)
			itemsServer := NewSecretsServer(itemsServiceMock, authServiceMock)
			response, err := itemsServer.AddTextInfo(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}

func TestItemsServer_AddCardInfo(t *testing.T) {
	userEmail := "test@email.com"
	item := dto.BaseSecret{
		Name:     "name",
		User:     userEmail,
		Metadata: "metadata",
	}
	cardInfo := dto.CardInfo{
		BaseSecret: item,
		Number:     "123123",
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request     *secrets.CardInfo
		expResponse *secrets.AddResponse
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().AddCardInfo(gomock.Any(), &cardInfo)
			},
			request: &secrets.CardInfo{
				Name:     cardInfo.Name,
				Number:   cardInfo.Number,
				Metadata: cardInfo.Metadata,
			},
			expResponse: &secrets.AddResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authServiceMock := mock_handlers.NewMockAuthService(ctrl)
			itemsServiceMock := mock_handlers.NewMockSecretsService(ctrl)
			tt.setupMocks(authServiceMock, itemsServiceMock)
			itemsServer := NewSecretsServer(itemsServiceMock, authServiceMock)
			response, err := itemsServer.AddCardInfo(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}

func TestItemsServer_ListItems(t *testing.T) {
	userEmail := "test@email.com"
	itemsList := dto.SecretsList{
		Passwords: []*dto.LoginPassword{{Password: "pwd1"}, {Password: "pwd2"}},
		Texts:     []*dto.TextInfo{{Text: "text1"}, {Text: "text2"}},
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request     *secrets.EmptyRequest
		expResponse *secrets.ListSecretResponse
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().ListSecrets(gomock.Any(), userEmail).Return(itemsList, nil)
			},
			request: &secrets.EmptyRequest{},
			expResponse: &secrets.ListSecretResponse{
				Passwords: []*secrets.Password{
					{Password: itemsList.Passwords[0].Password},
					{Password: itemsList.Passwords[1].Password},
				},
				Texts: []*secrets.TextInfo{
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
			itemsServiceMock := mock_handlers.NewMockSecretsService(ctrl)
			tt.setupMocks(authServiceMock, itemsServiceMock)
			itemsServer := NewSecretsServer(itemsServiceMock, authServiceMock)
			response, err := itemsServer.ListSecrets(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}
