package commands

import (
	"fmt"
)

type ListCommand struct {
	view ItemsView
}

func NewListCommand(view ItemsView) *ListCommand {
	return &ListCommand{view: view}
}

func (c *ListCommand) Execute(args []string) error {
	fmt.Println("listing items")
	c.view.ListItemsPage()
	return nil
}
