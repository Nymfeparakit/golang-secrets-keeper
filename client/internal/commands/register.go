package commands

type RegisterCommand struct {
	view AuthView
}

func NewRegisterCommand(view AuthView) *RegisterCommand {
	return &RegisterCommand{view: view}
}

func (c *RegisterCommand) Execute(args []string) error {
	c.view.RegisterUserPage()
	return nil
}
