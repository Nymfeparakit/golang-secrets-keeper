package handlers

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService interface {
	GetUserFromContext(ctx context.Context) (string, bool)
	Register(ctx context.Context, user *dto.User) error
}

type AuthServer struct {
	auth.UnimplementedAuthManagementServer
	authService AuthService
}

func NewAuthServer(authService AuthService) *AuthServer {
	return &AuthServer{authService: authService}
}

func (s *AuthServer) SignUp(ctx context.Context, in *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	user := dto.User{
		Email:    in.Login,
		Password: in.Password,
	}
	err := s.authService.Register(ctx, &user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &auth.SignUpResponse{}, nil
}
