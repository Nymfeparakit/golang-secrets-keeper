package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
)

type AuthMetadataService interface {
	AddAuthMetadata(ctx context.Context) (context.Context, error)
}

type SecretCryptoService interface {
	EncryptSecret(source any) error
	DecryptSecret(source any) error
}

type SecretsService struct {
	authService   AuthMetadataService
	storageClient secrets.SecretsManagementClient
	cryptoService SecretCryptoService
}

func NewSecretsService(
	client secrets.SecretsManagementClient,
	service AuthMetadataService,
	cryptoService SecretCryptoService,
) *SecretsService {
	return &SecretsService{storageClient: client, authService: service, cryptoService: cryptoService}
}

func (s *SecretsService) AddCredentials(loginPwd *dto.LoginPassword) error {
	// todo: context should be passed from argument
	ctx, err := s.authService.AddAuthMetadata(context.Background())
	if err != nil {
		return err
	}
	err = s.cryptoService.EncryptSecret(loginPwd)
	if err != nil {
		return fmt.Errorf("can't encrypt item: %s", err)
	}
	request := secrets.Password{
		Name:     loginPwd.Name,
		Login:    loginPwd.Login,
		Password: loginPwd.Password,
		Metadata: loginPwd.Metadata,
	}
	response, err := s.storageClient.AddCredentials(ctx, &request)
	if err != nil {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding password: %s", response.Error)
	}
	return nil
}

func (s *SecretsService) AddTextInfo(text *dto.TextInfo) error {
	ctx, err := s.authService.AddAuthMetadata(context.Background())
	if err != nil {
		return err
	}
	err = s.cryptoService.EncryptSecret(text)
	if err != nil {
		return fmt.Errorf("can't encrypt item: %s", err)
	}
	request := secrets.TextInfo{
		Name:     text.Name,
		Text:     text.Text,
		Metadata: text.Metadata,
	}
	response, err := s.storageClient.AddTextInfo(ctx, &request)
	if err != nil {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding password: %s", response.Error)
	}
	return nil
}

func (s *SecretsService) AddCardInfo(card *dto.CardInfo) error {
	ctx, err := s.authService.AddAuthMetadata(context.Background())
	if err != nil {
		return err
	}
	err = s.cryptoService.EncryptSecret(card)
	if err != nil {
		return fmt.Errorf("can't encrypt item: %s", err)
	}
	request := secrets.CardInfo{
		Name:            card.Name,
		Number:          card.Number,
		ExpirationMonth: card.ExpirationMonth,
		ExpirationYear:  card.ExpirationYear,
		Cvv:             card.CVV,
		Metadata:        card.Metadata,
	}
	response, err := s.storageClient.AddCardInfo(ctx, &request)
	if err != nil {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding password: %s", response.Error)
	}
	return nil
}

func (s *SecretsService) ListSecrets() (dto.SecretsList, error) {
	ctx, err := s.authService.AddAuthMetadata(context.Background())
	if err != nil {
		return dto.SecretsList{}, err
	}

	request := secrets.EmptyRequest{}
	response, err := s.storageClient.ListSecrets(ctx, &request)
	if err != nil {
		return dto.SecretsList{}, err
	}

	var secretsList dto.SecretsList
	for _, pwd := range response.Passwords {
		itemDest := dto.Secret{
			Name:     pwd.Name,
			Metadata: pwd.Metadata,
		}
		pwdDest := dto.LoginPassword{
			Secret:   itemDest,
			Login:    pwd.Login,
			Password: pwd.Password,
		}
		err := s.cryptoService.DecryptSecret(&pwdDest)
		if err != nil {
			return dto.SecretsList{}, err
		}
		secretsList.Passwords = append(secretsList.Passwords, &pwdDest)
	}
	for _, txt := range response.Texts {
		itemDest := dto.Secret{
			Name:     txt.Name,
			Metadata: txt.Metadata,
		}
		txtDest := &dto.TextInfo{
			Text:   txt.Text,
			Secret: itemDest,
		}
		err := s.cryptoService.DecryptSecret(txtDest)
		if err != nil {
			return dto.SecretsList{}, err
		}
		secretsList.Texts = append(secretsList.Texts, txtDest)
	}
	for _, crd := range response.Cards {
		itemDest := dto.Secret{
			Name:     crd.Name,
			Metadata: crd.Metadata,
		}
		crdDest := &dto.CardInfo{
			Secret:          itemDest,
			Number:          crd.Number,
			CVV:             crd.Cvv,
			ExpirationMonth: crd.ExpirationMonth,
			ExpirationYear:  crd.ExpirationYear,
		}
		err := s.cryptoService.DecryptSecret(crdDest)
		if err != nil {
			return dto.SecretsList{}, err
		}
		secretsList.Cards = append(secretsList.Cards, crdDest)
	}

	return secretsList, nil
}
