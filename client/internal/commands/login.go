package commands

import (
	"context"
	"os/signal"
)

// LoginCommand is a command to login existing user.
type LoginCommand struct {
	view AuthView
}

// NewLoginCommand instantiates new login command.
func NewLoginCommand(view AuthView) *LoginCommand {
	return &LoginCommand{view: view}
}

// Execute performs logic to execute login command.
func (c *LoginCommand) Execute(args []string) error {
	ctx, cancel := context.WithCancel(context.Background())
	notifyCtx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	go func() {
		c.view.LoginUserPage(ctx)
	}()
	<-notifyCtx.Done()
	stop()
	cancel()
	return nil
}
