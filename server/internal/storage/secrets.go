package storage

import (
	commonstorage "github.com/Nymfeparakit/gophkeeper/common/storage"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type SecretsStorage struct {
	db *sqlx.DB
	commonstorage.BaseDBItemsStorage
}

func NewSecretsStorage(db *sqlx.DB) *SecretsStorage {
	baseStorage := commonstorage.NewBaseItemsStorage(db)
	return &SecretsStorage{db: db, BaseDBItemsStorage: *baseStorage}
}
