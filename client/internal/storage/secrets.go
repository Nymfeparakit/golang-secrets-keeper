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
		queryCards := `INSERT INTO card_info (id, name, metadata, user_email, card_number, cvv, expiration_month, expiration_year) VALUES`
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
		queryTxts := `INSERT INTO text_info (id, name, metadata, user_email, text) VALUES`
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

	if len(secretsList.Bins) != 0 {
		query := `INSERT INTO binary_info (id, name, metadata, user_email, data) VALUES`
		var binQueryArgs []interface{}
		for _, secret := range secretsList.Bins {
			binQueryArgs = append(binQueryArgs, &secret.ID, &secret.Name, &secret.Metadata, &secret.User, &secret.Data)
		}

		query += s.createBulkInsertArgsString(len(secretsList.Texts), 5)
		_, err = tx.ExecContext(ctx, query, binQueryArgs...)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *LocalDBSecretsStorage) UpdateCredentials(ctx context.Context, pwd *dto.LoginPassword) error {
	query := `UPDATE login_pwd SET name=:name, metadata=:metadata, login=:login, password=:password
WHERE id=:id AND updated_at < :updated_at`
	return s.UpdateSecret(ctx, query, &pwd)
}

func (s *LocalDBSecretsStorage) UpdateCardInfo(ctx context.Context, crd *dto.CardInfo) error {
	query := `UPDATE card_info SET name=:name, metadata=:metadata, card_number=:card_number, cvv=:cvv,
	expiration_month=:expiration_month, expiration_year=:expiration_year, updated_at=:updated_at
WHERE id=:id AND updated_at < :updated_at`
	return s.UpdateSecret(ctx, query, &crd)
}

func (s *LocalDBSecretsStorage) UpdateTextInfo(ctx context.Context, txt *dto.TextInfo) error {
	query := `UPDATE text_info SET name=:name, metadata=:metadata, text=:text, updated_at=:updated_at
WHERE id=:id AND updated_at < :updated_at`
	return s.UpdateSecret(ctx, query, &txt)
}

func (s *LocalDBSecretsStorage) UpdateBinaryInfo(ctx context.Context, crd *dto.BinaryInfo) error {
	query := `UPDATE binary_info SET name=:name, metadata=:metadata, data=:data, updated_at=:updated_at
WHERE id=:id AND updated_at < :updated_at`
	return s.UpdateSecret(ctx, query, &crd)
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
