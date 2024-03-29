package handlers

import (
	"context"
	"errors"
	"github.com/Nymfeparakit/gophkeeper/common"
	"github.com/Nymfeparakit/gophkeeper/common/storage"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SecretsService - service for performing CRUD operations with secrets.
type SecretsService interface {
	AddCredentials(ctx context.Context, password *dto.LoginPassword) (string, error)
	AddTextInfo(ctx context.Context, textInfo *dto.TextInfo) (string, error)
	AddCardInfo(ctx context.Context, cardInfo *dto.CardInfo) (string, error)
	AddBinaryInfo(ctx context.Context, binInfo *dto.BinaryInfo) (string, error)

	ListSecrets(ctx context.Context, user string) (dto.SecretsList, error)
	GetCredentialsById(ctx context.Context, id string, user string) (*dto.LoginPassword, error)
	GetCardById(ctx context.Context, id string, user string) (*dto.CardInfo, error)
	GetTextById(ctx context.Context, id string, user string) (*dto.TextInfo, error)
	GetBinaryById(ctx context.Context, id string, user string) (*dto.BinaryInfo, error)

	UpdateCredentials(ctx context.Context, password *dto.LoginPassword) error
	UpdateTextInfo(ctx context.Context, secret *dto.TextInfo) error
	UpdateBinaryInfo(ctx context.Context, secret *dto.BinaryInfo) error
	UpdateCardInfo(ctx context.Context, secret *dto.CardInfo) error

	DeleteCredentials(ctx context.Context, id string) error
	DeleteTextInfo(ctx context.Context, id string) error
	DeleteCardInfo(ctx context.Context, id string) error
	DeleteBinaryInfo(ctx context.Context, id string) error
}

// SecretsServer - server with methods to perform CRUD operations with secrets.
type SecretsServer struct {
	secrets.UnimplementedSecretsManagementServer
	authService    AuthService
	secretsService SecretsService
}

// NewSecretsServer creates new SecretsServer object.
func NewSecretsServer(secretsService SecretsService, authService AuthService) *SecretsServer {
	return &SecretsServer{secretsService: secretsService, authService: authService}
}

// AddCredentials creates new credentials secret.
func (s *SecretsServer) AddCredentials(ctx context.Context, in *secrets.Password) (*secrets.AddResponse, error) {
	var response secrets.AddResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	password := common.PasswordFromProto(in)
	password.User = user
	createdID, err := s.secretsService.AddCredentials(ctx, &password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	response.Id = createdID

	return &response, nil
}

// AddTextInfo creates new text secret.
func (s *SecretsServer) AddTextInfo(ctx context.Context, in *secrets.TextInfo) (*secrets.AddResponse, error) {
	var response secrets.AddResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	textInfo := common.TextFromProto(in)
	textInfo.User = user

	createdID, err := s.secretsService.AddTextInfo(ctx, &textInfo)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	response.Id = createdID

	return &response, nil
}

// AddCardInfo creates new card secret.
func (s *SecretsServer) AddCardInfo(ctx context.Context, in *secrets.CardInfo) (*secrets.AddResponse, error) {
	var response secrets.AddResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	cardInfo := common.CardFromProto(in)
	cardInfo.User = user

	createdID, err := s.secretsService.AddCardInfo(ctx, &cardInfo)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	response.Id = createdID

	return &response, nil
}

// AddBinaryInfo creates new binary secret.
func (s *SecretsServer) AddBinaryInfo(ctx context.Context, in *secrets.BinaryInfo) (*secrets.AddResponse, error) {
	var response secrets.AddResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	bin := common.BinaryFromProto(in)
	bin.User = user

	createdID, err := s.secretsService.AddBinaryInfo(ctx, &bin)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	response.Id = createdID

	return &response, nil
}

// ListSecrets lists all user's secrets.
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
		pwdCopy := common.CredentialsToProto(pwd)
		response.Passwords = append(response.Passwords, pwdCopy)
	}
	for _, txt := range secretsList.Texts {
		txtCopy := common.TextToProto(txt)
		response.Texts = append(response.Texts, txtCopy)
	}
	for _, crd := range secretsList.Cards {
		crdCopy := common.CardToProto(crd)
		response.Cards = append(response.Cards, crdCopy)
	}
	for _, bin := range secretsList.Bins {
		binCopy := common.BinaryToProto(bin)
		response.Bins = append(response.Bins, binCopy)
	}

	return &response, nil
}

// GetCredentialsByID returns credentials secret with the specified id.
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
	if err = s.processGetSecretError(err); err != nil {
		return nil, err
	}

	response.Password = common.CredentialsToProto(pwd)

	return &response, nil
}

// GetCardByID returns card secret with the specified id.
func (s *SecretsServer) GetCardByID(
	ctx context.Context,
	in *secrets.GetSecretRequest,
) (*secrets.GetCardResponse, error) {
	var response secrets.GetCardResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	secret, err := s.secretsService.GetCardById(ctx, in.Id, user)
	if err = s.processGetSecretError(err); err != nil {
		return nil, err
	}

	response.Card = common.CardToProto(secret)

	return &response, nil
}

// GetTextByID returns text secret with the specified id.
func (s *SecretsServer) GetTextByID(
	ctx context.Context,
	in *secrets.GetSecretRequest,
) (*secrets.GetTextResponse, error) {
	var response secrets.GetTextResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	secret, err := s.secretsService.GetTextById(ctx, in.Id, user)
	if err = s.processGetSecretError(err); err != nil {
		return nil, err
	}

	response.Text = common.TextToProto(secret)

	return &response, nil
}

// GetBinaryByID returns binary secret with the specified id.
func (s *SecretsServer) GetBinaryByID(
	ctx context.Context,
	in *secrets.GetSecretRequest,
) (*secrets.GetBinaryResponse, error) {
	var response secrets.GetBinaryResponse

	user, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	secret, err := s.secretsService.GetBinaryById(ctx, in.Id, user)
	if err = s.processGetSecretError(err); err != nil {
		return nil, err
	}

	response.Bin = common.BinaryToProto(secret)

	return &response, nil
}

// UpdateCredentials updates specified by id credentials secret.
func (s *SecretsServer) UpdateCredentials(ctx context.Context, in *secrets.Password) (*secrets.EmptyResponse, error) {
	var response secrets.EmptyResponse

	_, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	password := common.PasswordFromProto(in)

	err := s.secretsService.UpdateCredentials(ctx, &password)
	if err = s.processUpdateSecretError(err); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateCardInfo updates specified by id card secret.
func (s *SecretsServer) UpdateCardInfo(ctx context.Context, in *secrets.CardInfo) (*secrets.EmptyResponse, error) {
	var response secrets.EmptyResponse

	_, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	secret := common.CardFromProto(in)
	err := s.secretsService.UpdateCardInfo(ctx, &secret)
	if err = s.processUpdateSecretError(err); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateTextInfo updates specified by id text secret.
func (s *SecretsServer) UpdateTextInfo(ctx context.Context, in *secrets.TextInfo) (*secrets.EmptyResponse, error) {
	var response secrets.EmptyResponse

	_, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	secret := common.TextFromProto(in)
	err := s.secretsService.UpdateTextInfo(ctx, &secret)
	if err = s.processUpdateSecretError(err); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateBinaryInfo updates specified by id binary secret.
func (s *SecretsServer) UpdateBinaryInfo(ctx context.Context, in *secrets.BinaryInfo) (*secrets.EmptyResponse, error) {
	var response secrets.EmptyResponse

	_, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	secret := common.BinaryFromProto(in)
	err := s.secretsService.UpdateBinaryInfo(ctx, &secret)
	if err = s.processUpdateSecretError(err); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteCredentials deletes specified by id credentials secret.
func (s *SecretsServer) DeleteCredentials(ctx context.Context, in *secrets.DeleteSecretRequest) (*secrets.ResponseWithError, error) {
	var response secrets.ResponseWithError

	_, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	err := s.secretsService.DeleteCredentials(ctx, in.Id)
	if err = s.processDeleteSecretError(err); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteCardInfo deletes specified by id card secret.
func (s *SecretsServer) DeleteCardInfo(ctx context.Context, in *secrets.DeleteSecretRequest) (*secrets.ResponseWithError, error) {
	var response secrets.ResponseWithError

	_, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	err := s.secretsService.DeleteCardInfo(ctx, in.Id)
	if err = s.processDeleteSecretError(err); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteBinaryInfo deletes specified by id binary secret.
func (s *SecretsServer) DeleteBinaryInfo(ctx context.Context, in *secrets.DeleteSecretRequest) (*secrets.ResponseWithError, error) {
	var response secrets.ResponseWithError

	_, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	err := s.secretsService.DeleteBinaryInfo(ctx, in.Id)
	if err = s.processDeleteSecretError(err); err != nil {
		return nil, err
	}

	return &response, nil
}

// DeleteTextInfo deletes specified by id text secret.
func (s *SecretsServer) DeleteTextInfo(ctx context.Context, in *secrets.DeleteSecretRequest) (*secrets.ResponseWithError, error) {
	var response secrets.ResponseWithError

	_, ok := s.authService.GetUserFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "User is not authenticated")
	}

	err := s.secretsService.DeleteTextInfo(ctx, in.Id)
	if err = s.processDeleteSecretError(err); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *SecretsServer) processDeleteSecretError(err error) error {
	if err != nil && errors.Is(err, storage.ErrSecretDoesNotExistOrWasDeleted) {
		return status.Error(codes.FailedPrecondition, "Can't delete secret")
	}

	return err
}

func (s *SecretsServer) processGetSecretError(err error) error {
	if err == storage.ErrSecretNotFound {
		return status.Error(codes.PermissionDenied, "BaseSecret does not exist")
	}
	if err != nil {
		return status.Error(codes.Internal, "Internal server error")
	}

	return nil
}

func (s *SecretsServer) processUpdateSecretError(err error) error {
	if err == storage.CantUpdateSecret {
		return status.Error(codes.PermissionDenied, "Can not update secret")
	}
	if err != nil {
		return status.Error(codes.Internal, "Internal server error")
	}

	return nil
}
