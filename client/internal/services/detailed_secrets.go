package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"time"
)

type UpdateRetrieveDeleteSecretServiceInterface[T dto.Secret] interface {
	GetSecretByID(id string) (T, error)
	UpdateSecret(secret T) error
}

type SecretInstanceService[T dto.Secret] interface {
	GetSecretByID(ctx context.Context, id string) (T, error)
	GetLocalSecretByID(id string, email string) (T, error)
	UpdateSecret(ctx context.Context, secret T) error
	UpdateLocalSecret(secret T) error
}

type UpdateRetrieveDeleteSecretService[T dto.Secret] struct {
	secretInstanceService SecretInstanceService[T]
	authService           AuthMetadataService
	cryptoService         SecretCryptoService
	userCredsStorage      UserCredentialsStorage
}

func NewUpdateRetrieveDeletePasswordService(
	authService AuthMetadataService,
	cryptoService SecretCryptoService,
	userCredsStorage UserCredentialsStorage,
	localStorage LocalItemsStorage,
	client secrets.SecretsManagementClient,
) UpdateRetrieveDeleteSecretServiceInterface[dto.LoginPassword] {
	secretInstanceService := &PasswordInstanceService{
		storageClient: client,
		localStorage:  localStorage,
	}
	return &UpdateRetrieveDeleteSecretService[dto.LoginPassword]{
		secretInstanceService: secretInstanceService,
		authService:           authService,
		cryptoService:         cryptoService,
		userCredsStorage:      userCredsStorage,
	}
}

func NewUpdateRetrieveDeleteCardService(
	authService AuthMetadataService,
	cryptoService SecretCryptoService,
	userCredsStorage UserCredentialsStorage,
	localStorage LocalItemsStorage,
	client secrets.SecretsManagementClient,
) UpdateRetrieveDeleteSecretServiceInterface[dto.CardInfo] {
	secretInstanceService := &CardInstanceService{
		storageClient: client,
		localStorage:  localStorage,
	}
	return &UpdateRetrieveDeleteSecretService[dto.CardInfo]{
		secretInstanceService: secretInstanceService,
		authService:           authService,
		cryptoService:         cryptoService,
		userCredsStorage:      userCredsStorage,
	}
}

func NewUpdateRetrieveDeleteTextService(
	authService AuthMetadataService,
	cryptoService SecretCryptoService,
	userCredsStorage UserCredentialsStorage,
	localStorage LocalItemsStorage,
	client secrets.SecretsManagementClient,
) UpdateRetrieveDeleteSecretServiceInterface[dto.TextInfo] {
	secretInstanceService := &TextInstanceService{
		storageClient: client,
		localStorage:  localStorage,
	}
	return &UpdateRetrieveDeleteSecretService[dto.TextInfo]{
		secretInstanceService: secretInstanceService,
		authService:           authService,
		cryptoService:         cryptoService,
		userCredsStorage:      userCredsStorage,
	}
}

func NewUpdateRetrieveDeleteBinaryService(
	authService AuthMetadataService,
	cryptoService SecretCryptoService,
	userCredsStorage UserCredentialsStorage,
	localStorage LocalItemsStorage,
	client secrets.SecretsManagementClient,
) UpdateRetrieveDeleteSecretServiceInterface[dto.BinaryInfo] {
	secretInstanceService := &BinaryInstanceService{
		storageClient: client,
		localStorage:  localStorage,
	}
	return &UpdateRetrieveDeleteSecretService[dto.BinaryInfo]{
		secretInstanceService: secretInstanceService,
		authService:           authService,
		cryptoService:         cryptoService,
		userCredsStorage:      userCredsStorage,
	}
}

func (s *UpdateRetrieveDeleteSecretService[T]) GetSecretByID(id string) (T, error) {
	var secret T

	// first get secret from remote storage
	secret, err := s.getSecretById(id)
	if err != nil && !checkUnavailableError(err) {
		return secret, err
	}

	// get secret from local storage
	localSecret, err := s.getLocalSecretByID(id)

	// compare updated_at timestamp
	//return secret with the latest timestamp
	if localSecret.GetUpdatedAt().After(secret.GetUpdatedAt()) {
		secret = localSecret
	}

	return secret, nil
}

func (s *UpdateRetrieveDeleteSecretService[T]) UpdateSecret(secret T) error {
	err := s.cryptoService.EncryptSecret(&secret)
	if err != nil {
		return fmt.Errorf("can't encrypt item: %s", err)
	}
	secret.SetUpdatedAt(time.Now())

	// do update in remote
	ctx, err := s.authenticateContext(context.Background())
	if err != nil {
		return err
	}

	err = s.secretInstanceService.UpdateSecret(ctx, secret)
	if err != nil && !checkUnavailableError(err) {
		return err
	}

	// do update in local storage
	err = s.secretInstanceService.UpdateLocalSecret(secret)
	if err != nil {
		return err
	}

	return nil
}

func (s *UpdateRetrieveDeleteSecretService[T]) getSecretById(id string) (T, error) {
	var secret T

	ctx, err := s.authenticateContext(context.Background())
	if err != nil {
		return secret, err
	}

	secret, err = s.secretInstanceService.GetSecretByID(ctx, id)
	if err != nil {
		return secret, err
	}

	err = s.cryptoService.DecryptSecret(&secret)
	if err != nil {
		return secret, err
	}

	return secret, nil
}

func (s *UpdateRetrieveDeleteSecretService[T]) getLocalSecretByID(id string) (T, error) {
	var secret T

	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return secret, err
	}

	secret, err = s.secretInstanceService.GetLocalSecretByID(id, credentials.Email)
	if err != nil {
		return secret, err
	}

	err = s.cryptoService.DecryptSecret(&secret)
	if err != nil {
		return secret, err
	}

	return secret, nil
}

func (s *UpdateRetrieveDeleteSecretService[T]) authenticateContext(ctx context.Context) (context.Context, error) {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return nil, err
	}

	ctx, err = s.authService.AddAuthMetadata(context.Background(), credentials.Token)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}
