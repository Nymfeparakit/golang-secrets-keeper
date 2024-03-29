package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
)

type CardsLocalStorage interface {
	GetCardById(ctx context.Context, id string, user string) (dto.CardInfo, error)
	UpdateCardInfo(ctx context.Context, crd *dto.CardInfo) error
	DeleteCardInfo(ctx context.Context, id string) error
}

// CardInstanceService - service to perform operations with single CardInfo instance.
type CardInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  CardsLocalStorage
}

// GetSecretByID gets secret specified by its id from remote storage.
func (s *CardInstanceService) GetSecretByID(ctx context.Context, id string) (dto.CardInfo, error) {
	request := secrets.GetSecretRequest{Id: id}
	response, err := s.storageClient.GetCardByID(ctx, &request)
	if err != nil {
		return dto.CardInfo{}, err
	}

	crdDest := common.CardFromProto(response.Card)
	return crdDest, nil
}

// GetLocalSecretByID gets CardInfo specified by its id from local storage.
func (s *CardInstanceService) GetLocalSecretByID(ctx context.Context, id string, email string) (dto.CardInfo, error) {
	return s.localStorage.GetCardById(ctx, id, email)
}

// UpdateSecret updates certain CardInfo in remote storage.
func (s *CardInstanceService) UpdateSecret(ctx context.Context, crd dto.CardInfo) error {
	request := common.CardToProto(&crd)
	_, err := s.storageClient.UpdateCardInfo(ctx, request)
	return err
}

// UpdateLocalSecret updates certain CardInfo in local storage.
func (s *CardInstanceService) UpdateLocalSecret(ctx context.Context, crd dto.CardInfo) error {
	return s.localStorage.UpdateCardInfo(ctx, &crd)
}

// DeleteSecret deletes certain CardInfo in remote storage.
func (s *CardInstanceService) DeleteSecret(ctx context.Context, id string) error {
	_, err := s.storageClient.DeleteCardInfo(ctx, &secrets.DeleteSecretRequest{Id: id})
	return err
}

// DeleteLocalSecret deletes certain CardInfo in local storage.
func (s *CardInstanceService) DeleteLocalSecret(ctx context.Context, id string) error {
	return s.localStorage.DeleteCardInfo(ctx, id)
}
