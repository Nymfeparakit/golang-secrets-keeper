package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"time"
)

type BinaryInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  LocalSecretsStorage
}

func (s *BinaryInstanceService) GetSecretByID(ctx context.Context, id string) (dto.BinaryInfo, error) {
	request := secrets.GetSecretRequest{Id: id}
	response, err := s.storageClient.GetBinaryByID(ctx, &request)
	if err != nil {
		return dto.BinaryInfo{}, err
	}

	binDest := common.BinaryFromProto(response.Bin)
	return binDest, nil
}

func (s *BinaryInstanceService) GetLocalSecretByID(id string, email string) (dto.BinaryInfo, error) {
	return s.localStorage.GetBinaryById(context.Background(), id, email)
}

func (s *BinaryInstanceService) UpdateSecret(ctx context.Context, bin dto.BinaryInfo) error {
	bin.UpdatedAt = time.Now().UTC()
	request := common.BinaryToProto(&bin)
	_, err := s.storageClient.UpdateBinaryInfo(ctx, request)
	return err
}

func (s *BinaryInstanceService) UpdateLocalSecret(bin dto.BinaryInfo) error {
	bin.UpdatedAt = time.Now().UTC()
	return s.localStorage.UpdateBinaryInfo(context.Background(), &bin)
}
