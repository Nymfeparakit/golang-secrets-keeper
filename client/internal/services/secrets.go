package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"golang.org/x/sync/errgroup"
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

type LocalSecretsStorage interface {
	AddCredentials(ctx context.Context, password *dto.LoginPassword) (string, error)
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) (string, error)
	AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) (string, error)
	AddBinaryInfo(ctx context.Context, binInfo *dto.BinaryInfo) (string, error)
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
	localUniqID    []string
	remoteUniqID   []string
	intersectionID map[string]bool
}

type SecretsService struct {
	authService        AuthMetadataService
	storageClient      secrets.SecretsManagementClient
	cryptoService      SecretCryptoService
	localStorage       LocalSecretsStorage
	userCredsStorage   UserCredentialsStorage
	pwdInstanceService UpdatePasswordsService
}

func NewSecretsService(
	client secrets.SecretsManagementClient,
	service AuthMetadataService,
	cryptoService SecretCryptoService,
	localStorage LocalSecretsStorage,
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

	text.User = credentials.Email
	// if remote was available
	if response != nil {
		text.ID = response.Id
	}
	_, err = s.localStorage.AddTextInfo(ctx, text)
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

	card.User = credentials.Email
	if response != nil {
		card.ID = response.Id
	}
	_, err = s.localStorage.AddCardInfo(ctx, card)
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
	secretsListRemote, err := s.listSecrets()
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

func (s *SecretsService) LoadSecrets(ctx context.Context) error {
	// load secrets from remote
	secretsList, err := s.listSecrets()
	if err != nil {
		return err
	}

	return s.localStorage.AddSecrets(ctx, secretsList)
}

func (s *SecretsService) listSecrets() (dto.SecretsList, error) {
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
		secretsList.Passwords = append(secretsList.Passwords, localSecrets.passwords[ID])
		wg.Go(func() error {
			request := common.CredentialsToProto(localSecrets.passwords[ID])
			_, err := s.storageClient.AddCredentials(ctx, request)
			return err
		})
	}
	for _, ID := range dataToSync.remoteUniqID {
		secretsList.Passwords = append(secretsList.Passwords, secrets.passwords[ID])
		wg.Go(func() error {
			_, err := s.localStorage.AddCredentials(ctx, secrets.passwords[ID])
			return err
		})
	}
	for ID := range dataToSync.intersectionID {
		secretsList.Passwords = append(secretsList.Passwords, secrets.passwords[ID])
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
		secretsList.Texts = append(secretsList.Texts, localSecrets.texts[ID])
		wg.Go(func() error {
			request := common.TextToProto(localSecrets.texts[ID])
			_, err := s.storageClient.AddTextInfo(ctx, request)
			return err
		})
	}
	for _, ID := range dataToSync.remoteUniqID {
		secretsList.Texts = append(secretsList.Texts, secrets.texts[ID])
		wg.Go(func() error {
			_, err := s.localStorage.AddTextInfo(ctx, secrets.texts[ID])
			return err
		})
	}
	for ID := range dataToSync.intersectionID {
		secretsList.Texts = append(secretsList.Texts, secrets.texts[ID])
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
		secretsList.Cards = append(secretsList.Cards, localSecrets.cards[ID])
		wg.Go(func() error {
			request := common.CardToProto(localSecrets.cards[ID])
			_, err := s.storageClient.AddCardInfo(ctx, request)
			return err
		})
	}
	for _, ID := range dataToSync.remoteUniqID {
		secretsList.Cards = append(secretsList.Cards, secrets.cards[ID])
		wg.Go(func() error {
			_, err := s.localStorage.AddCardInfo(ctx, secrets.cards[ID])
			return err
		})
	}
	for ID := range dataToSync.intersectionID {
		secretsList.Cards = append(secretsList.Cards, secrets.cards[ID])
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
		secretsList.Bins = append(secretsList.Bins, localSecrets.bins[ID])
		wg.Go(func() error {
			request := common.BinaryToProto(localSecrets.bins[ID])
			_, err := s.storageClient.AddBinaryInfo(ctx, request)
			return err
		})
	}
	for _, ID := range dataToSync.remoteUniqID {
		secretsList.Bins = append(secretsList.Bins, secrets.bins[ID])
		wg.Go(func() error {
			_, err := s.localStorage.AddBinaryInfo(ctx, secrets.bins[ID])
			return err
		})
	}
	for ID := range dataToSync.intersectionID {
		secretsList.Bins = append(secretsList.Bins, secrets.bins[ID])
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
