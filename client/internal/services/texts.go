package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"time"
)

type TextInstanceService struct {
	storageClient secrets.SecretsManagementClient
	localStorage  LocalSecretsStorage
}

func (s *TextInstanceService) GetSecretByID(ctx context.Context, id string) (dto.TextInfo, error) {
	request := secrets.GetSecretRequest{Id: id}
	response, err := s.storageClient.GetTextByID(ctx, &request)
	if err != nil {
		return dto.TextInfo{}, err
	}

	crdDest := common.TextFromProto(response.Text)
	return crdDest, nil
}

func (s *TextInstanceService) GetLocalSecretByID(id string, email string) (dto.TextInfo, error) {
	return s.localStorage.GetTextById(context.Background(), id, email)
}

func (s *TextInstanceService) UpdateSecret(ctx context.Context, txt dto.TextInfo) error {
	txt.UpdatedAt = time.Now().UTC()
	request := common.TextToProto(&txt)
	_, err := s.storageClient.UpdateTextInfo(ctx, request)
	return err
}

func (s *TextInstanceService) UpdateLocalSecret(txt dto.TextInfo) error {
	txt.UpdatedAt = time.Now().UTC()
	return s.localStorage.UpdateTextInfo(context.Background(), &txt)
}
