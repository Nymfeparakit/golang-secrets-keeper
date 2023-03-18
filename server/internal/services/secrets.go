package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
)

// SecretsService - service for performing CRUD operations with secrets.
type SecretsService struct {
	storage storage.SecretsStorage
}

// NewSecretsService creates new SecretsService object.
func NewSecretsService(storage storage.SecretsStorage) *SecretsService {
	return &SecretsService{storage: storage}
}

// AddCredentials creates new credentials secret in secrets storage.
func (s *SecretsService) AddCredentials(ctx context.Context, password *dto.LoginPassword) (string, error) {
	return s.storage.AddCredentials(ctx, password)
}

// AddTextInfo creates new text secret in secrets storage.
func (s *SecretsService) AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) (string, error) {
	return s.storage.AddTextInfo(ctx, textInfo)
}

// AddCardInfo creates new card secret in secrets storage.
func (s *SecretsService) AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) (string, error) {
	return s.storage.AddCardInfo(ctx, cardInfo)
}

// AddBinaryInfo creates new binary secret in secrets storage.
func (s *SecretsService) AddBinaryInfo(ctx context.Context, binInfo *dto.BinaryInfo) (string, error) {
	return s.storage.AddBinaryInfo(ctx, binInfo)
}

// ListSecrets returns from storage all user's secrets.
func (s *SecretsService) ListSecrets(ctx context.Context, user string) (dto.SecretsList, error) {
	return s.storage.ListSecrets(ctx, user)
}

// GetCredentialsById returns from storage credentials secret with specified id.
func (s *SecretsService) GetCredentialsById(ctx context.Context, id string, user string) (*dto.LoginPassword, error) {
	pwd, err := s.storage.GetCredentialsById(ctx, id, user)
	return &pwd, err
}

// GetCardById returns from storage card secret with specified id.
func (s *SecretsService) GetCardById(ctx context.Context, id string, user string) (*dto.CardInfo, error) {
	secret, err := s.storage.GetCardById(ctx, id, user)
	return &secret, err
}

// GetTextById returns from storage text secret with specified id.
func (s *SecretsService) GetTextById(ctx context.Context, id string, user string) (*dto.TextInfo, error) {
	secret, err := s.storage.GetTextById(ctx, id, user)
	return &secret, err
}

// GetBinaryById returns from storage binary secret with specified id.
func (s *SecretsService) GetBinaryById(ctx context.Context, id string, user string) (*dto.BinaryInfo, error) {
	secret, err := s.storage.GetBinaryById(ctx, id, user)
	return &secret, err
}

// UpdateCredentials updates in storage credentials secret with specified id.
func (s *SecretsService) UpdateCredentials(ctx context.Context, password *dto.LoginPassword) error {
	return s.storage.UpdateCredentials(ctx, password)
}

// UpdateTextInfo updates in storage text secret with specified id.
func (s *SecretsService) UpdateTextInfo(ctx context.Context, secret *dto.TextInfo) error {
	return s.storage.UpdateTextInfo(ctx, secret)
}

// UpdateBinaryInfo updates in storage binary secret with specified id.
func (s *SecretsService) UpdateBinaryInfo(ctx context.Context, secret *dto.BinaryInfo) error {
	return s.storage.UpdateBinaryInfo(ctx, secret)
}

// UpdateCardInfo updates in storage card secret with specified id.
func (s *SecretsService) UpdateCardInfo(ctx context.Context, secret *dto.CardInfo) error {
	return s.storage.UpdateCardInfo(ctx, secret)
}

// DeleteCredentials deletes from storage credentials secret with specified id.
func (s *SecretsService) DeleteCredentials(ctx context.Context, id string) error {
	return s.storage.DeleteCredentials(ctx, id)
}

// DeleteTextInfo deletes from storage text secret with specified id.
func (s *SecretsService) DeleteTextInfo(ctx context.Context, id string) error {
	return s.storage.DeleteTextInfo(ctx, id)
}

// DeleteCardInfo deletes from storage card secret with specified id.
func (s *SecretsService) DeleteCardInfo(ctx context.Context, id string) error {
	return s.storage.DeleteCardInfo(ctx, id)
}

// DeleteBinaryInfo deletes from storage binary secret with specified id.
func (s *SecretsService) DeleteBinaryInfo(ctx context.Context, id string) error {
	return s.storage.DeleteBinaryInfo(ctx, id)
}
