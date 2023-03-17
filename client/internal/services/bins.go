package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"time"
)

// BinaryInstanceService - service to perform operations with single BinaryInfo instance.
type BinaryInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  LocalSecretsStorage
}

// GetSecretByID gets BinaryInfo specified by its id from remote storage.
func (s *BinaryInstanceService) GetSecretByID(ctx context.Context, id string) (dto.BinaryInfo, error) {
	request := secrets.GetSecretRequest{Id: id}
	response, err := s.storageClient.GetBinaryByID(ctx, &request)
	if err != nil {
		return dto.BinaryInfo{}, err
	}

	binDest := common.BinaryFromProto(response.Bin)
	return binDest, nil
}

// GetLocalSecretByID gets BinaryInfo specified by its id from local storage.
func (s *BinaryInstanceService) GetLocalSecretByID(id string, email string) (dto.BinaryInfo, error) {
	return s.localStorage.GetBinaryById(context.Background(), id, email)
}

// UpdateSecret updates certain BinaryInfo in remote storage.
func (s *BinaryInstanceService) UpdateSecret(ctx context.Context, bin dto.BinaryInfo) error {
	bin.UpdatedAt = time.Now().UTC()
	request := common.BinaryToProto(&bin)
	_, err := s.storageClient.UpdateBinaryInfo(ctx, request)
	return err
}

// UpdateLocalSecret updates certain BinaryInfo in local storage.
func (s *BinaryInstanceService) UpdateLocalSecret(bin dto.BinaryInfo) error {
	bin.UpdatedAt = time.Now().UTC()
	return s.localStorage.UpdateBinaryInfo(context.Background(), &bin)
}

// DeleteSecret deletes certain BinaryInfo in remote storage.
func (s *BinaryInstanceService) DeleteSecret(ctx context.Context, id string) error {
	_, err := s.storageClient.DeleteBinaryInfo(ctx, &secrets.DeleteSecretRequest{Id: id})
	return err
}

// DeleteLocalSecret deletes certain BinaryInfo in local storage.
func (s *BinaryInstanceService) DeleteLocalSecret(id string) error {
	return s.localStorage.DeleteBinaryInfo(context.Background(), id)
}
