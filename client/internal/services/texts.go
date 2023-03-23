package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
)

type TextLocalStorage interface {
	GetTextById(ctx context.Context, id string, user string) (dto.TextInfo, error)
	UpdateTextInfo(ctx context.Context, txt *dto.TextInfo) error
	DeleteTextInfo(ctx context.Context, id string) error
}

// TextInstanceService - service to perform operations with single TextInfo instance.
type TextInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  TextLocalStorage
}

// GetSecretByID gets TextInfo specified by its id from remote storage.
func (s *TextInstanceService) GetSecretByID(ctx context.Context, id string) (dto.TextInfo, error) {
	request := secrets.GetSecretRequest{Id: id}
	response, err := s.storageClient.GetTextByID(ctx, &request)
	if err != nil {
		return dto.TextInfo{}, err
	}

	crdDest := common.TextFromProto(response.Text)
	return crdDest, nil
}

// GetLocalSecretByID gets TextInfo specified by its id from local storage.
func (s *TextInstanceService) GetLocalSecretByID(ctx context.Context, id string, email string) (dto.TextInfo, error) {
	return s.localStorage.GetTextById(ctx, id, email)
}

// UpdateSecret updates certain TextInfo in remote storage.
func (s *TextInstanceService) UpdateSecret(ctx context.Context, txt dto.TextInfo) error {
	request := common.TextToProto(&txt)
	_, err := s.storageClient.UpdateTextInfo(ctx, request)
	return err
}

// UpdateLocalSecret updates certain TextInfo in local storage.
func (s *TextInstanceService) UpdateLocalSecret(ctx context.Context, txt dto.TextInfo) error {
	return s.localStorage.UpdateTextInfo(ctx, &txt)
}

// DeleteSecret deletes certain CardInfo in remote storage.
func (s *TextInstanceService) DeleteSecret(ctx context.Context, id string) error {
	_, err := s.storageClient.DeleteTextInfo(ctx, &secrets.DeleteSecretRequest{Id: id})
	return err
}

// DeleteLocalSecret deletes certain CardInfo in local storage.
func (s *TextInstanceService) DeleteLocalSecret(ctx context.Context, id string) error {
	return s.localStorage.DeleteTextInfo(ctx, id)
}
