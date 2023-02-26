package main

import (
	"github.com/Nymfeparakit/gophkeeper/client/internal/commands"
	"github.com/Nymfeparakit/gophkeeper/client/internal/services"
	"github.com/Nymfeparakit/gophkeeper/client/internal/view"
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
	itemsService := services.NewItemsService(conn)
	itemsView := view.NewItemsView(itemsService)
	authService := services.NewAuthService(conn)
	authView := view.NewAuthView(authService)
	commandParser := commands.NewCommandParser()
	if err := commandParser.InitCommands(itemsView, authView); err != nil {
		log.Fatal().Err(err).Msg("")
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
