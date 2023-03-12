package handlers

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SecretsService interface {
	AddCredentials(ctx context.Context, password *dto.LoginPassword) error
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error
	AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error
	ListSecrets(ctx context.Context, user string) (dto.SecretsList, error)
}

type SecretsServer struct {
	secrets.UnimplementedSecretsManagementServer
	authService    AuthService
	secretsService SecretsService
}

func NewSecretsServer(secretsService SecretsService, authService AuthService) *SecretsServer {
	return &SecretsServer{secretsService: secretsService, authService: authService}
}

func (s *SecretsServer) AddCredentials(ctx context.Context, in *secrets.Password) (*secrets.Response, error) {
	var response secrets.Response

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	item := dto.Secret{
		Name:     in.Name,
		User:     user,
		Metadata: in.Metadata,
	}
	password := dto.LoginPassword{
		Login:    in.Login,
		Password: in.Password,
		Secret:   item,
	}

	err := s.secretsService.AddCredentials(ctx, &password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &response, nil
}

func (s *SecretsServer) AddTextInfo(ctx context.Context, in *secrets.TextInfo) (*secrets.Response, error) {
	var response secrets.Response

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	item := dto.Secret{
		Name:     in.Name,
		User:     user,
		Metadata: in.Metadata,
	}
	textInfo := dto.TextInfo{
		Text:   in.Text,
		Secret: item,
	}

	err := s.secretsService.AddTextInfo(ctx, &textInfo)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &response, nil
}

func (s *SecretsServer) AddCardInfo(ctx context.Context, in *secrets.CardInfo) (*secrets.Response, error) {
	var response secrets.Response

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	item := dto.Secret{
		Name:     in.Name,
		User:     user,
		Metadata: in.Metadata,
	}
	cardInfo := dto.CardInfo{
		Secret:          item,
		Number:          in.Number,
		CVV:             in.Cvv,
		ExpirationMonth: in.ExpirationMonth,
		ExpirationYear:  in.ExpirationYear,
	}

	err := s.secretsService.AddCardInfo(ctx, &cardInfo)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &response, nil
}

func (s *SecretsServer) ListSecrets(ctx context.Context, in *secrets.EmptyRequest) (*secrets.ListSecretResponse, error) {
	var response secrets.ListSecretResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	secretsList, err := s.secretsService.ListSecrets(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	for _, pwd := range secretsList.Passwords {
		pwdCopy := secrets.Password{
			Name:     pwd.Name,
			Login:    pwd.Login,
			Password: pwd.Password,
			Metadata: pwd.Metadata,
		}
		response.Passwords = append(response.Passwords, &pwdCopy)
	}
	for _, txt := range secretsList.Texts {
		txtCopy := secrets.TextInfo{
			Name:     txt.Name,
			Text:     txt.Text,
			Metadata: txt.Metadata,
		}
		response.Texts = append(response.Texts, &txtCopy)
	}
	for _, crd := range secretsList.Cards {
		crdCopy := secrets.CardInfo{
			Name:            crd.Name,
			Metadata:        crd.Metadata,
			Number:          crd.Number,
			ExpirationMonth: crd.ExpirationMonth,
			ExpirationYear:  crd.ExpirationYear,
			Cvv:             crd.CVV,
		}
		response.Cards = append(response.Cards, &crdCopy)
	}

	return &response, nil
}
