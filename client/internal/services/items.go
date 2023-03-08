package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/items"
)

type AuthMetadataService interface {
	AddAuthMetadata(ctx context.Context) (context.Context, error)
}

type ItemCryptoService interface {
	EncryptItem(source any) error
	DecryptItem(source any) error
}

type ItemsService struct {
	authService   AuthMetadataService
	storageClient items.ItemsManagementClient
	cryptoService ItemCryptoService
}

func NewItemsService(
	client items.ItemsManagementClient,
	service AuthMetadataService,
	cryptoService ItemCryptoService,
) *ItemsService {
	return &ItemsService{storageClient: client, authService: service, cryptoService: cryptoService}
}

func (s *ItemsService) AddPassword(loginPwd *dto.LoginPassword) error {
	// todo: context should be passed from argument
	ctx, err := s.authService.AddAuthMetadata(context.Background())
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
	if err != nil {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding password: %s", response.Error)
	}
	return nil
}

func (s *ItemsService) AddTextInfo(text *dto.TextInfo) error {
	ctx, err := s.authService.AddAuthMetadata(context.Background())
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
	ctx, err := s.authService.AddAuthMetadata(context.Background())
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
	ctx, err := s.authService.AddAuthMetadata(context.Background())
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
			Name:     pwd.Name,
			Metadata: pwd.Metadata,
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
			Name:     txt.Name,
			Metadata: txt.Metadata,
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
			Name:     crd.Name,
			Metadata: crd.Metadata,
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
