package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"os"
)

func ConnectToServer(enableHTTPs bool, serverAddr string) (*grpc.ClientConn, error) {
	if enableHTTPs {
		pemServerCA, err := ioutil.ReadFile("cert.pem")
		if err != nil {
			return nil, err
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			return nil, fmt.Errorf("failed to add server certificate")
		}

		cfg := &tls.Config{
			RootCAs: certPool,
		}

		creds := credentials.NewTLS(cfg)

		return grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds))
	}

	return grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func main() {
	cfg := &config.Config{}
	err := commonconfig.InitConfig(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("can not initialize config")
	}

	conn, err := ConnectToServer(cfg.EnableHTTPS, cfg.ServerAddress)
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
	localConn, err := storage.GetLocalStorageConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to local storage")
	}
	localStorage := storage.NewSecretsStorage(localConn)
	usersLocalStorage := storage.NewUsersStorage(localConn)

	cryptoService := services.NewCryptoService(credentialStorage)
	authMetadataService := services.NewMetadataService()
	pwdInstanceService := services.NewUpdateRetrieveDeletePasswordService(authMetadataService, cryptoService, credentialStorage, localStorage, secretsClient)
	crdInstanceService := services.NewUpdateRetrieveDeleteCardService(authMetadataService, cryptoService, credentialStorage, localStorage, secretsClient)
	txtInstanceService := services.NewUpdateRetrieveDeleteTextService(authMetadataService, cryptoService, credentialStorage, localStorage, secretsClient)
	binInstanceService := services.NewUpdateRetrieveDeleteBinaryService(authMetadataService, cryptoService, credentialStorage, localStorage, secretsClient)
	secretsService := services.NewSecretsService(
		secretsClient,
		authMetadataService,
		cryptoService,
		localStorage,
		credentialStorage,
		pwdInstanceService,
	)
	authService := services.NewAuthService(authClient, credentialStorage, cryptoService, usersLocalStorage, secretsService)

	authView := view.NewAuthView(authService)
	secretsView := view.NewSecretsView(
		secretsService,
		pwdInstanceService,
		crdInstanceService,
		txtInstanceService,
		binInstanceService,
	)

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
