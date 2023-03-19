package commands

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/rs/zerolog/log"
	"os/signal"
)

// AddCommand is a command to add new secret.
type AddCommand struct {
	Type string `short:"t" long:"type" description:"Type of secret that should be added"`
	view SecretsView
}

// NewAddCommand instantiates AddCommand.
func NewAddCommand(view SecretsView) *AddCommand {
	return &AddCommand{view: view}
}

// Execute performs logic to execute add command.
func (c *AddCommand) Execute(args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	notifyCtx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	go func() {
		itemType, err := dto.SecretTypeFromString(c.Type)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
		c.view.AddSecretPage(ctx, itemType)
	}()
	<-notifyCtx.Done()
	stop()
	cancel()
	return nil
}
