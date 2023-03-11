package commands

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
)

type AddCommand struct {
	Type string `short:"t" long:"type" description:"Type of secret that should be added"`
	view SecretsView
}

func NewAddCommand(view SecretsView) *AddCommand {
	return &AddCommand{view: view}
}

func (c *AddCommand) Execute(args []string) error {
	itemType, err := dto.SecretTypeFromString(c.Type)
	if err != nil {
		return err
	}
	c.view.AddSecretPage(itemType)
	return nil
}
