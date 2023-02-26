package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
)

type ItemsStorage interface {
	AddPassword(ctx context.Context, password *dto.LoginPassword) error
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error
	ListItems(ctx context.Context, user string) (dto.ItemsList, error)
}

type ItemsService struct {
	storage ItemsStorage
}

func NewItemsService(storage ItemsStorage) *ItemsService {
	return &ItemsService{storage: storage}
}

func (s *ItemsService) AddPassword(ctx context.Context, password *dto.LoginPassword) error {
	return s.storage.AddPassword(ctx, password)
}

func (s *ItemsService) AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error {
	return s.storage.AddTextInfo(ctx, textInfo)
}

func (s *ItemsService) ListItems(ctx context.Context, user string) (dto.ItemsList, error) {
	return s.storage.ListItems(ctx, user)
}
