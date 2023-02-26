package storage

import (
	"context"
	"errors"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

type UsersStorage struct {
	db *sqlx.DB
}

func NewUsersStorage(db *sqlx.DB) *UsersStorage {
	return &UsersStorage{db: db}
}

func (s *UsersStorage) CreateUser(ctx context.Context, user *dto.User) error {
	query := `INSERT INTO auth_user (email, password) VALUES ($1, $2)`
	_, err := s.db.ExecContext(ctx, query, user.Email, user.Password)

	var pgErr pgx.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgerrcode.UniqueViolation {
			return ErrUserAlreadyExists
		}
	}
	if err != nil {
		return err
	}

	return nil
}
