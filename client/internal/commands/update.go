package commands

import "github.com/Nymfeparakit/gophkeeper/dto"

type UpdateCommand struct {
	Type string `short:"t" long:"type" description:"Type of item that should be updated"`
	view SecretsView
}

func NewUpdateCommand(view SecretsView) *UpdateCommand {
	return &UpdateCommand{view: view}
}

func (c *UpdateCommand) Execute(args []string) error {
	itemType, err := dto.SecretTypeFromString(c.Type)
	if err != nil {
		return err
	}
	c.view.AddSecretPage(itemType)
	return nil
}
