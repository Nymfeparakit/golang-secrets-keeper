package storage

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type ItemsStorage struct {
	db *sqlx.DB
}

func NewItemsStorage(db *sqlx.DB) *ItemsStorage {
	return &ItemsStorage{db: db}
}

func (s *ItemsStorage) AddPassword(ctx context.Context, password *dto.LoginPassword) error {
	query := `INSERT INTO login_pwd (name, metadata, user_email, login, password)
VALUES (:name, :metadata, :user_email, :login, :password)`
	_, err := s.db.NamedExecContext(ctx, query, &password)
	if err != nil {
		return err
	}

	return nil
}

func (s *ItemsStorage) AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error {
	query := `INSERT INTO text_info (name, metadata, user_email, text) VALUES (:name, :metadata, :user_email, :text)`
	_, err := s.db.NamedExecContext(ctx, query, &textInfo)
	if err != nil {
		return err
	}

	return nil
}

func (s *ItemsStorage) AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error {
	query := `INSERT INTO card_info (name, metadata, user_email, card_number, cvv, expiration_month, expiration_year)
VALUES (:name, :metadata, :user_email, :card_number, :cvv, :expiration_month, :expiration_year)`
	_, err := s.db.NamedExecContext(ctx, query, &cardInfo)
	if err != nil {
		return err
	}

	return nil
}

func (s *ItemsStorage) ListSecrets(ctx context.Context, user string) (dto.SecretsList, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return dto.SecretsList{}, err
	}
	defer tx.Rollback()

	var itemsList dto.SecretsList

	passwords, err := s.listPasswords(ctx, tx, user)
	if err != nil {
		return dto.SecretsList{}, err
	}
	itemsList.Passwords = passwords

	texts, err := s.listTexts(ctx, tx, user)
	if err != nil {
		return dto.SecretsList{}, err
	}
	itemsList.Texts = texts

	cards, err := s.listCardInfo(ctx, tx, user)
	if err != nil {
		return dto.SecretsList{}, err
	}
	itemsList.Cards = cards

	return itemsList, nil
}

func (s *ItemsStorage) listPasswords(ctx context.Context, tx *sqlx.Tx, user string) ([]*dto.LoginPassword, error) {
	query := `SELECT * FROM login_pwd WHERE user_email=$1`
	rows, err := tx.QueryxContext(ctx, query, user)
	if err != nil {
		return []*dto.LoginPassword{}, err
	}
	var passwords []*dto.LoginPassword
	for rows.Next() {
		var pwd dto.LoginPassword
		err = rows.StructScan(&pwd)
		if err != nil {
			return []*dto.LoginPassword{}, err
		}
		passwords = append(passwords, &pwd)
	}

	return passwords, nil
}

func (s *ItemsStorage) listTexts(ctx context.Context, tx *sqlx.Tx, user string) ([]*dto.TextInfo, error) {
	query := `SELECT * FROM text_info WHERE user_email=$1`
	rows, err := tx.QueryxContext(ctx, query, user)
	if err != nil {
		return []*dto.TextInfo{}, err
	}
	var texts []*dto.TextInfo
	for rows.Next() {
		var txt dto.TextInfo
		err = rows.StructScan(&txt)
		if err != nil {
			return []*dto.TextInfo{}, err
		}
		texts = append(texts, &txt)
	}

	return texts, nil
}

func (s *ItemsStorage) listCardInfo(ctx context.Context, tx *sqlx.Tx, user string) ([]*dto.CardInfo, error) {
	query := `SELECT * FROM card_info WHERE user_email=$1`
	rows, err := tx.QueryxContext(ctx, query, user)
	if err != nil {
		return []*dto.CardInfo{}, err
	}
	var cards []*dto.CardInfo
	for rows.Next() {
		var card dto.CardInfo
		err = rows.StructScan(&card)
		if err != nil {
			return []*dto.CardInfo{}, err
		}
		cards = append(cards, &card)
	}

	return cards, nil
}
