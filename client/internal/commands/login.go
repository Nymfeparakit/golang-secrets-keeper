package commands

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
	c.view.LoginUserPage()
	return nil
}
