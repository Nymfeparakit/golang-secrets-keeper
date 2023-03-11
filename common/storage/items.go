package storage

import (
	"context"
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

type BaseDBItemsStorage struct {
	db *sqlx.DB
}

func NewBaseItemsStorage(db *sqlx.DB) *BaseDBItemsStorage {
	return &BaseDBItemsStorage{db: db}
}

func (s *BaseDBItemsStorage) ListItems(ctx context.Context, user string) (dto.ItemsList, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return dto.ItemsList{}, err
	}
	defer tx.Rollback()

	var itemsList dto.ItemsList

	passwords, err := s.listPasswords(ctx, tx, user)
	if err != nil {
		return dto.ItemsList{}, err
	}
	itemsList.Passwords = passwords

	texts, err := s.listTexts(ctx, tx, user)
	if err != nil {
		return dto.ItemsList{}, err
	}
	itemsList.Texts = texts

	cards, err := s.listCardInfo(ctx, tx, user)
	if err != nil {
		return dto.ItemsList{}, err
	}
	itemsList.Cards = cards

	return itemsList, nil
}

func (s *BaseDBItemsStorage) listPasswords(ctx context.Context, tx *sqlx.Tx, user string) ([]*dto.LoginPassword, error) {
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

func (s *BaseDBItemsStorage) listTexts(ctx context.Context, tx *sqlx.Tx, user string) ([]*dto.TextInfo, error) {
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

func (s *BaseDBItemsStorage) listCardInfo(ctx context.Context, tx *sqlx.Tx, user string) ([]*dto.CardInfo, error) {
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