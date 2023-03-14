package commands

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/jessevdk/go-flags"
)

type UpdateCommand struct {
	Type string `short:"t" long:"type" description:"Type of item that should be updated"`
	view SecretsView
}

func NewUpdateCommand(view SecretsView) *UpdateCommand {
	return &UpdateCommand{view: view}
}

func (c *UpdateCommand) Execute(args []string) error {
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
	c.view.UpdateSecretPage(itemType, secretID)
	return nil
}
