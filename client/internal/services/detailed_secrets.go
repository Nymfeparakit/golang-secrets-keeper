package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateRetrieveDeleteSecretServiceInterface service to perform update/retrieve operations with local and remote storage.
type UpdateRetrieveDeleteSecretServiceInterface[T dto.Secret] interface {
	GetSecretByID(ctx context.Context, id string) (T, error)
	UpdateSecret(ctx context.Context, secret T) error
	UpdateRemoteSecret(ctx context.Context, secret T) error
	UpdateLocalSecret(ctx context.Context, secret T) error
	DeleteSecret(ctx context.Context, id string) error
	DeleteRemoteSecret(ctx context.Context, id string) error
	DeleteLocalSecret(ctx context.Context, id string) error
}

// SecretInstanceService - service to perform operations with single secret instance.
type SecretInstanceService[T dto.Secret] interface {
	GetSecretByID(ctx context.Context, id string) (T, error)
	GetLocalSecretByID(ctx context.Context, id string, email string) (T, error)
	UpdateSecret(ctx context.Context, secret T) error
	UpdateLocalSecret(ctx context.Context, secret T) error
	DeleteSecret(ctx context.Context, id string) error
	DeleteLocalSecret(ctx context.Context, id string) error
}

// UpdateRetrieveDeleteSecretService - service to perform update/retrieve operations with local and remote storage.
type UpdateRetrieveDeleteSecretService[T dto.Secret] struct {
	secretInstanceService SecretInstanceService[T]
	authService           AuthMetadataService
	cryptoService         SecretCryptoService
	userCredsStorage      UserCredentialsStorage
}

// NewUpdateRetrieveDeletePasswordService creates service to perform update/retrieve operations with LoginPassword instance.
func NewUpdateRetrieveDeletePasswordService(
	authService AuthMetadataService,
	cryptoService SecretCryptoService,
	userCredsStorage UserCredentialsStorage,
	localStorage CredentialsLocalStorage,
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

// NewUpdateRetrieveDeleteCardService creates service to perform update/retrieve operations with CardInfo instance.
func NewUpdateRetrieveDeleteCardService(
	authService AuthMetadataService,
	cryptoService SecretCryptoService,
	userCredsStorage UserCredentialsStorage,
	localStorage CardsLocalStorage,
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

// NewUpdateRetrieveDeleteTextService creates service to perform update/retrieve operations with TextInfo instance.
func NewUpdateRetrieveDeleteTextService(
	authService AuthMetadataService,
	cryptoService SecretCryptoService,
	userCredsStorage UserCredentialsStorage,
	localStorage TextLocalStorage,
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

// NewUpdateRetrieveDeleteBinaryService creates service to perform update/retrieve operations with BinaryInfo instance.
func NewUpdateRetrieveDeleteBinaryService(
	authService AuthMetadataService,
	cryptoService SecretCryptoService,
	userCredsStorage UserCredentialsStorage,
	localStorage BinaryLocalStorage,
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

// GetSecretByID gets secret from remote and local storage, then compares them and returns the one with the latest
// UpdatedAt timestamp.
func (s *UpdateRetrieveDeleteSecretService[T]) GetSecretByID(ctx context.Context, id string) (T, error) {
	var secret T

	// first get secret from remote storage
	secret, err := s.getSecretById(ctx, id)
	gotSecretFromRemote := true
	if err != nil {
		if !checkUnavailableError(err) {
			return secret, err
		}
		gotSecretFromRemote = false
	}
	// check if it has been deleted
	if gotSecretFromRemote && secret.IsDeleted() {
		return secret, ErrSecretDoesNotExist
	}

	// get secret from local storage
	localSecret, err := s.getLocalSecretByID(ctx, id)
	gotLocalSecret := true
	if err != nil {
		if err != storage.ErrSecretNotFound {
			return secret, fmt.Errorf("can't get secret from local storage")
		}
		gotLocalSecret = false
	}
	// check if it has been deleted
	if gotLocalSecret && localSecret.IsDeleted() {
		return secret, ErrSecretDoesNotExist
	}

	// if there's no such secret on remote and in local storage
	if !gotLocalSecret && !gotSecretFromRemote {
		return secret, storage.ErrSecretNotFound
	}
	// if such a secret exists only in local storage
	if !gotSecretFromRemote {
		return localSecret, nil
	}
	if !gotLocalSecret {
		return secret, nil
	}
	// compare updated_at timestamp
	//return secret with the latest timestamp
	// if local secret was updated later
	if localSecret.GetUpdatedAt().After(secret.GetUpdatedAt()) {
		// update old secret in remote
		err = s.UpdateRemoteSecret(ctx, localSecret)
		if err != nil && !checkUnavailableError(err) {
			return secret, nil
		}
		secret = localSecret
	} else {
		// update old secret locally
		err = s.UpdateLocalSecret(ctx, secret)
		if err != nil {
			return secret, err
		}
	}

	err = s.cryptoService.DecryptSecret(&secret)
	if err != nil {
		return secret, err
	}

	return secret, nil
}

// UpdateSecret updates secret in remote and local storage.
func (s *UpdateRetrieveDeleteSecretService[T]) UpdateSecret(ctx context.Context, secret T) error {
	err := s.cryptoService.EncryptSecret(&secret)
	if err != nil {
		return fmt.Errorf("can't encrypt item: %s", err)
	}

	// do update in remote
	err = s.UpdateRemoteSecret(ctx, secret)
	if err != nil && !checkUnavailableError(err) {
		return err
	}

	// do update in local storage
	err = s.UpdateLocalSecret(ctx, secret)

	return err
}

func (s *UpdateRetrieveDeleteSecretService[T]) UpdateRemoteSecret(ctx context.Context, secret T) error {
	ctx, err := s.withContext(ctx)
	if err != nil {
		return err
	}

	err = s.secretInstanceService.UpdateSecret(ctx, secret)
	return err
}

func (s *UpdateRetrieveDeleteSecretService[T]) UpdateLocalSecret(ctx context.Context, secret T) error {
	err := s.secretInstanceService.UpdateLocalSecret(ctx, secret)
	return err
}

func (s *UpdateRetrieveDeleteSecretService[T]) withContext(ctx context.Context) (context.Context, error) {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return nil, err
	}

	ctx, err = s.authService.AddAuthMetadata(ctx, credentials.Token)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}

// DeleteSecret deletes secret in remote and local storage.
func (s *UpdateRetrieveDeleteSecretService[T]) DeleteSecret(ctx context.Context, id string) error {
	// try to delete in remote
	err := s.DeleteRemoteSecret(ctx, id)
	st, ok := status.FromError(err)
	deletedFromRemote := true
	if err != nil {
		deletedFromRemote = false
		// failed precondition can be in case when secret was already deleted from remote
		if ok && st.Code() != codes.FailedPrecondition && !checkUnavailableError(err) {
			return err
		}
	}

	// try to delete locally
	deletedFromLocal := true
	err = s.DeleteLocalSecret(ctx, id)
	if err != nil {
		deletedFromLocal = false
		if !errors.Is(err, storage.ErrSecretDoesNotExistOrWasDeleted) {
			return err
		}
	}

	if !deletedFromRemote && !deletedFromLocal {
		return fmt.Errorf("secret with provided id doesn't exist")
	}

	return nil
}

func (s *UpdateRetrieveDeleteSecretService[T]) DeleteRemoteSecret(ctx context.Context, id string) error {
	ctx, err := s.withContext(ctx)
	if err != nil {
		return err
	}
	err = s.secretInstanceService.DeleteSecret(ctx, id)
	return err
}

func (s *UpdateRetrieveDeleteSecretService[T]) DeleteLocalSecret(ctx context.Context, id string) error {
	err := s.secretInstanceService.DeleteLocalSecret(ctx, id)
	return err
}

func (s *UpdateRetrieveDeleteSecretService[T]) getSecretById(ctx context.Context, id string) (T, error) {
	var secret T

	ctx, err := s.withContext(ctx)
	if err != nil {
		return secret, err
	}

	secret, err = s.secretInstanceService.GetSecretByID(ctx, id)
	if err != nil {
		return secret, err
	}

	return secret, nil
}

func (s *UpdateRetrieveDeleteSecretService[T]) getLocalSecretByID(ctx context.Context, id string) (T, error) {
	var secret T

	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return secret, err
	}

	secret, err = s.secretInstanceService.GetLocalSecretByID(ctx, id, credentials.Email)
	if err != nil {
		return secret, err
	}

	return secret, nil
}
