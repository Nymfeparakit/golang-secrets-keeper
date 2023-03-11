package dto

import (
	"fmt"
	"time"
)

type SecretType int

const (
	PASSWORD SecretType = iota
	TEXT
	CARD
	BINARY
	UNKNOWN
)

var ErrUnknownSecretType = fmt.Errorf("unknown secret type")

func SecretTypeFromString(s string) (SecretType, error) {
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
		return UNKNOWN, ErrUnknownSecretType
	}
}

type Secret struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	User      string `db:"user_email"`
	Metadata  string `db:"metadata"`
	UpdatedAt time.Time
}

type LoginPassword struct {
	Secret
	Login    string `db:"login"`
	Password string `db:"password"`
}

type TextInfo struct {
	Secret
	Text string `db:"text"`
}

type CardInfo struct {
	Secret
	Number          string `db:"card_number"`
	CVV             string `db:"cvv"`
	ExpirationMonth string `db:"expiration_month" json:"expiration_month"`
	ExpirationYear  string `db:"expiration_year" json:"expiration_year"`
}

type SecretsList struct {
	Passwords []*LoginPassword
	Texts     []*TextInfo
	Cards     []*CardInfo
}
