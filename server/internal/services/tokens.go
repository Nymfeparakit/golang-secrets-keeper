package services

import "github.com/golang-jwt/jwt/v5"

// JWTClaims - stores claims in jwt token.
type JWTClaims struct {
	jwt.RegisteredClaims
	Email string
}

// AuthJWTTokenService - service for performing operations with jwt tokens.
type AuthJWTTokenService struct {
	secretKey string
}

// NewAuthJWTTokenService creates new AuthJWTTokenService object.
func NewAuthJWTTokenService(secretKey string) *AuthJWTTokenService {
	return &AuthJWTTokenService{secretKey: secretKey}
}

func (s *AuthJWTTokenService) generateAuthToken(login string) (string, error) {
	claims := JWTClaims{
		Email:            login,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *AuthJWTTokenService) parseJWTToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", ErrInvalidAccessToken
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return "", ErrInvalidAccessToken
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims.Email, nil
	}

	return "", ErrInvalidAccessToken
}
