package commands

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/jessevdk/go-flags"
)

// DeleteCommand is a command to update existing secret.
type DeleteCommand struct {
	Type string `short:"t" long:"type" description:"Type of item that should be deleted"`
	view SecretsView
}

// NewDeleteCommand instantiates new delete command object.
func NewDeleteCommand(view SecretsView) *DeleteCommand {
	return &DeleteCommand{view: view}
}

// Execute performs logic to execute delete command.
func (c *DeleteCommand) Execute(args []string) error {
	if len(args) == 0 {
		errMsg := "secret ID should be specified"
		return &flags.Error{Message: errMsg}
	}
	if len(args) > 1 {
		errMsg := "only one secret ID can be specified"
		return &flags.Error{Message: errMsg}
	}

	secretID := args[0]

	itemType, err := dto.SecretTypeFromString(c.Type)
	if err != nil {
		return err
	}
	c.view.DeleteSecretPage(itemType, secretID)
	return nil
}
