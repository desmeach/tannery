package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	signingKey []byte
}

type UserClaims struct {
	UserId string
}

func NewManager(signingKey []byte) *Manager {
	return &Manager{signingKey: signingKey}
}

func (m Manager) Build(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
	})

	return token.SignedString(m.signingKey)
}

func (m Manager) Parse(tokenString string) (UserClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return m.signingKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return UserClaims{}, err
	}

	if !token.Valid {
		return UserClaims{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return UserClaims{}, errors.New("invalid token")
	}

	return UserClaims{UserId: claims["userId"].(string)}, nil
}
