package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, secret string, days int) (string, error) {
	if secret == "" {
		return "", errors.New("missing jwt secret")
	}
	claims := TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * time.Duration(days))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenStr, secret string) (bool, *TokenClaims) {
	if secret == "" {
		return false, nil
	}
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false, nil
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return true, claims
	}
	return false, nil
}
