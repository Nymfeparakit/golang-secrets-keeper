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

type Secret interface {
	GetUpdatedAt() time.Time
	SetUpdatedAt(t time.Time)
	GetID() string
}

type BaseSecret struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	User      string    `db:"user_email"`
	Metadata  string    `db:"metadata"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (s BaseSecret) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s BaseSecret) SetUpdatedAt(t time.Time) {
	s.UpdatedAt = t
}

func (s BaseSecret) GetID() string {
	return s.ID
}

type LoginPassword struct {
	BaseSecret
	Login    string `db:"login"`
	Password string `db:"password"`
}

type TextInfo struct {
	BaseSecret
	Text string `db:"text"`
}

type CardInfo struct {
	BaseSecret
	Number          string `db:"card_number"`
	CVV             string `db:"cvv"`
	ExpirationMonth string `db:"expiration_month" json:"expiration_month"`
	ExpirationYear  string `db:"expiration_year" json:"expiration_year"`
}

type BinaryInfo struct {
	BaseSecret
	Data string `db:"data"`
}

type SecretsList struct {
	Passwords []*LoginPassword
	Texts     []*TextInfo
	Cards     []*CardInfo
	Bins      []*BinaryInfo
}
