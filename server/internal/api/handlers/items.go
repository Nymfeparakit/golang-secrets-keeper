package handlers

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	items2 "github.com/Nymfeparakit/gophkeeper/server/proto/items"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ItemsService interface {
	AddPassword(ctx context.Context, password *dto.LoginPassword) error
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error
	ListItems(ctx context.Context, user string) (dto.ItemsList, error)
}

type ItemsServer struct {
	items2.UnimplementedItemsManagementServer
	authService  AuthService
	itemsService ItemsService
}

func NewItemsServer(itemsService ItemsService, authService AuthService) *ItemsServer {
	return &ItemsServer{itemsService: itemsService, authService: authService}
}

func (s *ItemsServer) AddPassword(ctx context.Context, in *items2.Password) (*items2.Response, error) {
	var response items2.Response

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	item := dto.Item{
		Name:     in.Name,
		User:     user,
		Metadata: in.Metadata,
	}
	password := dto.LoginPassword{
		Login:    in.Login,
		Password: in.Password,
		Item:     item,
	}

	err := s.itemsService.AddPassword(ctx, &password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &response, nil
}

func (s *ItemsServer) AddTextInfo(ctx context.Context, in *items2.TextInfo) (*items2.Response, error) {
	var response items2.Response

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	item := dto.Item{
		Name:     in.Name,
		User:     user,
		Metadata: in.Metadata,
	}
	textInfo := dto.TextInfo{
		Text: in.Text,
		Item: item,
	}

	err := s.itemsService.AddTextInfo(ctx, &textInfo)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &response, nil
}

func (s *ItemsServer) ListItems(ctx context.Context, in *items2.EmptyRequest) (*items2.ListItemResponse, error) {
	var response items2.ListItemResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	itemsList, err := s.itemsService.ListItems(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	// todo: we have to parse every item to change it's type. What to do with it?
	// maybe don't use dto, but use just structs from proto
	for _, pwd := range itemsList.Passwords {
		pwdCopy := items2.Password{
			Name:     pwd.Name,
			Login:    pwd.Login,
			Password: pwd.Password,
			Metadata: pwd.Metadata,
		}
		response.Passwords = append(response.Passwords, &pwdCopy)
	}
	for _, txt := range itemsList.Texts {
		txtCopy := items2.TextInfo{
			Name:     txt.Name,
			Text:     txt.Text,
			Metadata: txt.Metadata,
		}
		response.Texts = append(response.Texts, &txtCopy)
	}

	return &response, nil
}
