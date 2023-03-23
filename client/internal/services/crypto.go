package services

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"reflect"
)

// KeyStorage - local storage of user keys.
type KeyStorage interface {
	SaveUserKey([]byte) error
	GetUserKey() ([]byte, error)
}

// CryptoService - service for encrypting and decrypting secrets.
type CryptoService struct {
	userKey    []byte
	keyStorage KeyStorage
}

// NewCryptoService creates new CryptoService object.
func NewCryptoService(keyStorage KeyStorage) *CryptoService {
	return &CryptoService{keyStorage: keyStorage}
}

// CreateUserKey creates key for encryption from user password, then saves it in user keys storage.
func (s *CryptoService) CreateUserKey(password string) error {
	key := sha256.Sum256([]byte(password))
	s.userKey = key[:]
	err := s.keyStorage.SaveUserKey(s.userKey)
	if err != nil {
		return fmt.Errorf("can not save user key: %v", err)
	}

	return nil
}

// EncryptSecret encrypts provided secret.
func (s *CryptoService) EncryptSecret(source any) error {
	// todo: check that pointers were passed
	pSourceValue := reflect.ValueOf(source)
	sourceValue := pSourceValue.Elem()
	sourceObjType := sourceValue.Type()

	for i := 0; i < sourceValue.NumField(); i++ {
		field := sourceObjType.Field(i)
		fieldValue := sourceValue.Field(i)

		if !field.IsExported() {
			continue
		}

		if field.Name == "BaseSecret" {
			fieldType := fieldValue.Type()
			for ii := 0; ii < fieldValue.NumField(); ii++ {
				ffield := fieldType.Field(i)
				ffieldValue := fieldValue.Field(i)
				var encryptedValue string
				if ffield.Name == "Metadata" {
					var err error
					encryptedValue, err = s.encryptStructField(ffieldValue.Interface(), ffield.Name)
					if err != nil {
						return err
					}
					ffieldValue.SetString(encryptedValue)
				}
			}
			continue
		}

		encryptedValue, err := s.encryptStructField(fieldValue.Interface(), field.Name)
		if err != nil {
			return err
		}

		fieldValue.SetString(encryptedValue)
	}

	return nil
}

func (s *CryptoService) encryptStructField(sourceValue any, fieldName string) (string, error) {
	// TODO: sourceValue should be only string
	var bytesValue []byte
	if sourceValueStr, ok := sourceValue.(string); ok {
		bytesValue = []byte(sourceValueStr)
	} else {
		var buff bytes.Buffer
		err := binary.Write(&buff, binary.LittleEndian, sourceValue)
		if err != nil {
			return "", fmt.Errorf("can not convert field %s: %v", fieldName, err)
		}
		bytesValue = buff.Bytes()
	}

	encryptErrMsg := "can't encrypt data: %v"
	aesgcm, err := s.createAesgcm()
	if err != nil {
		return "", fmt.Errorf(encryptErrMsg, err)
	}

	nonce, err := s.createNonce(aesgcm)
	if err != nil {
		return "", fmt.Errorf(encryptErrMsg, err)
	}

	encrypted := aesgcm.Seal(nil, nonce, bytesValue, nil)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (s *CryptoService) getUserKey() ([]byte, error) {
	if s.userKey == nil {
		key, err := s.keyStorage.GetUserKey()
		if err != nil {
			return nil, err
		}
		s.userKey = key
	}

	return s.userKey, nil
}

// DecryptSecret decrypts provided secret.
func (s *CryptoService) DecryptSecret(source any) error {
	// TODO: check that pointers were passed
	pSourceValue := reflect.ValueOf(source)
	sourceValue := pSourceValue.Elem()
	sourceObjType := sourceValue.Type()

	for i := 0; i < sourceValue.NumField(); i++ {
		field := sourceObjType.Field(i)
		fieldValue := sourceValue.Field(i)

		if !field.IsExported() {
			continue
		}

		if field.Name == "BaseSecret" {
			fieldType := fieldValue.Type()
			for ii := 0; ii < fieldValue.NumField(); ii++ {
				ffield := fieldType.Field(i)
				ffieldValue := fieldValue.Field(i)
				var decryptedValue string
				if ffield.Name == "Metadata" {
					var err error
					decryptedValue, err = s.decryptStructField(ffieldValue.Interface().(string))
					if err != nil {
						return err
					}
					ffieldValue.SetString(decryptedValue)
				}
			}
			continue
		}

		decryptedValue, err := s.decryptStructField(fieldValue.Interface().(string))
		if err != nil {
			return err
		}

		fieldValue.SetString(decryptedValue)
	}

	return nil
}

func (s *CryptoService) decryptStructField(sourceValue string) (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(sourceValue)
	if err != nil {
		return "", fmt.Errorf("unable to decode data: %v", err)
	}

	decryptErrMsg := "unable to decrypt data: %v"
	aesgcm, err := s.createAesgcm()
	if err != nil {
		return "", fmt.Errorf(decryptErrMsg, err)
	}

	nonce, err := s.createNonce(aesgcm)
	if err != nil {
		return "", fmt.Errorf(decryptErrMsg, err)
	}

	decrypted, err := aesgcm.Open(nil, nonce, decodedData, nil)
	if err != nil {
		return "", fmt.Errorf(decryptErrMsg, err)
	}

	return string(decrypted), nil
}

func (s *CryptoService) createAesgcm() (cipher.AEAD, error) {
	key, err := s.getUserKey()
	if err != nil {
		return nil, err
	}
	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		return nil, err
	}

	return aesgcm, nil
}

func (s *CryptoService) createNonce(aesgcm cipher.AEAD) ([]byte, error) {
	key, err := s.getUserKey()
	if err != nil {
		return nil, err
	}
	nonce := key[len(key)-aesgcm.NonceSize():]
	return nonce, nil
}
