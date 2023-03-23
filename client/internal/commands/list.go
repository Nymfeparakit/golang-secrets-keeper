package commands

import (
	"context"
	"os/signal"
)

// ListCommand is a command to list all secrets.
type ListCommand struct {
	view SecretsView
}

// NewListCommand instantiates ListCommand.
func NewListCommand(view SecretsView) *ListCommand {
	return &ListCommand{view: view}
}

// Execute performs logic to execute list command.
func (c *ListCommand) Execute(args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	notifyCtx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	go func() {
		c.view.ListSecretsPage(ctx)
	}()
	<-notifyCtx.Done()
	stop()
	cancel()
	return nil
}
