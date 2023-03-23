package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
)

// UsersStorage - storage of users in database.
type UsersStorage struct {
	db *sqlx.DB
}

// NewUsersStorage creates new UsersStorage object.
func NewUsersStorage(db *sqlx.DB) *UsersStorage {
	return &UsersStorage{db: db}
}

// CreateUser creates new user in database.
func (s *UsersStorage) CreateUser(ctx context.Context, email string) error {
	query := `INSERT INTO auth_user (email) VALUES ($1)`
	_, err := s.db.ExecContext(ctx, query, email)
	if err != nil {
		return err
	}

	return nil
}
