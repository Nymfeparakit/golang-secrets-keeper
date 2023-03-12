package main

import (
	"github.com/Nymfeparakit/gophkeeper/client/internal/commands"
	"github.com/Nymfeparakit/gophkeeper/client/internal/config"
	"github.com/Nymfeparakit/gophkeeper/client/internal/services"
	"github.com/Nymfeparakit/gophkeeper/client/internal/storage"
	"github.com/Nymfeparakit/gophkeeper/client/internal/view"
	commonconfig "github.com/Nymfeparakit/gophkeeper/common/config"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	cfg := &config.Config{}
	err := commonconfig.InitConfig(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("can not initialize config")
	}

	conn, err := services.ConnectToServer(cfg.ServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to server")
	}
	defer conn.Close()

	secretsClient := secrets.NewSecretsManagementClient(conn)
	credentialStorage, err := storage.OpenCredentialsStorage()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to open credentials storage")
	}
	authClient := auth.NewAuthManagementClient(conn)

	cryptoService := services.NewCryptoService(credentialStorage)
	authService := services.NewAuthService(authClient, credentialStorage, cryptoService)
	secretsService := services.NewSecretsService(secretsClient, authService, cryptoService)

	authView := view.NewAuthView(authService)
	secretsView := view.NewSecretsView(secretsService)

	commandParser := commands.NewCommandParser()
	if err := commandParser.InitCommands(secretsView, authView); err != nil {
		log.Fatal().Err(err).Msg("can't init commands")
	}
	if _, err := commandParser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}
