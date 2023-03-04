package storage

import (
	"errors"
	"github.com/99designs/keyring"
)

const serviceName = "gophkeeper"
const ringTokenKey = "token"

type TokenStorage struct{}

func NewTokenStorage() *TokenStorage {
	return &TokenStorage{}
}

func (s *TokenStorage) SaveToken(token string) error {
	ring, err := keyring.Open(keyring.Config{ServiceName: serviceName})
	if err != nil {
		return err
	}

	err = ring.Set(keyring.Item{Key: ringTokenKey, Data: []byte(token)})
	if err != nil {
		return err
	}

	return nil
}

func (s *TokenStorage) GetToken() (string, error) {
	ring, err := keyring.Open(keyring.Config{ServiceName: serviceName})
	if err != nil {
		return "", err
	}

	item, err := ring.Get(ringTokenKey)
	if errors.Is(err, keyring.ErrKeyNotFound) {
		return "", ErrTokenNotFound
	}
	if err != nil {
		return "", err
	}
	return string(item.Data), nil
}
