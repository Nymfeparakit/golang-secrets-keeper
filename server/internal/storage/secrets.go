package storage

import (
	"context"
	commonstorage "github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type SecretsStorage struct {
	db *sqlx.DB
	commonstorage.BaseDBItemsStorage
}

func NewSecretsStorage(db *sqlx.DB) *SecretsStorage {
	baseStorage := commonstorage.NewBaseItemsStorage(db)
	return &SecretsStorage{db: db, BaseDBItemsStorage: *baseStorage}
}

func (s *SecretsStorage) AddCredentials(ctx context.Context, password *dto.LoginPassword) (string, error) {
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

	return "", nil
}

func (s *SecretsStorage) AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error {
	query := `INSERT INTO text_info (name, metadata, user_email, text) VALUES (:name, :metadata, :user_email, :text)`
	_, err := s.db.NamedExecContext(ctx, query, &textInfo)
	if err != nil {
		return err
	}

	return nil
}

func (s *SecretsStorage) AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error {
	query := `INSERT INTO card_info (name, metadata, user_email, card_number, cvv, expiration_month, expiration_year)
VALUES (:name, :metadata, :user_email, :card_number, :cvv, :expiration_month, :expiration_year)`
	_, err := s.db.NamedExecContext(ctx, query, &cardInfo)
	if err != nil {
		return err
	}

	return nil
}

func (s *SecretsStorage) AddBinaryInfo(ctx context.Context, cardInfo *dto.CardInfo) error {
	query := `INSERT INTO binary_info (name, metadata, user_email, data)
VALUES (:name, :metadata, :user_email, :data)`
	_, err := s.db.NamedExecContext(ctx, query, &cardInfo)
	if err != nil {
		return err
	}

	return nil
}
