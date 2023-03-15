package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"time"
)

type CardInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  LocalSecretsStorage
}

func (s *CardInstanceService) GetSecretByID(ctx context.Context, id string) (dto.CardInfo, error) {
	request := secrets.GetSecretRequest{Id: id}
	response, err := s.storageClient.GetCardByID(ctx, &request)
	if err != nil {
		return dto.CardInfo{}, err
	}

	crdDest := common.CardFromProto(response.Card)
	return crdDest, nil
}

func (s *CardInstanceService) GetLocalSecretByID(id string, email string) (dto.CardInfo, error) {
	return s.localStorage.GetCardById(context.Background(), id, email)
}

func (s *CardInstanceService) UpdateSecret(ctx context.Context, crd dto.CardInfo) error {
	crd.UpdatedAt = time.Now().UTC()
	request := common.CardToProto(&crd)
	_, err := s.storageClient.UpdateCardInfo(ctx, request)
	return err
}

func (s *CardInstanceService) UpdateLocalSecret(crd dto.CardInfo) error {
	crd.UpdatedAt = time.Now().UTC()
	return s.localStorage.UpdateCardInfo(context.Background(), &crd)
}
