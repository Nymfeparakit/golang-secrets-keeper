package commands

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
	c.view.ListSecretsPage()
	return nil
}
