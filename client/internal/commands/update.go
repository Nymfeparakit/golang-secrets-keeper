package commands

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/jessevdk/go-flags"
	"os/signal"
)

// UpdateCommand is a command to update existing secret.
type UpdateCommand struct {
	Type string `short:"t" long:"type" description:"Type of item that should be updated"`
	view SecretsView
}

// NewUpdateCommand instantiates new update command object.
func NewUpdateCommand(view SecretsView) *UpdateCommand {
	return &UpdateCommand{view: view}
}

// Execute performs login to execute update command.
func (c *UpdateCommand) Execute(args []string) error {
	if len(args) == 0 {
		errMsg := "secret ID should be specified"
		return &flags.Error{Message: errMsg}
	}
	if len(args) > 1 {
		errMsg := "only one secret ID can be specified"
		return &flags.Error{Message: errMsg}
	}

	secretID := args[0]

	itemType, err := dto.SecretTypeFromString(c.Type)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	notifyCtx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	go func() {
		c.view.UpdateSecretPage(ctx, itemType, secretID)
	}()
	<-notifyCtx.Done()
	stop()
	cancel()
	return nil
}
