package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthMetadataService interface {
	AddAuthMetadata(ctx context.Context, token string) (context.Context, error)
}

type SecretCryptoService interface {
	EncryptSecret(source any) error
	DecryptSecret(source any) error
}

type UserCredentialsStorage interface {
	GetCredentials() (*dto.UserCredentials, error)
}

type LocalItemsStorage interface {
	AddCredentials(ctx context.Context, password *dto.LoginPassword) error
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error
	AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error
	ListSecrets(ctx context.Context, user string) (dto.SecretsList, error)
	AddSecrets(ctx context.Context, secretsList dto.SecretsList) error
}

type SecretsService struct {
	authService      AuthMetadataService
	storageClient    secrets.SecretsManagementClient
	cryptoService    SecretCryptoService
	localStorage     LocalItemsStorage
	userCredsStorage UserCredentialsStorage
}

func (s *SecretsService) GetPasswordByID(id string) (*dto.LoginPassword, error) {
	//TODO implement me
	panic("implement me")
}

func NewSecretsService(
	client secrets.SecretsManagementClient,
	service AuthMetadataService,
	cryptoService SecretCryptoService,
	localStorage LocalItemsStorage,
	userCredsStorage UserCredentialsStorage,
) *SecretsService {
	return &SecretsService{
		storageClient:    client,
		authService:      service,
		cryptoService:    cryptoService,
		localStorage:     localStorage,
		userCredsStorage: userCredsStorage,
	}
}

func (s *SecretsService) AddCredentials(loginPwd *dto.LoginPassword) error {
	// todo: context should be passed from argument
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}

	ctx, err := s.authService.AddAuthMetadata(context.Background(), credentials.Token)
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
	st, ok := status.FromError(err)
	if ok && st.Code() == codes.Unavailable {
		log.Error().Err(err).Msg("remote storage is not available:")
		return nil
	}
	if err != nil {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding password: %s", response.Error)
	}

	loginPwd.User = credentials.Email
	loginPwd.ID = response.Id
	err = s.localStorage.AddCredentials(ctx, loginPwd)
	if err != nil {
		return err
	}

	return nil
}

func (s *SecretsService) AddTextInfo(text *dto.TextInfo) error {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}
	ctx, err := s.authService.AddAuthMetadata(context.Background(), credentials.Token)
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
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}

	ctx, err := s.authService.AddAuthMetadata(context.Background(), credentials.Token)
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
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return dto.SecretsList{}, err
	}

	ctx, err := s.authService.AddAuthMetadata(context.Background(), credentials.Token)
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
			ID:       pwd.Id,
			Name:     pwd.Name,
			Metadata: pwd.Metadata,
			User:     pwd.User,
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
			ID:       txt.Id,
			Name:     txt.Name,
			Metadata: txt.Metadata,
			User:     txt.User,
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
			ID:       crd.Id,
			Name:     crd.Name,
			Metadata: crd.Metadata,
			User:     crd.User,
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

func (s *SecretsService) LoadSecrets(ctx context.Context) error {
	secretsList, err := s.ListSecrets()
	if err != nil {
		return err
	}

	return s.localStorage.AddSecrets(ctx, secretsList)
}
