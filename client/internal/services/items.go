package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/items"
	"google.golang.org/grpc"
)

type ItemsService struct {
	storageClient items.ItemsManagementClient
}

func NewItemsService(conn *grpc.ClientConn) *ItemsService {
	storageClient := items.NewItemsManagementClient(conn)
	return &ItemsService{storageClient: storageClient}
}

func (s *ItemsService) AddPassword(loginPwd *dto.LoginPassword) error {
	// todo: context should be passed from argument
	request := items.Password{
		Name:     loginPwd.Name,
		Login:    loginPwd.Login,
		Password: loginPwd.Password,
	}
	response, err := s.storageClient.AddPassword(context.Background(), &request)
	if err != nil {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding password: %s", response.Error)
	}
	return nil
}

func (s *ItemsService) AddTextInfo(text *dto.TextInfo) error {
	// todo: context should be passed from argument
	request := items.TextInfo{
		Name: text.Name,
		Text: text.Text,
	}
	response, err := s.storageClient.AddTextInfo(context.Background(), &request)
	if err != nil {
		return err
	}
	if response.Error != "" {
		return fmt.Errorf("error occured on adding password: %s", response.Error)
	}
	return nil
}

func (s *ItemsService) ListItems() (*items.ListItemResponse, error) {
	request := items.EmptyRequest{}
	response, err := s.storageClient.ListItems(context.Background(), &request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
