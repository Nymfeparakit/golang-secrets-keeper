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

type LocalSecretsStorage interface {
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
	localUniqID  []string
	remoteUniqID []string
	allID        map[string]bool
}

type SecretsService struct {
	authService        AuthMetadataService
	storageClient      secrets.SecretsManagementClient
	cryptoService      SecretCryptoService
	localStorage       LocalSecretsStorage
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

	text.User = credentials.Email
	// if remote was available
	if response != nil {
		text.ID = response.Id
	}
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

	card.User = credentials.Email
	if response != nil {
		card.ID = response.Id
	}
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

	bin.User = credentials.Email
	if response != nil {
		bin.ID = response.Id
	}
	err = s.localStorage.AddBinaryInfo(ctx, bin)
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
	request := secrets.EmptyRequest{}
	response, err := s.storageClient.ListSecrets(ctx, &request)
	if err != nil && !checkUnavailableError(err) {
		return dto.SecretsList{}, err
	}

	var secretsList dto.SecretsList
	secretsMap := SecretsMap{
		passwords: make(map[string]*dto.LoginPassword, len(response.Passwords)),
		cards:     make(map[string]*dto.CardInfo, len(response.Cards)),
		texts:     make(map[string]*dto.TextInfo, len(response.Texts)),
		bins:      make(map[string]*dto.BinaryInfo, len(response.Bins)),
	}
	for _, pwd := range response.Passwords {
		pwdDest := common.PasswordFromProto(pwd)
		secretsList.Passwords = append(secretsList.Passwords, &pwdDest)
		secretsMap.passwords[pwdDest.ID] = &pwdDest
	}
	for _, txt := range response.Texts {
		txtDest := common.TextFromProto(txt)
		secretsList.Texts = append(secretsList.Texts, &txtDest)
		secretsMap.texts[txtDest.ID] = &txtDest
	}
	for _, crd := range response.Cards {
		crdDest := common.CardFromProto(crd)
		secretsList.Cards = append(secretsList.Cards, &crdDest)
		secretsMap.cards[crdDest.ID] = &crdDest
	}
	for _, bin := range response.Bins {
		dest := common.BinaryFromProto(bin)
		secretsList.Bins = append(secretsList.Bins, &dest)
		secretsMap.bins[dest.ID] = &dest
	}

	//get secrets from local storage
	//localSecretsList, err := s.localStorage.ListSecrets(ctx, credentials.Email)
	//
	//localSecretsMap := SecretsMap{
	//	passwords: make(map[string]*dto.LoginPassword, len(response.Passwords)),
	//	cards:     make(map[string]*dto.CardInfo, len(response.Cards)),
	//	texts:     make(map[string]*dto.TextInfo, len(response.Texts)),
	//	bins:      make(map[string]*dto.BinaryInfo, len(response.Bins)),
	//}
	//for _, pwd := range localSecretsList.Passwords {
	//	localSecretsMap.passwords[pwd.ID] = pwd
	//}
	//for _, txt := range localSecretsList.Texts {
	//	localSecretsMap.texts[txt.ID] = txt
	//}
	//for _, crd := range localSecretsList.Cards {
	//	localSecretsMap.cards[crd.ID] = crd
	//}
	//for _, bin := range localSecretsList.Bins {
	//	localSecretsMap.bins[bin.ID] = bin
	//}

	return secretsList, nil
}

func (s *SecretsService) LoadSecrets(ctx context.Context) error {
	secretsList, err := s.ListSecrets()
	if err != nil {
		return err
	}

	return s.localStorage.AddSecrets(ctx, secretsList)
}

// todo: change to base secret
//func (s *SecretsService) syncSecrets(secrets SecretsMap, localSecrets SecretsMap) dto.SecretsList {
//	var pwdID, localPwdID map[string]bool
//	for ID := range secrets.passwords {
//		pwdID[ID] = true
//	}
//	for ID := range localSecrets.passwords {
//		localPwdID[ID] = true
//	}
//	dataToSync := s.findAllAndUniqIDs(pwdID, localPwdID)
//
//	var allCards map[string]*dto.CardInfo
//	var localUniqCardID, remoteUniqCardID []string
//	for ID, secret := range secrets.cards {
//		if _, ok := localSecrets.cards[ID]; !ok {
//			remoteUniqCardID = append(remoteUniqCardID, ID)
//		}
//		allCards[ID] = secret
//	}
//	for ID, secret := range localSecrets.cards {
//		if _, ok := secrets.cards[ID]; !ok {
//			localUniqCardID = append(localUniqCardID, ID)
//		}
//		allCards[ID] = secret
//	}
//
//}
//
//func (s *SecretsService) findAllAndUniqIDs(secrets map[string]bool, localSecrets map[string]bool) DataToSync {
//	var allIDs map[string]bool
//	var localUniqID, remoteUniqID []string
//	for ID := range secrets {
//		if _, ok := localSecrets[ID]; !ok {
//			remoteUniqID = append(remoteUniqID, ID)
//		}
//		allIDs[ID] = true
//	}
//	for ID := range localSecrets {
//		if _, ok := secrets[ID]; !ok {
//			localUniqID = append(localUniqID, ID)
//		}
//		allIDs[ID] = true
//	}
//
//	return DataToSync{
//		localUniqID:  localUniqID,
//		remoteUniqID: remoteUniqID,
//		allID:        allIDs,
//	}
//}
