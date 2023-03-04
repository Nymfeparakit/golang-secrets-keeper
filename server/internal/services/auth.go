package services

import (
	"context"
	"errors"
	"github.com/Nymfeparakit/gophkeeper/dto"
	"github.com/Nymfeparakit/gophkeeper/server/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user *dto.User) error
	GetUserByEmail(ctx context.Context, email string) (*dto.User, error)
}

type AuthService struct {
	userStorage  UserStorage
	tokenService *AuthJWTTokenService
}

func NewAuthService(userStorage UserStorage, secretKey string) *AuthService {
	tokenService := NewAuthJWTTokenService(secretKey)
	return &AuthService{userStorage: userStorage, tokenService: tokenService}
}

// UserCtxKey user key in a context
type UserCtxKey string

// AddUserToContext adds user email in context
func (a *AuthService) AddUserToContext(ctx context.Context, userEmail string) context.Context {
	return context.WithValue(ctx, UserCtxKey("userEmail"), userEmail)
}

// GetUserFromContext возвращает id пользователя из контекста, если он в нем присутствует.
func (a *AuthService) GetUserFromContext(ctx context.Context) (string, bool) {
	userValue := ctx.Value(UserCtxKey("userEmail"))
	if userValue == nil {
		return "", false
	}
	userEmail, ok := userValue.(string)
	if !ok {
		return "", false
	}

	return userEmail, true
}

func (a *AuthService) Register(ctx context.Context, user *dto.User) error {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pwdHash)
	err = a.userStorage.CreateUser(ctx, user)
	if errors.Is(err, storage.ErrUserAlreadyExists) {
		return ErrUserAlreadyExists
	}
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthService) Login(ctx context.Context, email string, pwd string) (string, error) {
	// находим пользователя по логину
	existingUser, err := a.userStorage.GetUserByEmail(ctx, email)
	if errors.Is(err, ErrUserDoesNotExist) {
		return "", ErrInvalidCredentials
	}
	if err != nil {
		return "", err
	}

	// проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(pwd))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", ErrInvalidCredentials
	}
	if err != nil {
		return "", err
	}

	// генерируем токены для пользователя
	tokenString, err := a.tokenService.generateAuthToken(existingUser.Email)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *AuthService) ParseUserToken(tokenString string) (string, error) {
	return a.tokenService.parseJWTToken(tokenString)
}
