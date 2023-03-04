package handlers

import (
	"context"
	"errors"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/internal/services"
	"github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService interface {
	GetUserFromContext(ctx context.Context) (string, bool)
	Register(ctx context.Context, user *dto.User) error
	Login(ctx context.Context, email string, pwd string) (string, error)
	AddUserToContext(ctx context.Context, userEmail string) context.Context
	ParseUserToken(tokenString string) (string, error)
}

type AuthServer struct {
	auth.UnimplementedAuthManagementServer
	authService AuthService
}

func NewAuthServer(authService AuthService) *AuthServer {
	return &AuthServer{authService: authService}
}

func (s *AuthServer) SignUp(ctx context.Context, in *auth.SignUpRequest) (*auth.SignUpResponse, error) {
	// todo: add error if user with such email exists
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

func (s *AuthServer) Login(ctx context.Context, in *auth.LoginRequest) (*auth.LoginResponse, error) {
	token, err := s.authService.Login(ctx, in.Login, in.Password)
	if errors.Is(err, services.ErrInvalidCredentials) {
		return nil, status.Error(codes.Unauthenticated, "Invalid login or password")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &auth.LoginResponse{Token: token}, nil
}
