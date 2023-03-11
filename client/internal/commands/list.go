package commands

type ListCommand struct {
	view SecretsView
}

func NewListCommand(view SecretsView) *ListCommand {
	return &ListCommand{view: view}
}

func (c *ListCommand) Execute(args []string) error {
	c.view.ListSecretsPage()
	return nil
}
