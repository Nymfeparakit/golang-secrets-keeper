package commands

import (
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/jessevdk/go-flags"
)

type Options struct {
}

var options Options

type ItemsView interface {
	ListItemsPage()
	AddItemPage(itemType dto.ItemType)
}

type AuthView interface {
	RegisterUserPage()
	LoginUserPage()
}

type CommandParser struct {
	parser *flags.Parser
}

func NewCommandParser() *CommandParser {
	parser := flags.NewParser(&options, flags.Default)
	return &CommandParser{parser: parser}
}

func (p *CommandParser) InitCommands(itemsView ItemsView, authView AuthView) error {
	err := p.initAuthCommands(authView)
	if err != nil {
		return err
	}
	err = p.initItemsCommands(itemsView)
	return err
}

func (p *CommandParser) initItemsCommands(view ItemsView) error {
	addCmd := NewAddCommand(view)
	_, err := p.parser.AddCommand("add",
		"Add a file",
		"The add command adds new item.",
		addCmd,
	)
	if err != nil {
		return err
	}
	listCmd := NewListCommand(view)
	_, err = p.parser.AddCommand("list",
		"List items",
		"The list command lists all existing items.",
		listCmd,
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

func (p *CommandParser) Parse() ([]string, error) {
	return p.parser.Parse()
}
