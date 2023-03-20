package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"golang.org/x/sync/errgroup"
)

// AuthMetadataService service for adding auth metadata to provided context.
type AuthMetadataService interface {
	AddAuthMetadata(ctx context.Context, token string) (context.Context, error)
}

// SecretCryptoService service for encrypting and decrypting secrets.
type SecretCryptoService interface {
	EncryptSecret(source any) error
	DecryptSecret(source any) error
}

// UserCredentialsStorage - storage of user credentials (email and token).
type UserCredentialsStorage interface {
	GetCredentials() (*dto.UserCredentials, error)
}

// LocalSecretsStorage - local storage of secrets.
type LocalSecretsStorage interface {
	AddCredentials(ctx context.Context, password *dto.LoginPassword) (string, error)
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) (string, error)
	AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) (string, error)
	AddBinaryInfo(ctx context.Context, binInfo *dto.BinaryInfo) (string, error)

	ListSecrets(ctx context.Context, user string) (dto.SecretsList, error)
	AddSecrets(ctx context.Context, secretsList dto.SecretsList) error
}

// UpdateDeletePasswordService - service for updating LoginPassword instance.
type UpdateDeletePasswordService interface {
	UpdateRemoteSecret(ctx context.Context, secret dto.LoginPassword) error
	UpdateLocalSecret(ctx context.Context, secret dto.LoginPassword) error
	DeleteRemoteSecret(ctx context.Context, id string) error
	DeleteLocalSecret(ctx context.Context, id string) error
}

// UpdateDeleteCardService - service for updating CardInfo instance.
type UpdateDeleteCardService interface {
	UpdateRemoteSecret(ctx context.Context, secret dto.CardInfo) error
	UpdateLocalSecret(ctx context.Context, secret dto.CardInfo) error
	DeleteRemoteSecret(ctx context.Context, id string) error
	DeleteLocalSecret(ctx context.Context, id string) error
}

// UpdateDeleteTextService - service for updating TextInfo instance.
type UpdateDeleteTextService interface {
	UpdateRemoteSecret(ctx context.Context, secret dto.TextInfo) error
	UpdateLocalSecret(ctx context.Context, secret dto.TextInfo) error
	DeleteRemoteSecret(ctx context.Context, id string) error
	DeleteLocalSecret(ctx context.Context, id string) error
}

// UpdateDeleteBinaryService - service for updating BinaryInfo instance.
type UpdateDeleteBinaryService interface {
	UpdateRemoteSecret(ctx context.Context, secret dto.BinaryInfo) error
	UpdateLocalSecret(ctx context.Context, secret dto.BinaryInfo) error
	DeleteRemoteSecret(ctx context.Context, id string) error
	DeleteLocalSecret(ctx context.Context, id string) error
}

// DataToSync - data to sync between remote and local storage.
type DataToSync struct {
	localUniqID    []string
	remoteUniqID   []string
	intersectionID map[string]bool
}

// SecretsService service for listing and adding new secrets.
type SecretsService struct {
	authService        AuthMetadataService
	storageClient      secrets.SecretsManagementClient
	cryptoService      SecretCryptoService
	localStorage       LocalSecretsStorage
	userCredsStorage   UserCredentialsStorage
	pwdInstanceService UpdateDeletePasswordService
	crdInstanceService UpdateDeleteCardService
	txtInstanceService UpdateDeleteTextService
	binInstanceService UpdateDeleteBinaryService
}

// NewSecretsService creates new SecretsService object.
func NewSecretsService(
	client secrets.SecretsManagementClient,
	service AuthMetadataService,
	cryptoService SecretCryptoService,
	localStorage LocalSecretsStorage,
	userCredsStorage UserCredentialsStorage,
	pwdInstanceService UpdateDeletePasswordService,
	crdInstanceService UpdateDeleteCardService,
	txtInstanceService UpdateDeleteTextService,
	binInstanceService UpdateDeleteBinaryService,
) *SecretsService {
	return &SecretsService{
		storageClient:      client,
		authService:        service,
		cryptoService:      cryptoService,
		localStorage:       localStorage,
		userCredsStorage:   userCredsStorage,
		pwdInstanceService: pwdInstanceService,
		crdInstanceService: crdInstanceService,
		txtInstanceService: txtInstanceService,
		binInstanceService: binInstanceService,
	}
}

// AddCredentials adds LoginPwd to remote and local storage.
func (s *SecretsService) AddCredentials(ctx context.Context, loginPwd *dto.LoginPassword) error {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}
	ctx, err = s.authService.AddAuthMetadata(ctx, credentials.Token)
	if err != nil {
		return err
	}

	err = s.cryptoService.EncryptSecret(loginPwd)
	if err != nil {
		return fmt.Errorf("can't encrypt item: %s", err)
	}
	request := common.CredentialsToProto(loginPwd)

	response, err := s.storageClient.AddCredentials(ctx, request)
	if err != nil && !checkUnavailableError(err) {
		return err
	}

	loginPwd.User = credentials.Email
	// if remote was available
	if response != nil {
		loginPwd.ID = response.Id
	}
	_, err = s.localStorage.AddCredentials(ctx, loginPwd)
	return err
}

// AddTextInfo adds TextInfo to remote and local storage.
func (s *SecretsService) AddTextInfo(ctx context.Context, text *dto.TextInfo) error {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}
	ctx, err = s.authService.AddAuthMetadata(ctx, credentials.Token)
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

	text.User = credentials.Email
	// if remote was available
	if response != nil {
		text.ID = response.Id
	}
	_, err = s.localStorage.AddTextInfo(ctx, text)
	return err
}

// AddCardInfo adds CardInfo to remote and local storage.
func (s *SecretsService) AddCardInfo(ctx context.Context, card *dto.CardInfo) error {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}

	ctx, err = s.authService.AddAuthMetadata(ctx, credentials.Token)
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

	card.User = credentials.Email
	if response != nil {
		card.ID = response.Id
	}
	_, err = s.localStorage.AddCardInfo(ctx, card)
	return err
}

// AddBinaryInfo adds BinaryInfo to remote and local storage.
func (s *SecretsService) AddBinaryInfo(ctx context.Context, bin *dto.BinaryInfo) error {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}

	ctx, err = s.authService.AddAuthMetadata(ctx, credentials.Token)
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

	bin.User = credentials.Email
	if response != nil {
		bin.ID = response.Id
	}
	_, err = s.localStorage.AddBinaryInfo(ctx, bin)
	return err
}

type SecretsMap struct {
	passwords map[string]*dto.LoginPassword
	cards     map[string]*dto.CardInfo
	texts     map[string]*dto.TextInfo
	bins      map[string]*dto.BinaryInfo
}

// ListSecrets syncs all secrets for current user between local and remote storage, and then returns synchronized
// secrets list.
func (s *SecretsService) ListSecrets(ctx context.Context) (dto.SecretsList, error) {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return dto.SecretsList{}, err
	}

	ctx, err = s.authService.AddAuthMetadata(ctx, credentials.Token)
	if err != nil {
		return dto.SecretsList{}, err
	}

	// get secrets from remote
	secretsListRemote, err := s.listSecrets(ctx)
	if err != nil && !checkUnavailableError(err) {
		return dto.SecretsList{}, err
	}

	secretsMap := SecretsMap{
		passwords: make(map[string]*dto.LoginPassword, len(secretsListRemote.Passwords)),
		cards:     make(map[string]*dto.CardInfo, len(secretsListRemote.Cards)),
		texts:     make(map[string]*dto.TextInfo, len(secretsListRemote.Texts)),
		bins:      make(map[string]*dto.BinaryInfo, len(secretsListRemote.Bins)),
	}
	for _, pwd := range secretsListRemote.Passwords {
		secretsMap.passwords[pwd.ID] = pwd
	}
	for _, txt := range secretsListRemote.Texts {
		secretsMap.texts[txt.ID] = txt
	}
	for _, crd := range secretsListRemote.Cards {
		secretsMap.cards[crd.ID] = crd
	}
	for _, bin := range secretsListRemote.Bins {
		secretsMap.bins[bin.ID] = bin
	}

	//get secrets from local storage
	localSecretsList, err := s.localStorage.ListSecrets(ctx, credentials.Email)

	localSecretsMap := SecretsMap{
		passwords: make(map[string]*dto.LoginPassword, len(localSecretsList.Passwords)),
		cards:     make(map[string]*dto.CardInfo, len(localSecretsList.Cards)),
		texts:     make(map[string]*dto.TextInfo, len(localSecretsList.Texts)),
		bins:      make(map[string]*dto.BinaryInfo, len(localSecretsList.Bins)),
	}
	for _, pwd := range localSecretsList.Passwords {
		localSecretsMap.passwords[pwd.ID] = pwd
	}
	for _, txt := range localSecretsList.Texts {
		localSecretsMap.texts[txt.ID] = txt
	}
	for _, crd := range localSecretsList.Cards {
		localSecretsMap.cards[crd.ID] = crd
	}
	for _, bin := range localSecretsList.Bins {
		localSecretsMap.bins[bin.ID] = bin
	}

	finalSecretsList := localSecretsList
	if len(secretsListRemote.Passwords) != 0 || len(secretsListRemote.Texts) != 0 || len(secretsListRemote.Cards) != 0 || len(secretsListRemote.Bins) != 0 {
		// sync secrets to get final list of all of them
		finalSecretsList, err = s.syncSecrets(ctx, secretsMap, localSecretsMap)
		if err != nil {
			return dto.SecretsList{}, fmt.Errorf("failed to sync data: %v", err)
		}
	}
	// decrypt all secrets
	for _, secret := range finalSecretsList.Passwords {
		err := s.cryptoService.DecryptSecret(secret)
		if err != nil {
			return dto.SecretsList{}, err
		}
	}
	for _, secret := range finalSecretsList.Texts {
		err := s.cryptoService.DecryptSecret(secret)
		if err != nil {
			return dto.SecretsList{}, err
		}
	}
	for _, secret := range finalSecretsList.Cards {
		err := s.cryptoService.DecryptSecret(secret)
		if err != nil {
			return dto.SecretsList{}, err
		}
	}
	for _, secret := range finalSecretsList.Bins {
		err := s.cryptoService.DecryptSecret(secret)
		if err != nil {
			return dto.SecretsList{}, err
		}
	}

	return finalSecretsList, nil
}

// LoadSecrets gets all user secrets from remote and then adds them to local storage.
func (s *SecretsService) LoadSecrets(ctx context.Context) error {
	// load secrets from remote
	secretsList, err := s.listSecrets(ctx)
	if err != nil {
		return err
	}

	return s.localStorage.AddSecrets(ctx, secretsList)
}

func (s *SecretsService) listSecrets(ctx context.Context) (dto.SecretsList, error) {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return dto.SecretsList{}, err
	}

	ctx, err = s.authService.AddAuthMetadata(ctx, credentials.Token)
	if err != nil {
		return dto.SecretsList{}, err
	}

	// get secrets from remote
	request := secrets.EmptyRequest{}
	response, err := s.storageClient.ListSecrets(ctx, &request)
	if err != nil && !checkUnavailableError(err) {
		return dto.SecretsList{}, err
	}
	if response == nil {
		return dto.SecretsList{}, nil
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

	return secretsList, nil
}

func (s *SecretsService) syncSecrets(ctx context.Context, secrets SecretsMap, localSecrets SecretsMap) (dto.SecretsList, error) {
	wg := new(errgroup.Group)
	var secretsList dto.SecretsList

	pwdID := make(map[string]bool, len(secrets.passwords))
	localPwdID := make(map[string]bool, len(secrets.passwords))
	for ID := range secrets.passwords {
		pwdID[ID] = true
	}
	for ID := range localSecrets.passwords {
		localPwdID[ID] = true
	}
	dataToSync := s.findAllAndUniqIDs(pwdID, localPwdID)

	for _, ID := range dataToSync.localUniqID {
		if !localSecrets.passwords[ID].IsDeleted() {
			secretsList.Passwords = append(secretsList.Passwords, localSecrets.passwords[ID])
		}
		wg.Go(func() error {
			request := common.CredentialsToProto(localSecrets.passwords[ID])
			_, err := s.storageClient.AddCredentials(ctx, request)
			return err
		})
	}
	for _, ID := range dataToSync.remoteUniqID {
		if !secrets.passwords[ID].IsDeleted() {
			secretsList.Passwords = append(secretsList.Passwords, secrets.passwords[ID])
		}
		wg.Go(func() error {
			_, err := s.localStorage.AddCredentials(ctx, secrets.passwords[ID])
			return err
		})
	}
	for ID := range dataToSync.intersectionID {
		secret := secrets.passwords[ID]
		localSecret := localSecrets.passwords[ID]
		s.syncIntersectingSecrets(
			wg,
			*secret,
			*localSecret,
			func() error { return s.pwdInstanceService.DeleteLocalSecret(ctx, ID) },
			func() error { return s.pwdInstanceService.DeleteRemoteSecret(ctx, ID) },
			func() error { return s.pwdInstanceService.UpdateLocalSecret(ctx, *secret) },
			func() error { return s.pwdInstanceService.UpdateRemoteSecret(ctx, *localSecret) },
			func() { secretsList.Passwords = append(secretsList.Passwords, localSecret) },
			func() { secretsList.Passwords = append(secretsList.Passwords, secret) },
		)
	}

	txtID := make(map[string]bool, len(secrets.texts))
	localTxtID := make(map[string]bool, len(secrets.texts))
	for ID := range secrets.texts {
		txtID[ID] = true
	}
	for ID := range localSecrets.texts {
		localTxtID[ID] = true
	}
	dataToSync = s.findAllAndUniqIDs(txtID, localTxtID)

	for _, ID := range dataToSync.localUniqID {
		if !localSecrets.texts[ID].IsDeleted() {
			secretsList.Texts = append(secretsList.Texts, localSecrets.texts[ID])
		}
		wg.Go(func() error {
			request := common.TextToProto(localSecrets.texts[ID])
			_, err := s.storageClient.AddTextInfo(ctx, request)
			return err
		})
	}
	for _, ID := range dataToSync.remoteUniqID {
		if !secrets.texts[ID].IsDeleted() {
			secretsList.Texts = append(secretsList.Texts, secrets.texts[ID])
		}
		wg.Go(func() error {
			_, err := s.localStorage.AddTextInfo(ctx, secrets.texts[ID])
			return err
		})
	}
	for ID := range dataToSync.intersectionID {
		secret := secrets.texts[ID]
		localSecret := localSecrets.texts[ID]
		s.syncIntersectingSecrets(
			wg,
			*secret,
			*localSecret,
			func() error { return s.txtInstanceService.DeleteLocalSecret(ctx, ID) },
			func() error { return s.txtInstanceService.DeleteRemoteSecret(ctx, ID) },
			func() error { return s.txtInstanceService.UpdateLocalSecret(ctx, *secret) },
			func() error { return s.txtInstanceService.UpdateRemoteSecret(ctx, *localSecret) },
			func() { secretsList.Texts = append(secretsList.Texts, localSecret) },
			func() { secretsList.Texts = append(secretsList.Texts, secret) },
		)
	}

	crdID := make(map[string]bool, len(secrets.cards))
	localCrdID := make(map[string]bool, len(secrets.cards))
	for ID := range secrets.cards {
		crdID[ID] = true
	}
	for ID := range localSecrets.cards {
		localCrdID[ID] = true
	}
	dataToSync = s.findAllAndUniqIDs(crdID, localCrdID)

	for _, ID := range dataToSync.localUniqID {
		if !localSecrets.cards[ID].IsDeleted() {
			secretsList.Cards = append(secretsList.Cards, localSecrets.cards[ID])
		}
		wg.Go(func() error {
			request := common.CardToProto(localSecrets.cards[ID])
			_, err := s.storageClient.AddCardInfo(ctx, request)
			return err
		})
	}
	for _, ID := range dataToSync.remoteUniqID {
		if !secrets.cards[ID].IsDeleted() {
			secretsList.Cards = append(secretsList.Cards, secrets.cards[ID])
		}
		wg.Go(func() error {
			_, err := s.localStorage.AddCardInfo(ctx, secrets.cards[ID])
			return err
		})
	}
	for ID := range dataToSync.intersectionID {
		secret := secrets.cards[ID]
		localSecret := localSecrets.cards[ID]
		s.syncIntersectingSecrets(
			wg,
			*secret,
			*localSecret,
			func() error { return s.crdInstanceService.DeleteLocalSecret(ctx, ID) },
			func() error { return s.crdInstanceService.DeleteRemoteSecret(ctx, ID) },
			func() error { return s.crdInstanceService.UpdateLocalSecret(ctx, *secret) },
			func() error { return s.crdInstanceService.UpdateRemoteSecret(ctx, *localSecret) },
			func() { secretsList.Cards = append(secretsList.Cards, localSecret) },
			func() { secretsList.Cards = append(secretsList.Cards, secret) },
		)
	}

	binID := make(map[string]bool, len(secrets.bins))
	localBinID := make(map[string]bool, len(secrets.bins))
	for ID := range secrets.bins {
		binID[ID] = true
	}
	for ID := range localSecrets.bins {
		localBinID[ID] = true
	}
	dataToSync = s.findAllAndUniqIDs(binID, localBinID)

	for _, ID := range dataToSync.localUniqID {
		if !localSecrets.bins[ID].IsDeleted() {
			secretsList.Bins = append(secretsList.Bins, localSecrets.bins[ID])
		}
		wg.Go(func() error {
			request := common.BinaryToProto(localSecrets.bins[ID])
			_, err := s.storageClient.AddBinaryInfo(ctx, request)
			return err
		})
	}
	for _, ID := range dataToSync.remoteUniqID {
		if !secrets.bins[ID].IsDeleted() {
			secretsList.Bins = append(secretsList.Bins, secrets.bins[ID])
		}
		wg.Go(func() error {
			_, err := s.localStorage.AddBinaryInfo(ctx, secrets.bins[ID])
			return err
		})
	}
	for ID := range dataToSync.intersectionID {
		secret := secrets.bins[ID]
		localSecret := localSecrets.bins[ID]
		s.syncIntersectingSecrets(
			wg,
			*secret,
			*localSecret,
			func() error { return s.binInstanceService.DeleteLocalSecret(ctx, ID) },
			func() error { return s.binInstanceService.DeleteRemoteSecret(ctx, ID) },
			func() error { return s.binInstanceService.UpdateLocalSecret(ctx, *secret) },
			func() error { return s.binInstanceService.UpdateRemoteSecret(ctx, *localSecret) },
			func() { secretsList.Bins = append(secretsList.Bins, localSecret) },
			func() { secretsList.Bins = append(secretsList.Bins, secret) },
		)
	}

	if err := wg.Wait(); err != nil {
		return dto.SecretsList{}, err
	}

	return secretsList, nil
}

func (s *SecretsService) findAllAndUniqIDs(secrets map[string]bool, localSecrets map[string]bool) DataToSync {
	interIDs := make(map[string]bool, 0)
	var localUniqID, remoteUniqID []string
	for ID := range secrets {
		if _, ok := localSecrets[ID]; !ok {
			remoteUniqID = append(remoteUniqID, ID)
		} else {
			interIDs[ID] = true
		}
	}
	for ID := range localSecrets {
		if _, ok := secrets[ID]; !ok {
			localUniqID = append(localUniqID, ID)
		}
	}

	return DataToSync{
		localUniqID:    localUniqID,
		remoteUniqID:   remoteUniqID,
		intersectionID: interIDs,
	}
}

func (s *SecretsService) syncIntersectingSecrets(
	wg *errgroup.Group,
	secret dto.Secret,
	localSecret dto.Secret,
	deleteLocalSecretFunc func() error,
	deleteRemoteSecretFunc func() error,
	updateLocalSecretFunc func() error,
	updateRemoteSecretFunc func() error,
	appendLocalSecretFunc func(),
	appendRemoteSecretFunc func(),
) {
	if secret.IsDeleted() && !localSecret.IsDeleted() {
		wg.Go(deleteLocalSecretFunc)
		return
	}
	if localSecret.IsDeleted() && !secret.IsDeleted() {
		wg.Go(deleteRemoteSecretFunc)
		return
	}
	if secret.GetUpdatedAt().After(localSecret.GetUpdatedAt()) {
		wg.Go(updateLocalSecretFunc)
		appendRemoteSecretFunc()
	} else if localSecret.GetUpdatedAt().After(secret.GetUpdatedAt()) {
		wg.Go(updateRemoteSecretFunc)
		appendLocalSecretFunc()
	} else {
		appendRemoteSecretFunc()
	}
}
