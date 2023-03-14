package handlers

import (
	"context"
	"github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SecretsService interface {
	AddCredentials(ctx context.Context, password *dto.LoginPassword) (string, error)
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) error
	AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) error
	ListSecrets(ctx context.Context, user string) (dto.SecretsList, error)
	GetCredentialsById(ctx context.Context, id string, user string) (*dto.LoginPassword, error)
	UpdateCredentials(ctx context.Context, password *dto.LoginPassword) error
}

type SecretsServer struct {
	secrets.UnimplementedSecretsManagementServer
	authService    AuthService
	secretsService SecretsService
}

func NewSecretsServer(secretsService SecretsService, authService AuthService) *SecretsServer {
	return &SecretsServer{secretsService: secretsService, authService: authService}
}

func (s *SecretsServer) AddCredentials(ctx context.Context, in *secrets.Password) (*secrets.AddResponse, error) {
	var response secrets.AddResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	item := dto.BaseSecret{
		Name:     in.Name,
		User:     user,
		Metadata: in.Metadata,
	}
	password := dto.LoginPassword{
		Login:      in.Login,
		Password:   in.Password,
		BaseSecret: item,
	}

	createdID, err := s.secretsService.AddCredentials(ctx, &password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	response.Id = createdID

	return &response, nil
}

func (s *SecretsServer) AddTextInfo(ctx context.Context, in *secrets.TextInfo) (*secrets.AddResponse, error) {
	var response secrets.AddResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	item := dto.BaseSecret{
		Name:     in.Name,
		User:     user,
		Metadata: in.Metadata,
	}
	textInfo := dto.TextInfo{
		Text:       in.Text,
		BaseSecret: item,
	}

	err := s.secretsService.AddTextInfo(ctx, &textInfo)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &response, nil
}

func (s *SecretsServer) AddCardInfo(ctx context.Context, in *secrets.CardInfo) (*secrets.AddResponse, error) {
	var response secrets.AddResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	item := dto.BaseSecret{
		Name:     in.Name,
		User:     user,
		Metadata: in.Metadata,
	}
	cardInfo := dto.CardInfo{
		BaseSecret:      item,
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
			Id:       pwd.ID,
			Name:     pwd.Name,
			Login:    pwd.Login,
			Password: pwd.Password,
			Metadata: pwd.Metadata,
			User:     pwd.User,
		}
		response.Passwords = append(response.Passwords, &pwdCopy)
	}
	for _, txt := range secretsList.Texts {
		txtCopy := secrets.TextInfo{
			Id:       txt.ID,
			Name:     txt.Name,
			Text:     txt.Text,
			Metadata: txt.Metadata,
			User:     txt.User,
		}
		response.Texts = append(response.Texts, &txtCopy)
	}
	for _, crd := range secretsList.Cards {
		crdCopy := secrets.CardInfo{
			Id:              crd.ID,
			Name:            crd.Name,
			Metadata:        crd.Metadata,
			Number:          crd.Number,
			ExpirationMonth: crd.ExpirationMonth,
			ExpirationYear:  crd.ExpirationYear,
			Cvv:             crd.CVV,
			User:            crd.User,
		}
		response.Cards = append(response.Cards, &crdCopy)
	}

	return &response, nil
}

func (s *SecretsServer) GetCredentialsByID(
	ctx context.Context,
	in *secrets.GetSecretRequest,
) (*secrets.GetCredentialsResponse, error) {
	var response secrets.GetCredentialsResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	pwd, err := s.secretsService.GetCredentialsById(ctx, in.Id, user)
	if err == storage.ErrSecretNotFound {
		return nil, status.Error(codes.PermissionDenied, "BaseSecret does not exist")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	response.Password = &secrets.Password{
		Id:       pwd.ID,
		Name:     pwd.Name,
		Login:    pwd.Login,
		Password: pwd.Password,
		Metadata: pwd.Metadata,
		User:     pwd.User,
	}

	return &response, nil
}

func (s *SecretsServer) UpdateCredentials(ctx context.Context, in *secrets.Password) (*secrets.EmptyResponse, error) {
	var response secrets.EmptyResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	item := dto.BaseSecret{
		Name:     in.Name,
		User:     user,
		Metadata: in.Metadata,
	}
	password := dto.LoginPassword{
		Login:      in.Login,
		Password:   in.Password,
		BaseSecret: item,
	}

	err := s.secretsService.UpdateCredentials(ctx, &password)
	if err == storage.CantUpdateSecret {
		return nil, status.Error(codes.PermissionDenied, "Can not update secret")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &response, nil
}
