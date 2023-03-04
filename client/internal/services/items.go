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

type ItemsService struct {
	authService   AuthMetadataService
	storageClient items.ItemsManagementClient
}

func NewItemsService(client items.ItemsManagementClient, service AuthMetadataService) *ItemsService {
	return &ItemsService{storageClient: client, authService: service}
}

func (s *ItemsService) AddPassword(loginPwd *dto.LoginPassword) error {
	ctx, err := s.authService.AddAuthMetadata(context.Background())
	if err != nil {
		return err
	}
	// todo: context should be passed from argument
	request := items.Password{
		Name:     loginPwd.Name,
		Login:    loginPwd.Login,
		Password: loginPwd.Password,
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
	// todo: context should be passed from argument
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
	// todo: context should be passed from argument
	request := items.CardInfo{
		Name:            card.Name,
		Number:          card.CardNumber,
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
		item := dto.Item{
			Name:     pwd.Name,
			Metadata: pwd.Metadata,
		}
		pwdCopy := &dto.LoginPassword{
			Login:    pwd.Login,
			Password: pwd.Password,
			Item:     item,
		}
		itemsList.Passwords = append(itemsList.Passwords, pwdCopy)
	}
	for _, txt := range itemsList.Texts {
		item := dto.Item{
			Name:     txt.Name,
			Metadata: txt.Metadata,
		}
		txtCopy := &dto.TextInfo{
			Text: txt.Text,
			Item: item,
		}
		itemsList.Texts = append(itemsList.Texts, txtCopy)
	}

	return itemsList, nil
}
