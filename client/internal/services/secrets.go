package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
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
	AddBinaryInfo(ctx context.Context, binInfo *dto.BinaryInfo) error
	ListSecrets(ctx context.Context, user string) (dto.SecretsList, error)
	AddSecrets(ctx context.Context, secretsList dto.SecretsList) error
	GetCredentialsById(ctx context.Context, id string, user string) (dto.LoginPassword, error)
	GetTextById(ctx context.Context, id string, user string) (dto.TextInfo, error)
	GetBinaryById(ctx context.Context, id string, user string) (dto.BinaryInfo, error)
	GetCardById(ctx context.Context, id string, user string) (dto.CardInfo, error)
	UpdateCardInfo(ctx context.Context, crd *dto.CardInfo) error
	UpdateTextInfo(ctx context.Context, txt *dto.TextInfo) error
	UpdateCredentials(ctx context.Context, pwd *dto.LoginPassword) error
	UpdateBinaryInfo(ctx context.Context, crd *dto.BinaryInfo) error
}

type UpdatePasswordsService interface{}

type DataToSync struct {
	secretsToCreate      []dto.Secret
	localSecretsToCreate []dto.Secret
	secretsToUpdate      []dto.Secret
	localSecretsToUpdate []dto.Secret
}

type SecretsService struct {
	authService        AuthMetadataService
	storageClient      secrets.SecretsManagementClient
	cryptoService      SecretCryptoService
	localStorage       LocalItemsStorage
	userCredsStorage   UserCredentialsStorage
	pwdInstanceService UpdatePasswordsService
	//crdInstanceService UpdateRetrieveCardService
	//txtInstanceService UpdateRetrieveTextService
	//binInstanceService UpdateRetrieveBinaryService
}

func NewSecretsService(
	client secrets.SecretsManagementClient,
	service AuthMetadataService,
	cryptoService SecretCryptoService,
	localStorage LocalItemsStorage,
	userCredsStorage UserCredentialsStorage,
	pwdInstanceService UpdatePasswordsService,
) *SecretsService {
	return &SecretsService{
		storageClient:      client,
		authService:        service,
		cryptoService:      cryptoService,
		localStorage:       localStorage,
		userCredsStorage:   userCredsStorage,
		pwdInstanceService: pwdInstanceService,
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
	request := common.PasswordToProto(loginPwd)

	response, err := s.storageClient.AddCredentials(ctx, request)
	if err != nil && !checkUnavailableError(err) {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding password: %s", response.Error)
	}

	loginPwd.User = credentials.Email
	loginPwd.ID = response.Id
	err = s.localStorage.AddCredentials(ctx, loginPwd)
	return err
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
	request := common.TextToProto(text)

	response, err := s.storageClient.AddTextInfo(ctx, request)
	if err != nil && !checkUnavailableError(err) {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding password: %s", response.Error)
	}

	text.User = credentials.Email
	text.ID = response.Id
	err = s.localStorage.AddTextInfo(ctx, text)
	return err
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
	request := common.CardToProto(card)

	response, err := s.storageClient.AddCardInfo(ctx, request)
	if err != nil && !checkUnavailableError(err) {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding password: %s", response.Error)
	}

	card.User = credentials.Email
	card.ID = response.Id
	err = s.localStorage.AddCardInfo(ctx, card)
	return err
}

func (s *SecretsService) AddBinaryInfo(bin *dto.BinaryInfo) error {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}

	ctx, err := s.authService.AddAuthMetadata(context.Background(), credentials.Token)
	if err != nil {
		return err
	}

	err = s.cryptoService.EncryptSecret(bin)
	if err != nil {
		return fmt.Errorf("can't encrypt item: %s", err)
	}
	request := common.BinaryToProto(bin)

	response, err := s.storageClient.AddBinaryInfo(ctx, request)
	if err != nil {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding binary: %s", response.Error)
	}

	bin.User = credentials.Email
	bin.ID = response.Id
	err = s.localStorage.AddBinaryInfo(ctx, bin)
	return err
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

	// get secrets from remote
	request := secrets.EmptyRequest{}
	response, err := s.storageClient.ListSecrets(ctx, &request)
	if err != nil && !checkUnavailableError(err) {
		return dto.SecretsList{}, err
	}

	var secretsList dto.SecretsList
	for _, pwd := range response.Passwords {
		pwdDest := common.PasswordFromProto(pwd)
		secretsList.Passwords = append(secretsList.Passwords, &pwdDest)
	}
	for _, txt := range response.Texts {
		txtDest := common.TextFromProto(txt)
		secretsList.Texts = append(secretsList.Texts, &txtDest)
	}
	for _, crd := range response.Cards {
		crdDest := common.CardFromProto(crd)
		secretsList.Cards = append(secretsList.Cards, &crdDest)
	}
	for _, bin := range response.Bins {
		dest := common.BinaryFromProto(bin)
		secretsList.Bins = append(secretsList.Bins, &dest)
	}

	//get secrets from local storage
	localSecretsList, err := s.localStorage.ListSecrets(ctx, credentials.Email)

	// sync secrets
	//s.syncSecrets(secretsList.Passwords, localSecretsList.Passwords)

	//get secrets again from local storage
	localSecretsList, err = s.localStorage.ListSecrets(ctx, credentials.Email)
	// decrypt them and return to user
	for _, pwd := range localSecretsList.Passwords {
		err := s.cryptoService.DecryptSecret(pwd)
		if err != nil {
			return dto.SecretsList{}, err
		}
	}
	for _, txt := range localSecretsList.Texts {
		err := s.cryptoService.DecryptSecret(txt)
		if err != nil {
			return dto.SecretsList{}, err
		}
	}
	for _, crd := range localSecretsList.Cards {
		err := s.cryptoService.DecryptSecret(crd)
		if err != nil {
			return dto.SecretsList{}, err
		}
	}
	for _, bin := range localSecretsList.Bins {
		err := s.cryptoService.DecryptSecret(bin)
		if err != nil {
			return dto.SecretsList{}, err
		}
	}

	return localSecretsList, nil
}

func (s *SecretsService) LoadSecrets(ctx context.Context) error {
	secretsList, err := s.ListSecrets()
	if err != nil {
		return err
	}

	return s.localStorage.AddSecrets(ctx, secretsList)
}

// todo: change to base secret
func (s *SecretsService) syncSecrets(secrets []dto.Secret, localSecrets []dto.Secret) DataToSync {
	var dataToSync DataToSync
	secretsMap := make(map[string]dto.Secret, len(secrets))
	localSecretsMap := make(map[string]dto.Secret, len(localSecrets))
	for _, secret := range secrets {
		secretsMap[secret.GetID()] = secret
	}
	for _, secret := range localSecrets {
		secretID := secret.GetID()
		localSecretsMap[secretID] = secret
		remoteSecret, ok := secretsMap[secretID]
		if !ok {
			dataToSync.localSecretsToCreate = append(dataToSync.localSecretsToCreate, secret)
		} else {
			if remoteSecret.GetUpdatedAt() == secret.GetUpdatedAt() {
				continue
			}
			if remoteSecret.GetUpdatedAt().After(secret.GetUpdatedAt()) {
				dataToSync.localSecretsToUpdate = append(dataToSync.localSecretsToUpdate, remoteSecret)
			}
			dataToSync.secretsToUpdate = append(dataToSync.secretsToUpdate, secret)
		}
	}
	for _, secret := range secrets {
		secretID := secret.GetID()
		if _, ok := localSecretsMap[secretID]; !ok {
			dataToSync.secretsToCreate = append(dataToSync.secretsToCreate, secret)
		}
	}

	return dataToSync
}

func (s *SecretsService) addCredentials(authCtx context.Context, loginPwd *dto.LoginPassword) (*secrets.AddResponse, error) {
	err := s.cryptoService.EncryptSecret(loginPwd)
	if err != nil {
		return nil, fmt.Errorf("can't encrypt item: %s", err)
	}
	request := common.PasswordToProto(loginPwd)

	return s.storageClient.AddCredentials(authCtx, request)
}
