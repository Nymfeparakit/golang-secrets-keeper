package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"time"
)

// PasswordInstanceService - service to perform operations with single LoginPassword instance.
type PasswordInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  LocalSecretsStorage
}

// GetSecretByID gets LoginPassword specified by its id from remote storage.
func (s *PasswordInstanceService) GetSecretByID(ctx context.Context, id string) (dto.LoginPassword, error) {
	request := secrets.GetSecretRequest{Id: id}
	response, err := s.storageClient.GetCredentialsByID(ctx, &request)
	if err != nil {
		return dto.LoginPassword{}, err
	}

	pwd := response.Password
	dest := common.PasswordFromProto(pwd)

	return dest, nil
}

// GetLocalSecretByID gets LoginPassword specified by its id from local storage.
func (s *PasswordInstanceService) GetLocalSecretByID(id string, email string) (dto.LoginPassword, error) {
	return s.localStorage.GetCredentialsById(context.Background(), id, email)
}

// UpdateSecret updates certain LoginPassword in remote storage.
func (s *PasswordInstanceService) UpdateSecret(ctx context.Context, loginPwd dto.LoginPassword) error {
	loginPwd.UpdatedAt = time.Now().UTC()
	request := common.CredentialsToProto(&loginPwd)
	_, err := s.storageClient.UpdateCredentials(ctx, request)
	return err
}

// UpdateLocalSecret updates certain LoginPassword in local storage.
func (s *PasswordInstanceService) UpdateLocalSecret(loginPwd dto.LoginPassword) error {
	loginPwd.UpdatedAt = time.Now().UTC()
	return s.localStorage.UpdateCredentials(context.Background(), &loginPwd)
}
