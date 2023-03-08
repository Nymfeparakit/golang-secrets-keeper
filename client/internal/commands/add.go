package commands

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
)

type AddCommand struct {
	Type string `short:"t" long:"type" description:"Type of item that should be added"`
	view ItemsView
}

func NewAddCommand(view ItemsView) *AddCommand {
	return &AddCommand{view: view}
}

func (c *AddCommand) Execute(args []string) error {
	itemType, err := dto.ItemTypeFromString(c.Type)
	if err != nil {
		return err
	}
	c.view.AddItemPage(itemType)
	return nil
}
