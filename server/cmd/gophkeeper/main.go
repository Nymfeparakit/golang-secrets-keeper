package main

import (
	"context"
	commonconfig "github.com/Nymfeparakit/gophkeeper/common/config"
	"github.com/Nymfeparakit/gophkeeper/server/internal/api"
	"github.com/Nymfeparakit/gophkeeper/server/internal/config"
	"github.com/Nymfeparakit/gophkeeper/server/internal/services"
	"github.com/Nymfeparakit/gophkeeper/server/internal/storage"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"os/signal"
	"syscall"
)

func connectToDB(dbURL string) *sqlx.DB {
	db, err := sqlx.Open("pgx", dbURL)
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to database:")
	}
	return db
}

func main() {
	cfg := &config.Config{}
	err := commonconfig.InitConfig(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("can not initialize config")
	}
	db := connectToDB(cfg.DatabaseDSN)
	itemsStorage := storage.NewSecretsStorage(db)
	userStorage := storage.NewUsersStorage(db)
	itemsService := services.NewSecretsService(itemsStorage)
	// todo: get secret key from env
	authService := services.NewAuthService(userStorage, "123")
	server, err := api.NewServer(cfg.EnableHTTPS, authService, itemsService)
	if err != nil {
		log.Fatal().Err(err).Msg("can not initialize server:")
		return
	}
	if err = server.Start(cfg.ServerAddress); err != nil {
		log.Fatal().Err(err).Msg("can not start server:")
		return
	}

	notifyCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-notifyCtx.Done()
	stop()
}
