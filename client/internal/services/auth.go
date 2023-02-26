package services

import (
	"context"
	"fmt"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"google.golang.org/grpc"
)

type AuthService struct {
	storageClient auth.AuthManagementClient
}

func NewAuthService(conn *grpc.ClientConn) *AuthService {
	storageClient := auth.NewAuthManagementClient(conn)
	return &AuthService{storageClient: storageClient}
}

func (s *AuthService) Register(user *dto.User) error {
	request := auth.SignUpRequest{
		Login:    user.Email,
		Password: user.Password,
	}
	response, err := s.storageClient.SignUp(context.Background(), &request)
	errorMsg := "error occurred on registering user: %s"
	if err != nil {
		return fmt.Errorf(errorMsg, err)
	}
	if response.Error != "" {
		return fmt.Errorf(errorMsg, response.Error)
	}
	return nil
}
