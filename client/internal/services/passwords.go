package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PasswordInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  LocalItemsStorage
}

func (s *PasswordInstanceService) GetSecretByID(ctx context.Context, id string) (dto.LoginPassword, error) {
	request := secrets.GetSecretRequest{Id: id}
	response, err := s.storageClient.GetCredentialsByID(ctx, &request)
	if err != nil {
		return dto.LoginPassword{}, err
	}

	pwd := response.Password
	itemDest := dto.BaseSecret{
		ID:       pwd.Id,
		Name:     pwd.Name,
		Metadata: pwd.Metadata,
		User:     pwd.User,
	}
	pwdDest := dto.LoginPassword{
		BaseSecret: itemDest,
		Login:      pwd.Login,
		Password:   pwd.Password,
	}

	return pwdDest, nil
}

func (s *PasswordInstanceService) GetLocalSecretByID(id string, email string) (dto.LoginPassword, error) {
	return s.localStorage.GetCredentialsById(context.Background(), id, email)
}

func (s *PasswordInstanceService) UpdateSecret(ctx context.Context, loginPwd dto.LoginPassword) error {
	request := secrets.Password{
		Name:      loginPwd.Name,
		Login:     loginPwd.Login,
		Password:  loginPwd.Password,
		Metadata:  loginPwd.Metadata,
		UpdatedAt: timestamppb.New(loginPwd.UpdatedAt),
	}
	_, err := s.storageClient.UpdateCredentials(ctx, &request)
	return err
}

func (s *PasswordInstanceService) UpdateLocalSecret(loginPwd dto.LoginPassword) error {
	return s.localStorage.UpdateCredentials(context.Background(), &loginPwd)
}
