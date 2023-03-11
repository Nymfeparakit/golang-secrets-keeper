package main

import (
	"github.com/Nymfeparakit/gophkeeper/client/internal/commands"
	"github.com/Nymfeparakit/gophkeeper/client/internal/services"
	"github.com/Nymfeparakit/gophkeeper/client/internal/storage"
	"github.com/Nymfeparakit/gophkeeper/client/internal/view"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog/log"
	"os"
)

var serverAddress = ":8080"

func main() {
	conn, err := services.ConnectToServer(serverAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to server")
	}
	defer conn.Close()

	secretsClient := secrets.NewSecretsManagementClient(conn)
	credentialStorage := storage.NewCredentialsStorage()
	authClient := auth.NewAuthManagementClient(conn)
	localConn, err := storage.GetLocalStorageConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to local storage")
	}
	localStorage := storage.NewItemsStorage(localConn)
	usersLocalStorage := storage.NewUsersStorage(localConn)

	cryptoService := services.NewCryptoService(credentialStorage)
	authMetadataService := services.NewMetadataService()
	secretsService := services.NewSecretsService(secretsClient, authMetadataService, cryptoService, localStorage, credentialStorage)
	authService := services.NewAuthService(authClient, credentialStorage, cryptoService, usersLocalStorage, secretsService)

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
