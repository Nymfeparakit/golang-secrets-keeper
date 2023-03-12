package storage

import (
	"bytes"
	"context"
	commonstorage "github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type LocalDBSecretsStorage struct {
	db *sqlx.DB
	commonstorage.BaseDBItemsStorage
}

func NewSecretsStorage(db *sqlx.DB) *LocalDBSecretsStorage {
	baseStorage := commonstorage.NewBaseItemsStorage(db)
	return &LocalDBSecretsStorage{db: db, BaseDBItemsStorage: *baseStorage}
}

func (s *LocalDBSecretsStorage) AddCredentials(ctx context.Context, password *dto.LoginPassword) error {
	query := `INSERT INTO login_pwd (id, name, metadata, user_email, login, password)
VALUES (:id, :name, :metadata, :user_email, :login, :password)`
	_, err := s.db.NamedExecContext(ctx, query, &password)
	return err
}

func (s *LocalDBSecretsStorage) AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error {
	query := `INSERT INTO text_info (id, name, metadata, user_email, text) VALUES (:id, :name, :metadata, :user_email, :text)`
	_, err := s.db.NamedExecContext(ctx, query, &textInfo)
	if err != nil {
		return err
	}

	return nil
}

func (s *LocalDBSecretsStorage) AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error {
	query := `INSERT INTO card_info (id, name, metadata, user_email, card_number, cvv, expiration_month, expiration_year)
VALUES (:id, :name, :metadata, :user_email, :card_number, :cvv, :expiration_month, :expiration_year)`
	_, err := s.db.NamedExecContext(ctx, query, &cardInfo)
	if err != nil {
		return err
	}

	return nil
}

func (s *LocalDBSecretsStorage) AddSecrets(ctx context.Context, secretsList dto.SecretsList) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if len(secretsList.Passwords) != 0 {
		queryPwds := `INSERT INTO login_pwd (id, name, metadata, user_email, login, password) VALUES`
		var pwdQueryArgs []interface{}
		for _, pwd := range secretsList.Passwords {
			pwdQueryArgs = append(pwdQueryArgs, &pwd.ID, &pwd.Name, &pwd.Metadata, &pwd.User, &pwd.Login, &pwd.Password)
		}

		queryPwds += s.createBulkInsertArgsString(len(secretsList.Passwords), 6)
		_, err = tx.ExecContext(ctx, queryPwds, pwdQueryArgs...)
		if err != nil {
			return err
		}
	}

	if len(secretsList.Cards) != 0 {
		queryCards := `INSERT INTO card_info (id, name, metadata, user_email, card_number, cvv, expiration_month, expiration_year)`
		var crdQueryArgs []interface{}
		for _, crd := range secretsList.Cards {
			crdQueryArgs = append(crdQueryArgs, &crd.ID, &crd.Name, &crd.Metadata, &crd.User, &crd.Number, &crd.ExpirationMonth, &crd.ExpirationYear)
		}

		queryCards += s.createBulkInsertArgsString(len(secretsList.Cards), 8)
		_, err = tx.ExecContext(ctx, queryCards, crdQueryArgs...)
		if err != nil {
			return err
		}
	}

	if len(secretsList.Texts) != 0 {
		queryTxts := `INSERT INTO text_info (id, name, metadata, user_email, text)`
		var txtQueryArgs []interface{}
		for _, txt := range secretsList.Texts {
			txtQueryArgs = append(txtQueryArgs, &txt.ID, &txt.Name, &txt.Metadata, &txt.User, &txt.Text)
		}

		queryTxts += s.createBulkInsertArgsString(len(secretsList.Texts), 5)
		_, err = tx.ExecContext(ctx, queryTxts, txtQueryArgs...)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *LocalDBSecretsStorage) createBulkInsertArgsString(rowsNum int, columnsNum int) string {
	idx := 1
	var buffer bytes.Buffer
	buffer.WriteString("\n")
	for i := 0; i < rowsNum; i++ {
		buffer.WriteString("(")
		for i := 0; i < columnsNum; i++ {
			buffer.WriteString("$" + strconv.Itoa(idx))
			if i != columnsNum-1 {
				buffer.WriteString(", ")
			}
			idx++
		}
		if i == rowsNum-1 {
			buffer.WriteString(")")
		} else {
			buffer.WriteString("),\n")
		}
	}

	return buffer.String()
}
