package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	mock_services "github.com/Nymfeparakit/gophkeeper/client/internal/services/mocks"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCryptoService_EncryptSecret(t *testing.T) {
	secret := dto.BaseSecret{Name: "some name"}
	pwd := dto.LoginPassword{
		BaseSecret: secret,
		Login:      "login",
		Password:   "pwd",
	}
	pwdToEncrypt := pwd
	userKey := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	keyStorageMock := mock_services.NewMockKeyStorage(ctrl)
	keyStorageMock.EXPECT().GetUserKey().Return(userKey, nil)
	cryptoService := NewCryptoService(keyStorageMock)

	err := cryptoService.EncryptSecret(&pwdToEncrypt)
	require.NoError(t, err)

	decodedData, err := base64.StdEncoding.DecodeString(pwdToEncrypt.Login)
	aesblock, err := aes.NewCipher(userKey[:])
	require.NoError(t, err)
	aesgcm, err := cipher.NewGCM(aesblock)
	require.NoError(t, err)
	nonce := userKey[len(userKey)-aesgcm.NonceSize():]
	decrypted, err := aesgcm.Open(nil, nonce, decodedData, nil)
	require.NoError(t, err)

	assert.Equal(t, string(decrypted), pwd.Login)

	decodedData, err = base64.StdEncoding.DecodeString(pwdToEncrypt.Password)
	decrypted, err = aesgcm.Open(nil, nonce, decodedData, nil)
	require.NoError(t, err)

	assert.Equal(t, string(decrypted), pwd.Password)
}

func TestCryptoService_DecryptSecret(t *testing.T) {
	secret := dto.BaseSecret{Name: "some name"}
	pwd := dto.LoginPassword{
		BaseSecret: secret,
		Login:      "login",
		Password:   "pwd",
	}
	userKey := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	keyStorageMock := mock_services.NewMockKeyStorage(ctrl)
	keyStorageMock.EXPECT().GetUserKey().Return(userKey, nil)
	cryptoService := NewCryptoService(keyStorageMock)

	aesblock, err := aes.NewCipher(userKey[:])
	require.NoError(t, err)
	aesgcm, err := cipher.NewGCM(aesblock)
	require.NoError(t, err)
	nonce := userKey[len(userKey)-aesgcm.NonceSize():]
	encrypted := aesgcm.Seal(nil, nonce, []byte(pwd.Login), nil)
	require.NoError(t, err)
	encodedLogin := base64.StdEncoding.EncodeToString(encrypted)
	encrypted = aesgcm.Seal(nil, nonce, []byte(pwd.Password), nil)
	require.NoError(t, err)
	encodedPwd := base64.StdEncoding.EncodeToString(encrypted)
	encryptedPwd := dto.LoginPassword{Login: encodedLogin, Password: encodedPwd}

	err = cryptoService.DecryptSecret(&encryptedPwd)

	require.NoError(t, err)
	assert.Equal(t, pwd.Login, encryptedPwd.Login)
	assert.Equal(t, pwd.Password, encryptedPwd.Password)
}

func TestCryptoService_CreateUserKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	pwd := "123123"
	key := sha256.Sum256([]byte(pwd))
	userKey := key[:]
	keyStorageMock := mock_services.NewMockKeyStorage(ctrl)
	keyStorageMock.EXPECT().SaveUserKey(userKey).Return(nil)
	cryptoService := NewCryptoService(keyStorageMock)

	err := cryptoService.CreateUserKey(pwd)

	require.NoError(t, err)
}
