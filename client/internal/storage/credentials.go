package storage

import (
	"errors"
	"github.com/99designs/keyring"
	"github.com/Nymfeparakit/gophkeeper/dto"
)

const serviceName = "gophkeeper"
const ringEmailKey = "email"
const ringTokenKey = "token"
const ringUserKey = "user_key"

// CredentialsStorage storage for all user credentials in keyring.
type CredentialsStorage struct {
	keyring keyring.Keyring
}

// OpenCredentialsStorage opens keyring for service and creates new CredentialsStorage object.
func OpenCredentialsStorage() (*CredentialsStorage, error) {
	ring, err := keyring.Open(keyring.Config{ServiceName: serviceName})
	if err != nil {
		return nil, err
	}
	return &CredentialsStorage{keyring: ring}, nil
}

// SaveCredentials saves user credentials (email and token) in keyring.
func (s *CredentialsStorage) SaveCredentials(email string, token string) error {
	err := s.saveItem(ringTokenKey, []byte(token))
	if err != nil {
		return err
	}
	err = s.saveItem(ringEmailKey, []byte(email))
	return err
}

// GetCredentials gets user credentials from keyring.
func (s *CredentialsStorage) GetCredentials() (*dto.UserCredentials, error) {
	tokenData, err := s.getItem(ringTokenKey)
	if err != nil {
		return nil, err
	}
	emailData, err := s.getItem(ringEmailKey)
	if err != nil {
		return nil, err
	}

	return &dto.UserCredentials{Email: string(emailData), Token: string(tokenData)}, nil
}

// GetToken gets current user token from keyring.
func (s *CredentialsStorage) GetToken() (string, error) {
	ring, err := s.getKeyring()
	if err != nil {
		return "", err
	}

	item, err := ring.Get(ringTokenKey)
	if errors.Is(err, keyring.ErrKeyNotFound) {
		return "", ErrSecretNotFound
	}
	if err != nil {
		return "", err
	}
	return string(item.Data), nil
}

// SaveUserKey saves user key in keyring.
func (s *CredentialsStorage) SaveUserKey(key []byte) error {
	return s.saveItem(ringUserKey, key)
}

// GetUserKey gets user key from keyring.
func (s *CredentialsStorage) GetUserKey() ([]byte, error) {
	return s.getItem(ringUserKey)
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

func (s *CredentialsStorage) getItem(key string) ([]byte, error) {
	ring, err := s.getKeyring()
	if err != nil {
		return []byte{}, err
	}

	item, err := ring.Get(key)
	if errors.Is(err, keyring.ErrKeyNotFound) {
		return []byte{}, ErrSecretNotFound
	}
	if err != nil {
		return []byte{}, err
	}
	return item.Data, nil
}
