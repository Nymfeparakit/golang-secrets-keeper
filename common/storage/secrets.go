package storage

import (
	"context"
	"database/sql"
	"github.com/Nymfeparakit/gophkeeper/dto"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type SecretsStorage interface {
	AddCredentials(ctx context.Context, password *dto.LoginPassword) (string, error)
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) (string, error)
	AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) (string, error)
	AddBinaryInfo(ctx context.Context, secret *dto.BinaryInfo) (string, error)
	ListSecrets(ctx context.Context, user string) (dto.SecretsList, error)
	GetCredentialsById(ctx context.Context, id string, user string) (dto.LoginPassword, error)
	GetCardById(ctx context.Context, id string, user string) (dto.CardInfo, error)
	GetTextById(ctx context.Context, id string, user string) (dto.TextInfo, error)
	GetBinaryById(ctx context.Context, id string, user string) (dto.BinaryInfo, error)
	UpdateCredentials(ctx context.Context, pwd *dto.LoginPassword) error
	UpdateTextInfo(ctx context.Context, txt *dto.TextInfo) error
	UpdateBinaryInfo(ctx context.Context, crd *dto.BinaryInfo) error
	UpdateCardInfo(ctx context.Context, crd *dto.CardInfo) error
}

type BaseDBItemsStorage struct {
	db *sqlx.DB
}

func NewBaseItemsStorage(db *sqlx.DB) *BaseDBItemsStorage {
	return &BaseDBItemsStorage{db: db}
}

func (s *BaseDBItemsStorage) ListSecrets(ctx context.Context, user string) (dto.SecretsList, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return dto.SecretsList{}, err
	}
	defer tx.Rollback()

	var secretList dto.SecretsList

	passwords, err := s.listPasswords(ctx, tx, user)
	if err != nil {
		return dto.SecretsList{}, err
	}
	secretList.Passwords = passwords

	texts, err := s.listTexts(ctx, tx, user)
	if err != nil {
		return dto.SecretsList{}, err
	}
	secretList.Texts = texts

	cards, err := s.listCardInfo(ctx, tx, user)
	if err != nil {
		return dto.SecretsList{}, err
	}
	secretList.Cards = cards

	return secretList, nil
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

func (s *BaseDBItemsStorage) GetCredentialsById(ctx context.Context, id string, user string) (dto.LoginPassword, error) {
	query := `SELECT * FROM login_pwd WHERE id=$1 AND user_email=$2`
	var pwd dto.LoginPassword
	err := s.getSecretByID(ctx, id, user, query, &pwd)
	if err != nil {
		return dto.LoginPassword{}, err
	}

	return pwd, nil
}

func (s *BaseDBItemsStorage) GetCardById(ctx context.Context, id string, user string) (dto.CardInfo, error) {
	query := `SELECT * FROM card_info WHERE id=$1 AND user_email=$2`
	var crd dto.CardInfo
	err := s.getSecretByID(ctx, id, user, query, &crd)
	if err != nil {
		return dto.CardInfo{}, err
	}

	return crd, nil
}

func (s *BaseDBItemsStorage) GetTextById(ctx context.Context, id string, user string) (dto.TextInfo, error) {
	query := `SELECT * FROM text_info WHERE id=$1 AND user_email=$2`
	var txt dto.TextInfo
	err := s.getSecretByID(ctx, id, user, query, &txt)
	if err != nil {
		return dto.TextInfo{}, err
	}

	return txt, nil
}

func (s *BaseDBItemsStorage) GetBinaryById(ctx context.Context, id string, user string) (dto.BinaryInfo, error) {
	query := `SELECT * FROM binary_info WHERE id=$1 AND user_email=$2`
	var bin dto.BinaryInfo
	err := s.getSecretByID(ctx, id, user, query, &bin)
	if err != nil {
		return dto.BinaryInfo{}, err
	}

	return bin, nil
}

func (s *BaseDBItemsStorage) getSecretByID(ctx context.Context, id string, user string, query string, dest any) error {
	err := s.db.QueryRowxContext(ctx, query, &id, &user).StructScan(dest)
	if err == sql.ErrNoRows {
		return ErrSecretNotFound
	}
	return err
}

func (s *BaseDBItemsStorage) UpdateCredentials(ctx context.Context, pwd *dto.LoginPassword) error {
	query := `UPDATE login_pwd SET name=:name, metadata=:metadata, login=:login, password=:password, updated_at=:updated_at
WHERE text(id)=:id AND updated_at < :updated_at`
	return s.UpdateSecret(ctx, query, &pwd)
}

func (s *BaseDBItemsStorage) UpdateCardInfo(ctx context.Context, crd *dto.CardInfo) error {
	query := `UPDATE card_info SET name=:name, metadata=:metadata, card_number=:card_number, cvv=:cvv,
	expiration_month=:expiration_month, expiration_year=:expiration_year, updated_at=:updated_at
WHERE text(id)=:id AND updated_at < :updated_at`
	return s.UpdateSecret(ctx, query, &crd)
}

func (s *BaseDBItemsStorage) UpdateTextInfo(ctx context.Context, txt *dto.TextInfo) error {
	query := `UPDATE text_info SET name=:name, metadata=:metadata, text=:text, updated_at=:updated_at
WHERE text(id)=:id AND updated_at < :updated_at`
	return s.UpdateSecret(ctx, query, &txt)
}

func (s *BaseDBItemsStorage) UpdateBinaryInfo(ctx context.Context, crd *dto.BinaryInfo) error {
	query := `UPDATE binary_info SET name=:name, metadata=:metadata, data=:data, updated_at=:updated_at
WHERE text(id)=:id AND updated_at < :updated_at`
	return s.UpdateSecret(ctx, query, &crd)
}

func (s *BaseDBItemsStorage) UpdateSecret(ctx context.Context, query string, arg interface{}) error {
	_, err := s.db.NamedExecContext(ctx, query, arg)
	if err == sql.ErrNoRows {
		return CantUpdateSecret
	}
	return err
}
