package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
)

type SecretsService struct {
	storage storage.SecretsStorage
}

func NewSecretsService(storage storage.SecretsStorage) *SecretsService {
	return &SecretsService{storage: storage}
}

func (s *SecretsService) AddPassword(ctx context.Context, password *dto.LoginPassword) (string, error) {
	return s.storage.AddPassword(ctx, password)
}

func (s *SecretsService) AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error {
	return s.storage.AddTextInfo(ctx, textInfo)
}

func (s *SecretsService) AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error {
	return s.storage.AddCardInfo(ctx, cardInfo)
}

func (s *SecretsService) ListSecrets(ctx context.Context, user string) (dto.SecretsList, error) {
	return s.storage.ListSecrets(ctx, user)
}