package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
)

type SecretsStorage interface {
	AddCredentials(ctx context.Context, password *dto.LoginPassword) error
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error
	AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error
	ListSecrets(ctx context.Context, user string) (dto.SecretsList, error)
}

type SecretsService struct {
	storage SecretsStorage
}

func NewSecretsService(storage SecretsStorage) *SecretsService {
	return &SecretsService{storage: storage}
}

func (s *SecretsService) AddCredentials(ctx context.Context, password *dto.LoginPassword) error {
	return s.storage.AddCredentials(ctx, password)
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
