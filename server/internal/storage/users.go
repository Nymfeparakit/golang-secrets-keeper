package storage

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

// UsersStorage - storage of users in database.
type UsersStorage struct {
	db *sqlx.DB
}

// NewUsersStorage - creates new UsersStorage object.
func NewUsersStorage(db *sqlx.DB) *UsersStorage {
	return &UsersStorage{db: db}
}

// CreateUser creates new user in database.
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

// GetUserByEmail returns the user by email if the one exists.
func (s *UsersStorage) GetUserByEmail(ctx context.Context, email string) (*dto.User, error) {
	query := `SELECT email, password FROM auth_user WHERE email=$1`
	var existingUser dto.User
	err := s.db.QueryRowxContext(ctx, query, email).StructScan(&existingUser)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserDoesNotExist
	}
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}
