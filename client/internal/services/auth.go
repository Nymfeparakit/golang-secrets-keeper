package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"google.golang.org/grpc/metadata"
)

type CredentialsStorage interface {
	SaveCredentials(email string, token string) error
	GetToken() (string, error)
}

type UserCryptoService interface {
	CreateUserKey(string) error
}

type UsersStorage interface {
	CreateUser(ctx context.Context, email string) error
}

type SecretsLoadingService interface {
	LoadSecrets(ctx context.Context) error
}

type AuthService struct {
	storageClient      auth.AuthManagementClient
	userToken          string
	credentialsStorage CredentialsStorage
	cryptoService      UserCryptoService
	localUsersStorage  UsersStorage
	secretsService     SecretsLoadingService
}

func NewAuthService(
	client auth.AuthManagementClient,
	storage CredentialsStorage,
	cryptoService UserCryptoService,
	usersStorage UsersStorage,
	itemsService SecretsLoadingService,
) *AuthService {
	return &AuthService{
		storageClient:      client,
		credentialsStorage: storage,
		cryptoService:      cryptoService,
		localUsersStorage:  usersStorage,
		secretsService:     itemsService,
	}
}

// Register registers new user in remote storage.
func (s *AuthService) Register(ctx context.Context, user *dto.User) error {
	errorMsg := "error occurred on registering user: %s"

	request := auth.SignUpRequest{
		Login:    user.Email,
		Password: user.Password,
	}
	response, err := s.storageClient.SignUp(ctx, &request)
	if err != nil {
		return fmt.Errorf(errorMsg, err)
	}
	if response.Error != "" {
		return fmt.Errorf(errorMsg, response.Error)
	}

	return nil
}

// Login saves user credentials (such as access token) in local storage.
func (s *AuthService) Login(ctx context.Context, email string, pwd string) error {
	errorMsg := "error occurred during user login: %s"

	request := auth.LoginRequest{
		Login:    email,
		Password: pwd,
	}
	response, err := s.storageClient.Login(ctx, &request)
	if err != nil {
		return fmt.Errorf(errorMsg, err)
	}
	if response.Error != "" {
		return fmt.Errorf(errorMsg, response.Error)
	}

	token := response.Token
	err = s.credentialsStorage.SaveCredentials(email, token)
	if err != nil {
		return fmt.Errorf("failed to save user credentials in local storage: %s", err)
	}
	err = s.cryptoService.CreateUserKey(pwd)
	if err != nil {
		return fmt.Errorf("failed to save user credentials in local storage: %s", err)
	}

	return s.loadUserData(ctx, email)
}

func (s *AuthService) loadUserData(ctx context.Context, email string) error {
	err := s.localUsersStorage.CreateUser(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to save user data in local storage: %s", err)
	}
	err = s.secretsService.LoadSecrets(ctx)
	if err != nil {
		return fmt.Errorf("failed to load secrets in local storage: %s", err)
	}

	return nil
}

// MetadataService - service for performing operations with grpc metadata.
type MetadataService struct{}

// NewMetadataService creates new MetadataService object.
func NewMetadataService() *MetadataService {
	return &MetadataService{}
}

// AddAuthMetadata adds authentication metadata (authorization token).
func (s *MetadataService) AddAuthMetadata(ctx context.Context, token string) (context.Context, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	return ctx, nil
}
