package storage

import (
	"context"
	commonstorage "github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type ItemsStorage interface {
	AddPassword(ctx context.Context, password *dto.LoginPassword) (string, error)
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error
	AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error
	ListItems(ctx context.Context, user string) (dto.ItemsList, error)
}

type ServerDBItemsStorage struct {
	db *sqlx.DB
	commonstorage.BaseDBItemsStorage
}

func NewItemsStorage(db *sqlx.DB) *ServerDBItemsStorage {
	storage := commonstorage.NewBaseItemsStorage(db)
	return &ServerDBItemsStorage{db: db, BaseDBItemsStorage: *storage}
}

func (s *ServerDBItemsStorage) AddPassword(ctx context.Context, password *dto.LoginPassword) (string, error) {
	query := `INSERT INTO login_pwd (name, metadata, user_email, login, password)
VALUES (:name, :metadata, :user_email, :login, :password) RETURNING id`
	stmt, err := s.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return "", err
	}

	var createdID string
	err = stmt.QueryRowxContext(ctx, &password).Scan(&createdID)
	if err != nil {
		return "", err
	}

	return createdID, nil
}

func (s *ServerDBItemsStorage) AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error {
	query := `INSERT INTO text_info (name, metadata, user_email, text) VALUES (:name, :metadata, :user_email, :text)`
	_, err := s.db.NamedExecContext(ctx, query, &textInfo)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServerDBItemsStorage) AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error {
	query := `INSERT INTO card_info (name, metadata, user_email, card_number, cvv, expiration_month, expiration_year)
VALUES (:name, :metadata, :user_email, :card_number, :cvv, :expiration_month, :expiration_year)`
	_, err := s.db.NamedExecContext(ctx, query, &cardInfo)
	if err != nil {
		return err
	}

	return nil
}
