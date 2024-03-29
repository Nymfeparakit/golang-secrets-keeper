package storage

import (
	commonstorage "github.com/Nymfeparakit/gophkeeper/common/storage"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// SecretsStorage - storage of secrets in database.
type SecretsStorage struct {
	db *sqlx.DB
	commonstorage.BaseDBItemsStorage
}

// NewSecretsStorage - creates new SecretsStorage object.
func NewSecretsStorage(db *sqlx.DB) *SecretsStorage {
	baseStorage := commonstorage.NewBaseItemsStorage(db)
	return &SecretsStorage{db: db, BaseDBItemsStorage: *baseStorage}
}
