package handlers

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	commonstorage "github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
	mock_handlers "github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers/mocks"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestSecretsServer_AddCredentials(t *testing.T) {
	userEmail := "test@email.com"
	item := dto.BaseSecret{
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
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
			request:     common.CredentialsToProto(&password),
			expResponse: &secrets.AddResponse{Id: "123"},
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
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
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
			request:     common.TextToProto(&textInfo),
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
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
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
			request:     common.CardToProto(&cardInfo),
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
	now := time.Now().UTC()
	secret := dto.BaseSecret{UpdatedAt: now}
	itemsList := dto.SecretsList{
		Passwords: []*dto.LoginPassword{{Password: "pwd1", BaseSecret: secret}, {Password: "pwd2", BaseSecret: secret}},
		Texts:     []*dto.TextInfo{{Text: "text1", BaseSecret: secret}, {Text: "text2", BaseSecret: secret}},
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
					{Password: itemsList.Passwords[0].Password, UpdatedAt: timestamppb.New(now)},
					{Password: itemsList.Passwords[1].Password, UpdatedAt: timestamppb.New(now)},
				},
				Texts: []*secrets.TextInfo{
					{Text: itemsList.Texts[0].Text, UpdatedAt: timestamppb.New(now)},
					{Text: itemsList.Texts[1].Text, UpdatedAt: timestamppb.New(now)},
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

func TestSecretsServer_GetCredentialsByID(t *testing.T) {
	userEmail := "test@email.com"
	item := dto.BaseSecret{
		ID:        "123",
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
	}
	password := dto.LoginPassword{
		Login:      "login",
		Password:   "pwd",
		BaseSecret: item,
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request     *secrets.GetSecretRequest
		expResponse *secrets.GetCredentialsResponse
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().GetCredentialsById(gomock.Any(), "123", userEmail).Return(&password, nil)
			},
			request:     &secrets.GetSecretRequest{Id: "123"},
			expResponse: &secrets.GetCredentialsResponse{Password: common.CredentialsToProto(&password)},
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
			response, err := itemsServer.GetCredentialsByID(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}

func TestSecretsServer_GetCardByID(t *testing.T) {
	userEmail := "test@email.com"
	baseSecret := dto.BaseSecret{
		ID:        "123",
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
	}
	secret := dto.CardInfo{
		Number:     "123",
		BaseSecret: baseSecret,
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request     *secrets.GetSecretRequest
		expResponse *secrets.GetCardResponse
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().GetCardById(gomock.Any(), "123", userEmail).Return(&secret, nil)
			},
			request:     &secrets.GetSecretRequest{Id: "123"},
			expResponse: &secrets.GetCardResponse{Card: common.CardToProto(&secret)},
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
			response, err := itemsServer.GetCardByID(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}

func TestSecretsServer_GetTextByID(t *testing.T) {
	userEmail := "test@email.com"
	baseSecret := dto.BaseSecret{
		ID:        "123",
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
	}
	secret := dto.TextInfo{
		Text:       "some text",
		BaseSecret: baseSecret,
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request     *secrets.GetSecretRequest
		expResponse *secrets.GetTextResponse
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().GetTextById(gomock.Any(), "123", userEmail).Return(&secret, nil)
			},
			request:     &secrets.GetSecretRequest{Id: "123"},
			expResponse: &secrets.GetTextResponse{Text: common.TextToProto(&secret)},
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
			response, err := itemsServer.GetTextByID(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}

func TestSecretsServer_GetBinaryByID(t *testing.T) {
	userEmail := "test@email.com"
	baseSecret := dto.BaseSecret{
		ID:        "123",
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
	}
	secret := dto.BinaryInfo{
		Data:       "some text",
		BaseSecret: baseSecret,
	}

	tests := []struct {
		name        string
		setupMocks  func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request     *secrets.GetSecretRequest
		expResponse *secrets.GetBinaryResponse
		expError    error
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().GetBinaryById(gomock.Any(), "123", userEmail).Return(&secret, nil)
			},
			request:     &secrets.GetSecretRequest{Id: "123"},
			expResponse: &secrets.GetBinaryResponse{Bin: common.BinaryToProto(&secret)},
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
			response, err := itemsServer.GetBinaryByID(context.Background(), tt.request)

			assert.Equal(t, tt.expResponse, response)
			assert.Equal(t, tt.expError, err)
		})
	}
}

func initServerAndMocks(
	ctrl *gomock.Controller,
	setupMocks func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService),
) *SecretsServer {
	authServiceMock := mock_handlers.NewMockAuthService(ctrl)
	itemsServiceMock := mock_handlers.NewMockSecretsService(ctrl)
	setupMocks(authServiceMock, itemsServiceMock)
	itemsServer := NewSecretsServer(itemsServiceMock, authServiceMock)
	return itemsServer
}

func TestSecretsServer_UpdateCredentials(t *testing.T) {
	userEmail := "test@email.com"
	item := dto.BaseSecret{
		ID:        "123",
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
	}
	password := dto.LoginPassword{
		Login:      "login",
		Password:   "pwd",
		BaseSecret: item,
	}

	tests := []struct {
		name       string
		setupMocks func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request    *secrets.Password
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().UpdateCredentials(gomock.Any(), &password).Return(nil)
			},
			request: common.CredentialsToProto(&password),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			secretsServer := initServerAndMocks(ctrl, tt.setupMocks)
			_, err := secretsServer.UpdateCredentials(context.Background(), tt.request)
			require.NoError(t, err)
		})
	}
}

func TestSecretsServer_UpdateCardInfo(t *testing.T) {
	userEmail := "test@email.com"
	baseSecret := dto.BaseSecret{
		ID:        "123",
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
	}
	secret := dto.CardInfo{
		Number:     "123",
		BaseSecret: baseSecret,
	}

	tests := []struct {
		name       string
		setupMocks func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request    *secrets.CardInfo
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().UpdateCardInfo(gomock.Any(), &secret).Return(nil)
			},
			request: common.CardToProto(&secret),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			secretsServer := initServerAndMocks(ctrl, tt.setupMocks)
			_, err := secretsServer.UpdateCardInfo(context.Background(), tt.request)
			require.NoError(t, err)
		})
	}
}

func TestSecretsServer_UpdateTextInfo(t *testing.T) {
	userEmail := "test@email.com"
	baseSecret := dto.BaseSecret{
		ID:        "123",
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
	}
	secret := dto.TextInfo{
		Text:       "some text",
		BaseSecret: baseSecret,
	}

	tests := []struct {
		name       string
		setupMocks func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request    *secrets.TextInfo
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().UpdateTextInfo(gomock.Any(), &secret).Return(nil)
			},
			request: common.TextToProto(&secret),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			secretsServer := initServerAndMocks(ctrl, tt.setupMocks)
			_, err := secretsServer.UpdateTextInfo(context.Background(), tt.request)
			require.NoError(t, err)
		})
	}
}

func TestSecretsServer_UpdateBinaryInfo(t *testing.T) {
	userEmail := "test@email.com"
	baseSecret := dto.BaseSecret{
		ID:        "123",
		Name:      "name",
		User:      userEmail,
		Metadata:  "metadata",
		UpdatedAt: time.Now().UTC(),
	}
	secret := dto.BinaryInfo{
		Data:       "some text",
		BaseSecret: baseSecret,
	}

	tests := []struct {
		name       string
		setupMocks func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService)
		request    *secrets.BinaryInfo
	}{
		{
			name: "positive test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, itemsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return(userEmail, true)
				itemsMock.EXPECT().UpdateBinaryInfo(gomock.Any(), &secret).Return(nil)
			},
			request: common.BinaryToProto(&secret),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			secretsServer := initServerAndMocks(ctrl, tt.setupMocks)
			_, err := secretsServer.UpdateBinaryInfo(context.Background(), tt.request)
			require.NoError(t, err)
		})
	}
}

type deleteSecretTestCase struct {
	name             string
	setupMocks       func(authMock *mock_handlers.MockAuthService, secretsMock *mock_handlers.MockSecretsService)
	expError         error
	callDeleteMethod func(server *SecretsServer) error
}

func TestSecretsServer_DeleteSecret(t *testing.T) {
	id := "123"
	tests := []deleteSecretTestCase{
		{
			name: "simple password test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, secretsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return("test@mail.com", true)
				secretsMock.EXPECT().DeleteCredentials(context.Background(), id).Return(nil)
			},
			callDeleteMethod: func(server *SecretsServer) error {
				request := &secrets.DeleteSecretRequest{Id: id}
				_, err := server.DeleteCredentials(context.Background(), request)
				return err
			},
		},
		{
			name: "password does not exist or was already deleted",
			setupMocks: func(authMock *mock_handlers.MockAuthService, secretsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return("test@mail.com", true)
				secretsMock.EXPECT().DeleteCredentials(context.Background(), id).Return(commonstorage.ErrSecretDoesNotExistOrWasDeleted)
			},
			expError: status.Error(codes.FailedPrecondition, "Can't delete secret"),
			callDeleteMethod: func(server *SecretsServer) error {
				request := &secrets.DeleteSecretRequest{Id: id}
				_, err := server.DeleteCredentials(context.Background(), request)
				return err
			},
		},
		{
			name: "simple text info test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, secretsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return("test@mail.com", true)
				secretsMock.EXPECT().DeleteTextInfo(context.Background(), id).Return(nil)
			},
			callDeleteMethod: func(server *SecretsServer) error {
				request := &secrets.DeleteSecretRequest{Id: id}
				_, err := server.DeleteTextInfo(context.Background(), request)
				return err
			},
		},
		{
			name: "text info does not exist or was already deleted",
			setupMocks: func(authMock *mock_handlers.MockAuthService, secretsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return("test@mail.com", true)
				secretsMock.EXPECT().DeleteTextInfo(context.Background(), id).Return(commonstorage.ErrSecretDoesNotExistOrWasDeleted)
			},
			expError: status.Error(codes.FailedPrecondition, "Can't delete secret"),
			callDeleteMethod: func(server *SecretsServer) error {
				request := &secrets.DeleteSecretRequest{Id: id}
				_, err := server.DeleteTextInfo(context.Background(), request)
				return err
			},
		},
		{
			name: "simple binary info test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, secretsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return("test@mail.com", true)
				secretsMock.EXPECT().DeleteBinaryInfo(context.Background(), id).Return(nil)
			},
			callDeleteMethod: func(server *SecretsServer) error {
				request := &secrets.DeleteSecretRequest{Id: id}
				_, err := server.DeleteBinaryInfo(context.Background(), request)
				return err
			},
		},
		{
			name: "binary info does not exist or was already deleted",
			setupMocks: func(authMock *mock_handlers.MockAuthService, secretsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return("test@mail.com", true)
				secretsMock.EXPECT().DeleteBinaryInfo(context.Background(), id).Return(commonstorage.ErrSecretDoesNotExistOrWasDeleted)
			},
			expError: status.Error(codes.FailedPrecondition, "Can't delete secret"),
			callDeleteMethod: func(server *SecretsServer) error {
				request := &secrets.DeleteSecretRequest{Id: id}
				_, err := server.DeleteBinaryInfo(context.Background(), request)
				return err
			},
		},
		{
			name: "simple card info test",
			setupMocks: func(authMock *mock_handlers.MockAuthService, secretsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return("test@mail.com", true)
				secretsMock.EXPECT().DeleteCardInfo(context.Background(), id).Return(nil)
			},
			callDeleteMethod: func(server *SecretsServer) error {
				request := &secrets.DeleteSecretRequest{Id: id}
				_, err := server.DeleteCardInfo(context.Background(), request)
				return err
			},
		},
		{
			name: "card info does not exist or was already deleted",
			setupMocks: func(authMock *mock_handlers.MockAuthService, secretsMock *mock_handlers.MockSecretsService) {
				authMock.EXPECT().GetUserFromContext(gomock.Any()).Return("test@mail.com", true)
				secretsMock.EXPECT().DeleteCardInfo(context.Background(), id).Return(commonstorage.ErrSecretDoesNotExistOrWasDeleted)
			},
			expError: status.Error(codes.FailedPrecondition, "Can't delete secret"),
			callDeleteMethod: func(server *SecretsServer) error {
				request := &secrets.DeleteSecretRequest{Id: id}
				_, err := server.DeleteCardInfo(context.Background(), request)
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			secretsServer := initServerAndMocks(ctrl, tt.setupMocks)
			err := tt.callDeleteMethod(secretsServer)
			if tt.expError != nil {
				assert.Equal(t, tt.expError, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
