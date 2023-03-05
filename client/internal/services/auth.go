package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"google.golang.org/grpc/metadata"
)

type CredentialsStorage interface {
	SaveToken(token string) error
	GetToken() (string, error)
}

type AuthService struct {
	storageClient      auth.AuthManagementClient
	userToken          string
	credentialsStorage CredentialsStorage
	cryptoService      CryptoService
}

func NewAuthService(client auth.AuthManagementClient, storage CredentialsStorage, cryptoService CryptoService) *AuthService {
	return &AuthService{storageClient: client, credentialsStorage: storage}
}

func (s *AuthService) Register(user *dto.User) error {
	errorMsg := "error occurred on registering user: %s"

	request := auth.SignUpRequest{
		Login:    user.Email,
		Password: user.Password,
	}
	response, err := s.storageClient.SignUp(context.Background(), &request)
	if err != nil {
		return fmt.Errorf(errorMsg, err)
	}
	if response.Error != "" {
		return fmt.Errorf(errorMsg, response.Error)
	}

	return nil
}

// Login saves user credentials (such as access token) in local storage.
func (s *AuthService) Login(email string, pwd string) error {
	errorMsg := "error occurred during user login: %s"

	request := auth.LoginRequest{
		Login:    email,
		Password: pwd,
	}
	response, err := s.storageClient.Login(context.Background(), &request)
	if err != nil {
		return fmt.Errorf(errorMsg, err)
	}
	if response.Error != "" {
		return fmt.Errorf(errorMsg, response.Error)
	}

	token := response.Token
	err = s.credentialsStorage.SaveToken(token)
	if err != nil {
		return fmt.Errorf("failed to save token in local storage: %s", err)
	}

	return nil
}

func (s *AuthService) getUserToken() (string, error) {
	if s.userToken == "" {
		errorMsg := "could not get user token: %s"
		token, err := s.credentialsStorage.GetToken()
		if err != nil {
			return "", fmt.Errorf(errorMsg, err)
		}
		s.userToken = token
	}

	return s.userToken, nil
}

func (s *AuthService) AddAuthMetadata(ctx context.Context) (context.Context, error) {
	token, err := s.getUserToken()
	if err != nil {
		return nil, err
	}
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	return ctx, nil
}
