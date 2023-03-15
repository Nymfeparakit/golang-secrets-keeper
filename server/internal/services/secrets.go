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

func (s *SecretsService) AddCredentials(ctx context.Context, password *dto.LoginPassword) (string, error) {
	return s.storage.AddCredentials(ctx, password)
}

func (s *SecretsService) AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) (string, error) {
	return s.storage.AddTextInfo(ctx, textInfo)
}

func (s *SecretsService) AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) (string, error) {
	return s.storage.AddCardInfo(ctx, cardInfo)
}

func (s *SecretsService) AddBinaryInfo(ctx context.Context, binInfo *dto.BinaryInfo) (string, error) {
	return s.storage.AddBinaryInfo(ctx, binInfo)
}

func (s *SecretsService) ListSecrets(ctx context.Context, user string) (dto.SecretsList, error) {
	return s.storage.ListSecrets(ctx, user)
}

func (s *SecretsService) GetCredentialsById(ctx context.Context, id string, user string) (*dto.LoginPassword, error) {
	pwd, err := s.storage.GetCredentialsById(ctx, id, user)
	return &pwd, err
}

func (s *SecretsService) GetCardById(ctx context.Context, id string, user string) (*dto.CardInfo, error) {
	secret, err := s.storage.GetCardById(ctx, id, user)
	return &secret, err
}

func (s *SecretsService) GetTextById(ctx context.Context, id string, user string) (*dto.TextInfo, error) {
	secret, err := s.storage.GetTextById(ctx, id, user)
	return &secret, err
}

func (s *SecretsService) GetBinaryById(ctx context.Context, id string, user string) (*dto.BinaryInfo, error) {
	secret, err := s.storage.GetBinaryById(ctx, id, user)
	return &secret, err
}

func (s *SecretsService) UpdateCredentials(ctx context.Context, password *dto.LoginPassword) error {
	return s.storage.UpdateCredentials(ctx, password)
}

func (s *SecretsService) UpdateTextInfo(ctx context.Context, secret *dto.TextInfo) error {
	return s.storage.UpdateTextInfo(ctx, secret)
}

func (s *SecretsService) UpdateBinaryInfo(ctx context.Context, secret *dto.BinaryInfo) error {
	return s.storage.UpdateBinaryInfo(ctx, secret)
}

func (s *SecretsService) UpdateCardInfo(ctx context.Context, secret *dto.CardInfo) error {
	return s.storage.UpdateCardInfo(ctx, secret)
}
