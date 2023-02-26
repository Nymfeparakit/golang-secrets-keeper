package services

import (
	"context"
	"errors"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/internal/storage"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user *dto.User) error
}

type AuthService struct {
	userStorage UserStorage
}

func NewAuthService(userStorage UserStorage) *AuthService {
	return &AuthService{userStorage: userStorage}
}

// UserCtxKey user key in a context
type UserCtxKey string

// AddUserToContext adds user email in context
func (a *AuthService) AddUserToContext(ctx context.Context, userEmail string) context.Context {
	return context.WithValue(ctx, UserCtxKey("userEmail"), userEmail)
}

// GetUserFromContext возвращает id пользователя из контекста, если он в нем присутствует.
func (a *AuthService) GetUserFromContext(ctx context.Context) (string, bool) {
	//userValue := ctx.Value(UserCtxKey("userEmail"))
	//if userValue == nil {
	//	return "", false
	//}
	//userEmail, ok := userValue.(string)
	//if !ok {
	//	return "", false
	//}

	userEmail := "example@example.com"

	return userEmail, true
}

func (a *AuthService) Register(ctx context.Context, user *dto.User) error {
	err := a.userStorage.CreateUser(ctx, user)
	if errors.Is(err, storage.ErrUserAlreadyExists) {
		return ErrUserAlreadyExists
	}
	if err != nil {
		return err
	}

	return nil
}
