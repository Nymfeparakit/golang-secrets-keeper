package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/items"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthMetadataService interface {
	AddAuthMetadata(ctx context.Context, token string) (context.Context, error)
}

type ItemCryptoService interface {
	EncryptItem(source any) error
	DecryptItem(source any) error
}

type UserCredentialsStorage interface {
	GetCredentials() (*dto.UserCredentials, error)
}

type LocalItemsStorage interface {
	AddPassword(ctx context.Context, password *dto.LoginPassword) error
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error
	AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error
	ListItems(ctx context.Context, user string) (dto.ItemsList, error)
	AddItems(ctx context.Context, itemsList dto.ItemsList) error
}

type ItemsService struct {
	authService      AuthMetadataService
	storageClient    items.ItemsManagementClient
	cryptoService    ItemCryptoService
	localStorage     LocalItemsStorage
	userCredsStorage UserCredentialsStorage
}

func NewItemsService(
	client items.ItemsManagementClient,
	service AuthMetadataService,
	cryptoService ItemCryptoService,
	localStorage LocalItemsStorage,
	userCredsStorage UserCredentialsStorage,
) *ItemsService {
	return &ItemsService{
		storageClient:    client,
		authService:      service,
		cryptoService:    cryptoService,
		localStorage:     localStorage,
		userCredsStorage: userCredsStorage,
	}
}

func (s *ItemsService) AddPassword(loginPwd *dto.LoginPassword) error {
	// todo: context should be passed from argument
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}

	ctx, err := s.authService.AddAuthMetadata(context.Background(), credentials.Token)
	if err != nil {
		return err
	}
	err = s.cryptoService.EncryptItem(loginPwd)
	if err != nil {
		return fmt.Errorf("can't encrypt item: %s", err)
	}
	request := items.Password{
		Name:     loginPwd.Name,
		Login:    loginPwd.Login,
		Password: loginPwd.Password,
		Metadata: loginPwd.Metadata,
	}
	response, err := s.storageClient.AddPassword(ctx, &request)
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
	err = s.localStorage.AddPassword(ctx, loginPwd)
	if err != nil {
		return err
	}

	return nil
}

func (s *ItemsService) AddTextInfo(text *dto.TextInfo) error {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}

	ctx, err := s.authService.AddAuthMetadata(context.Background(), credentials.Token)
	if err != nil {
		return err
	}
	err = s.cryptoService.EncryptItem(text)
	if err != nil {
		return fmt.Errorf("can't encrypt item: %s", err)
	}
	request := items.TextInfo{
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

func (s *ItemsService) AddCardInfo(card *dto.CardInfo) error {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return err
	}

	ctx, err := s.authService.AddAuthMetadata(context.Background(), credentials.Token)
	if err != nil {
		return err
	}
	err = s.cryptoService.EncryptItem(card)
	if err != nil {
		return fmt.Errorf("can't encrypt item: %s", err)
	}
	request := items.CardInfo{
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

func (s *ItemsService) ListItems() (dto.ItemsList, error) {
	credentials, err := s.userCredsStorage.GetCredentials()
	if err != nil {
		return dto.ItemsList{}, err
	}

	ctx, err := s.authService.AddAuthMetadata(context.Background(), credentials.Token)
	if err != nil {
		return dto.ItemsList{}, err
	}

	request := items.EmptyRequest{}
	response, err := s.storageClient.ListItems(ctx, &request)
	if err != nil {
		return dto.ItemsList{}, err
	}

	var itemsList dto.ItemsList
	for _, pwd := range response.Passwords {
		itemDest := dto.Item{
			ID:       pwd.Id,
			Name:     pwd.Name,
			Metadata: pwd.Metadata,
			User:     pwd.User,
		}
		pwdDest := dto.LoginPassword{
			Item:     itemDest,
			Login:    pwd.Login,
			Password: pwd.Password,
		}
		err := s.cryptoService.DecryptItem(&pwdDest)
		if err != nil {
			return dto.ItemsList{}, err
		}
		itemsList.Passwords = append(itemsList.Passwords, &pwdDest)
	}
	for _, txt := range response.Texts {
		itemDest := dto.Item{
			ID:       txt.Id,
			Name:     txt.Name,
			Metadata: txt.Metadata,
			User:     txt.User,
		}
		txtDest := &dto.TextInfo{
			Text: txt.Text,
			Item: itemDest,
		}
		err := s.cryptoService.DecryptItem(txtDest)
		if err != nil {
			return dto.ItemsList{}, err
		}
		itemsList.Texts = append(itemsList.Texts, txtDest)
	}
	for _, crd := range response.Cards {
		itemDest := dto.Item{
			ID:       crd.Id,
			Name:     crd.Name,
			Metadata: crd.Metadata,
			User:     crd.User,
		}
		crdDest := &dto.CardInfo{
			Item:            itemDest,
			Number:          crd.Number,
			CVV:             crd.Cvv,
			ExpirationMonth: crd.ExpirationMonth,
			ExpirationYear:  crd.ExpirationYear,
		}
		err := s.cryptoService.DecryptItem(crdDest)
		if err != nil {
			return dto.ItemsList{}, err
		}
		itemsList.Cards = append(itemsList.Cards, crdDest)
	}

	return itemsList, nil
}

func (s *ItemsService) LoadItems(ctx context.Context) error {
	itemsList, err := s.ListItems()
	if err != nil {
		return err
	}

	return s.localStorage.AddItems(ctx, itemsList)
}

func (s *ItemsService) GetPasswordByID(id string) (*dto.LoginPassword, error) {

}
