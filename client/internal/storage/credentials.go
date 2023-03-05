package storage

import (
	"errors"
	"github.com/99designs/keyring"
)

const serviceName = "gophkeeper"
const ringTokenKey = "token"
const ringUserKey = "user_key"

type CredentialsStorage struct {
	keyring keyring.Keyring
}

func NewCredentialsStorage() *CredentialsStorage {
	return &CredentialsStorage{}
}

func (s *CredentialsStorage) SaveToken(token string) error {
	return s.saveItem(ringTokenKey, []byte(token))
}

func (s *CredentialsStorage) GetToken() (string, error) {
	ring, err := s.getKeyring()
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

func (s *CredentialsStorage) SaveUserKey(key []byte) error {
	return s.saveItem(ringUserKey, key)
}

func (s *CredentialsStorage) getKeyring() (keyring.Keyring, error) {
	if s.keyring == nil {
		ring, err := keyring.Open(keyring.Config{ServiceName: serviceName})
		if err != nil {
			return nil, err
		}
		s.keyring = ring
	}
	return s.keyring, nil
}

func (s *CredentialsStorage) saveItem(key string, data []byte) error {
	ring, err := s.getKeyring()
	if err != nil {
		return err
	}

	err = ring.Set(keyring.Item{Key: key, Data: data})
	if err != nil {
		return err
	}

	return nil
}
