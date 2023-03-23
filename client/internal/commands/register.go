package commands

import (
	"context"
	"os/signal"
)

// RegisterCommand is a command to register new user.
type RegisterCommand struct {
	view AuthView
}

// NewRegisterCommand instantiates new reg command object.
func NewRegisterCommand(view AuthView) *RegisterCommand {
	return &RegisterCommand{view: view}
}

// Execute performs logic to execute register command.
func (c *RegisterCommand) Execute(args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	notifyCtx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	go func() {
		c.view.RegisterUserPage(ctx)
	}()
	<-notifyCtx.Done()
	stop()
	cancel()
	return nil
}
