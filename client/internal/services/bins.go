package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
)

type BinaryLocalStorage interface {
	GetBinaryById(ctx context.Context, id string, user string) (dto.BinaryInfo, error)
	UpdateBinaryInfo(ctx context.Context, crd *dto.BinaryInfo) error
	DeleteBinaryInfo(ctx context.Context, id string) error
}

// BinaryInstanceService - service to perform operations with single BinaryInfo instance.
type BinaryInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  BinaryLocalStorage
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
func (s *BinaryInstanceService) GetLocalSecretByID(ctx context.Context, id string, email string) (dto.BinaryInfo, error) {
	return s.localStorage.GetBinaryById(ctx, id, email)
}

// UpdateSecret updates certain BinaryInfo in remote storage.
func (s *BinaryInstanceService) UpdateSecret(ctx context.Context, bin dto.BinaryInfo) error {
	request := common.BinaryToProto(&bin)
	_, err := s.storageClient.UpdateBinaryInfo(ctx, request)
	return err
}

// UpdateLocalSecret updates certain BinaryInfo in local storage.
func (s *BinaryInstanceService) UpdateLocalSecret(ctx context.Context, bin dto.BinaryInfo) error {
	return s.localStorage.UpdateBinaryInfo(ctx, &bin)
}

// DeleteSecret deletes certain BinaryInfo in remote storage.
func (s *BinaryInstanceService) DeleteSecret(ctx context.Context, id string) error {
	_, err := s.storageClient.DeleteBinaryInfo(ctx, &secrets.DeleteSecretRequest{Id: id})
	return err
}

// DeleteLocalSecret deletes certain BinaryInfo in local storage.
func (s *BinaryInstanceService) DeleteLocalSecret(ctx context.Context, id string) error {
	return s.localStorage.DeleteBinaryInfo(ctx, id)
}
