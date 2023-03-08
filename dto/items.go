package dto

import "fmt"

type ItemType int

const (
	PASSWORD ItemType = iota
	TEXT
	CARD
	BINARY
	UNKNOWN
)

var ErrUnknownItemType = fmt.Errorf("unknown item type")

func ItemTypeFromString(s string) (ItemType, error) {
	switch s {
	case "password":
		return PASSWORD, nil
	case "text":
		return TEXT, nil
	case "card":
		return CARD, nil
	case "binary":
		return BINARY, nil
	default:
		return UNKNOWN, ErrUnknownItemType
	}
}

type Item struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	User     string `db:"user_email"`
	Metadata string `db:"metadata"`
}

type LoginPassword struct {
	Item
	Login    string `db:"login"`
	Password string `db:"password"`
}

type TextInfo struct {
	Item
	Text string `db:"text"`
}

type CardInfo struct {
	Item
	Number          string `db:"card_number"`
	CVV             string `db:"cvv"`
	ExpirationMonth string `db:"expiration_month" json:"expiration_month"`
	ExpirationYear  string `db:"expiration_year" json:"expiration_year"`
}

type ItemsList struct {
	Passwords []*LoginPassword
	Texts     []*TextInfo
	Cards     []*CardInfo
}
