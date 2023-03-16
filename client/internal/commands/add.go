package commands

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
)

// AddCommand is a command to add new secret.
type AddCommand struct {
	Type string `short:"t" long:"type" description:"Type of secret that should be added"`
	view SecretsView
}

// NewAddCommand instantiates AddCommand.
func NewAddCommand(view SecretsView) *AddCommand {
	return &AddCommand{view: view}
}

// Execute performs logic to execute add command.
func (c *AddCommand) Execute(args []string) error {
	itemType, err := dto.SecretTypeFromString(c.Type)
	if err != nil {
		return err
	}
	c.view.AddSecretPage(itemType)
	return nil
}
