package dto

import (
	"fmt"
	"time"
)

// SecretType - type of secret.
type SecretType int

const (
	PASSWORD SecretType = iota
	TEXT
	CARD
	BINARY
	UNKNOWN
)

var ErrUnknownSecretType = fmt.Errorf("unknown secret type")

// SecretTypeFromString converts strong to SecretType object.
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

// Secret - specifies secret object with ID and UpdatedAt fields.
type Secret interface {
	GetUpdatedAt() time.Time
	GetID() string
	IsDeleted() bool
}

// BaseSecret - base secret object with common fields for all objects.
type BaseSecret struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	User      string    `db:"user_email"`
	Metadata  string    `db:"metadata"`
	UpdatedAt time.Time `db:"updated_at"`
	Deleted   bool      `db:"deleted"`
}

// GetUpdatedAt - returns the time when the object was last updated.
func (s BaseSecret) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// GetID - returns the unique ID of secret.
func (s BaseSecret) GetID() string {
	return s.ID
}

func (s BaseSecret) IsDeleted() bool {
	return s.Deleted
}

// LoginPassword - stores the login and password data that the user uses in a particular service.
type LoginPassword struct {
	BaseSecret
	Login    string `db:"login"`
	Password string `db:"password"`
}

// TextInfo - stores secret text note.
type TextInfo struct {
	BaseSecret
	Text string `db:"text"`
}

// CardInfo - stores credit card information.
type CardInfo struct {
	BaseSecret
	Number          string `db:"card_number"`
	CVV             string `db:"cvv"`
	ExpirationMonth string `db:"expiration_month" json:"expiration_month"`
	ExpirationYear  string `db:"expiration_year" json:"expiration_year"`
}

// BinaryInfo - stores secret binary data(as base64 encoded).
type BinaryInfo struct {
	BaseSecret
	Data string `db:"data"`
}

// SecretsList - list of all secrets.
type SecretsList struct {
	Passwords []*LoginPassword
	Texts     []*TextInfo
	Cards     []*CardInfo
	Bins      []*BinaryInfo
}
