package commands

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
	c.view.RegisterUserPage()
	return nil
}
