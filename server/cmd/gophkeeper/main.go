package main

import (
	"context"
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
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("can not init config:")
		return
	}
	db := connectToDB(cfg.DatabaseDSN)
	itemsStorage := storage.NewItemsStorage(db)
	userStorage := storage.NewUsersStorage(db)
	itemsService := services.NewItemsService(itemsStorage)
	authService := services.NewAuthService(userStorage)
	server := api.NewServer()
	server.RegisterHandlers(itemsService, authService)
	if err = server.Start(cfg.ServerAddress); err != nil {
		log.Fatal().Err(err).Msg("can not start server:")
		return
	}

	notifyCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-notifyCtx.Done()
	stop()
}
