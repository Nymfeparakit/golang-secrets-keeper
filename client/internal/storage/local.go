package storage

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"os"
	"path"
	"path/filepath"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func GetLocalStorageConnection() (*sqlx.DB, error) {
	// todo: choose different folders depending on OS
	//if runtime.GOOS == "linux" {
	//	fmt.Println("Hello from Windows")
	//}
	//appDataPath := "/var/lib/" + common.AppName + "/"
	//if _, err := os.Stat(appDataPath); errors.Is(err, os.ErrNotExist) {
	//	err := os.Mkdir(appDataPath, os.ModePerm)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	// todo: creating in /var/lib/ not working, fix
	exPath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	appDataPath := filepath.Dir(exPath)
	dbFilePath := path.Join(appDataPath, common.AppName+".sqlite")
	dbCreated := false
	if _, err := os.Stat(dbFilePath); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(dbFilePath)
		if err != nil {
			return nil, err
		}
		dbCreated = true
	}
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not connect to local storage: %v", err)
	}

	if dbCreated {
		goose.SetBaseFS(embedMigrations)
		if err := goose.SetDialect("sqlite3"); err != nil {
			panic(err)
		}

		if err := goose.Up(db, "migrations"); err != nil {
			panic(err)
		}
	}

	return sqlx.NewDb(db, "sqlite3"), nil
}
