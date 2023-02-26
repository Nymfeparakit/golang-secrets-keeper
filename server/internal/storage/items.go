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
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	itemID, err := s.addItem(ctx, tx, password.Item)
	if err != nil {
		return err
	}
	query := `INSERT INTO login_pwd (login, password, item_id) VALUES ($1, $2, $3)`
	_, err = tx.ExecContext(ctx, query, &password.Login, &password.Password, &itemID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *ItemsStorage) AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	itemID, err := s.addItem(ctx, tx, textInfo.Item)
	if err != nil {
		return err
	}
	query := `INSERT INTO text_info (text, item_id) VALUES ($1, $2)`
	_, err = tx.ExecContext(ctx, query, &textInfo.Text, &itemID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *ItemsStorage) ListItems(ctx context.Context, user string) (dto.ItemsList, error) {
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

	return itemsList, nil
}

func (s *ItemsStorage) addItem(ctx context.Context, tx *sqlx.Tx, item dto.Item) (int, error) {
	query := `INSERT INTO item (name, metadata, user_email) VALUES ($1, $2, $3) RETURNING id`
	var itemID int
	err := tx.QueryRowxContext(ctx, query, &item.Name, &item.Metadata, &item.User).Scan(&itemID)
	if err != nil {
		return 0, err
	}

	return itemID, nil
}

func (s *ItemsStorage) listPasswords(ctx context.Context, tx *sqlx.Tx, user string) ([]*dto.LoginPassword, error) {
	query := `SELECT name, login, password FROM login_pwd INNER JOIN item ON login_pwd.item_id = item.id WHERE user_email=$1`
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
	query := `SELECT name, text FROM text_info INNER JOIN item ON text_info.item_id = item.id WHERE user_email=$1`
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
