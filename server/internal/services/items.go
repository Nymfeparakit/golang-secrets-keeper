package services

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
)

type ItemsService struct {
	storage storage.ItemsStorage
}

func NewItemsService(storage storage.ItemsStorage) *ItemsService {
	return &ItemsService{storage: storage}
}

func (s *ItemsService) AddPassword(ctx context.Context, password *dto.LoginPassword) (string, error) {
	return s.storage.AddPassword(ctx, password)
}

func (s *ItemsService) AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error {
	return s.storage.AddTextInfo(ctx, textInfo)
}

func (s *ItemsService) AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error {
	return s.storage.AddCardInfo(ctx, cardInfo)
}

func (s *ItemsService) ListItems(ctx context.Context, user string) (dto.ItemsList, error) {
	return s.storage.ListItems(ctx, user)
}
