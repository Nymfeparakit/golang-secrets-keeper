package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type UsersStorage struct {
	db *sqlx.DB
}

func NewUsersStorage(db *sqlx.DB) *UsersStorage {
	return &UsersStorage{db: db}
}

func (s *UsersStorage) CreateUser(ctx context.Context, email string) error {
	query := `INSERT INTO auth_user (email) VALUES ($1)`
	_, err := s.db.ExecContext(ctx, query, email)
	if err != nil {
		return err
	}

	return nil
}
