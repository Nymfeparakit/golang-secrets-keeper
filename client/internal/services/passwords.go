package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"time"
)

type PasswordInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  LocalSecretsStorage
}

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

func (s *PasswordInstanceService) GetLocalSecretByID(id string, email string) (dto.LoginPassword, error) {
	return s.localStorage.GetCredentialsById(context.Background(), id, email)
}

func (s *PasswordInstanceService) UpdateSecret(ctx context.Context, loginPwd dto.LoginPassword) error {
	loginPwd.UpdatedAt = time.Now().UTC()
	request := common.CredentialsToProto(&loginPwd)
	_, err := s.storageClient.UpdateCredentials(ctx, request)
	return err
}

func (s *PasswordInstanceService) UpdateLocalSecret(loginPwd dto.LoginPassword) error {
	loginPwd.UpdatedAt = time.Now().UTC()
	return s.localStorage.UpdateCredentials(context.Background(), &loginPwd)
}
