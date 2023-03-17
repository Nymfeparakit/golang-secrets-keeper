package commands

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/jessevdk/go-flags"
)

type Options struct {
}

var options Options

// SecretsView - view with pages for secrets, with page to list secrets, add and update them.
type SecretsView interface {
	ListSecretsPage()
	AddSecretPage(itemType dto.SecretType)
	UpdateSecretPage(itemType dto.SecretType, secretID string)
	DeleteSecretPage(itemType dto.SecretType, secretID string)
}

// AuthView - view with pages to register new users and login existing ones.
type AuthView interface {
	RegisterUserPage()
	LoginUserPage()
}

// CommandParser - parser of command line commands.
type CommandParser struct {
	parser *flags.Parser
}

// NewCommandParser instantiates new command parser object.
func NewCommandParser() *CommandParser {
	parser := flags.NewParser(&options, flags.Default)
	return &CommandParser{parser: parser}
}

// InitCommands instantiates objects for command line commands.
func (p *CommandParser) InitCommands(secretsView SecretsView, authView AuthView) error {
	err := p.initAuthCommands(authView)
	if err != nil {
		return err
	}
	err = p.initSecretsCommands(secretsView)
	return err
}

func (p *CommandParser) initSecretsCommands(view SecretsView) error {
	addCmd := NewAddCommand(view)
	_, err := p.parser.AddCommand("add",
		"Add new item",
		"The add command adds new item.",
		addCmd,
	)
	if err != nil {
		return err
	}
	listCmd := NewListCommand(view)
	_, err = p.parser.AddCommand("list",
		"List secrets",
		"The list command lists all existing secrets.",
		listCmd,
	)
	if err != nil {
		return err
	}
	updateCmd := NewUpdateCommand(view)
	_, err = p.parser.AddCommand("update",
		"Update secret",
		"The update command updates certain item.",
		updateCmd,
	)
	if err != nil {
		return err
	}
	deleteCmd := NewDeleteCommand(view)
	_, err = p.parser.AddCommand("delete",
		"Delete secret",
		"The delete command deletes certain secret.",
		deleteCmd,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *CommandParser) initAuthCommands(view AuthView) error {
	registerCmd := NewRegisterCommand(view)
	_, err := p.parser.AddCommand("reg",
		"Register new user",
		"The add command registers new user.",
		registerCmd,
	)
	if err != nil {
		return err
	}
	loginCmd := NewLoginCommand(view)
	_, err = p.parser.AddCommand(
		"login",
		"Login user",
		"The login command logins user.",
		loginCmd,
	)
	if err != nil {
		return err
	}

	return nil
}

// Parse parses command line arguments.
func (p *CommandParser) Parse() ([]string, error) {
	return p.parser.Parse()
}
