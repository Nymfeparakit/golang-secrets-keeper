package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"time"
)

// CardInstanceService - service to perform operations with single CardInfo instance.
type CardInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  LocalSecretsStorage
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
func (s *CardInstanceService) GetLocalSecretByID(id string, email string) (dto.CardInfo, error) {
	return s.localStorage.GetCardById(context.Background(), id, email)
}

// UpdateSecret updates certain CardInfo in remote storage.
func (s *CardInstanceService) UpdateSecret(ctx context.Context, crd dto.CardInfo) error {
	crd.UpdatedAt = time.Now().UTC()
	request := common.CardToProto(&crd)
	_, err := s.storageClient.UpdateCardInfo(ctx, request)
	return err
}

// UpdateLocalSecret updates certain CardInfo in local storage.
func (s *CardInstanceService) UpdateLocalSecret(crd dto.CardInfo) error {
	crd.UpdatedAt = time.Now().UTC()
	return s.localStorage.UpdateCardInfo(context.Background(), &crd)
}

// DeleteSecret deletes certain CardInfo in remote storage.
func (s *CardInstanceService) DeleteSecret(ctx context.Context, id string) error {
	_, err := s.storageClient.DeleteCardInfo(ctx, &secrets.DeleteSecretRequest{Id: id})
	return err
}

// DeleteLocalSecret deletes certain CardInfo in local storage.
func (s *CardInstanceService) DeleteLocalSecret(id string) error {
	return s.localStorage.DeleteCardInfo(context.Background(), id)
}
