package commands

import (
	"fmt"
)

type AddCommand struct {
	Type string `short:"t" long:"type" description:"Type of item that should be added"`
	view ItemsView
}

func NewAddCommand(view ItemsView) *AddCommand {
	return &AddCommand{view: view}
}

func (c *AddCommand) Execute(args []string) error {
	fmt.Println("adding new item")
	switch c.Type {
	case "password":
		c.view.AddPasswordPage()
	case "text":
		c.view.AddTextInfoPage()
	}

	return nil
}
