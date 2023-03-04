package interceptors

import (
	"context"
	"errors"
	"github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers"
	"github.com/Nymfeparakit/gophkeeper/server/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var doNotUseAuthMethods = []string{"/proto.AuthManagement/SignUp", "/proto.AuthManagement/Login"}

type AuthorizationServerInterceptor struct {
	authService handlers.AuthService
}

// NewAuthorizationServerInterceptor инициализирует AuthorizationServerInterceptor.
func NewAuthorizationServerInterceptor(authService handlers.AuthService) *AuthorizationServerInterceptor {
	return &AuthorizationServerInterceptor{authService: authService}
}

func (i *AuthorizationServerInterceptor) Unary(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if sliceContainsStr(doNotUseAuthMethods, info.FullMethod) {
		return handler(ctx, req)
	}

	var token string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("authorization")
		if len(values) > 0 {
			token = values[0]
		}
	}
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Authentication credentials not provided")
	}

	email, err := i.authService.ParseUserToken(token)
	if errors.Is(err, services.ErrInvalidAccessToken) {
		return nil, status.Error(codes.Unauthenticated, "Provided token is not valid")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	// добавляем пользователя в контекст
	ctx = i.authService.AddUserToContext(ctx, email)

	return handler(ctx, req)
}
