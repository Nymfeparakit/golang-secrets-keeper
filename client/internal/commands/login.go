package commands

type LoginCommand struct {
	view AuthView
}

func NewLoginCommand(view AuthView) *LoginCommand {
	return &LoginCommand{view: view}
}

func (c *LoginCommand) Execute(args []string) error {
	c.view.LoginUserPage()
	return nil
}
