package commands

type ListCommand struct {
	view ItemsView
}

func NewListCommand(view ItemsView) *ListCommand {
	return &ListCommand{view: view}
}

func (c *ListCommand) Execute(args []string) error {
	c.view.ListItemsPage()
	return nil
}
