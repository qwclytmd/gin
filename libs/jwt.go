package libs

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
	"time"
)

type TokenManager struct {
	SecretKey string
}

type LoginClaims struct {
	RId int `json:"rid"`
	jwt.RegisteredClaims
}

func (t *TokenManager) CreateToken(uid int, name string, rid int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims{
		RId: rid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "bcw",
			Subject:   name,
			ID:        cast.ToString(uid),
		},
	})

	return token.SignedString([]byte(t.SecretKey))
}

func (t *TokenManager) ParseToken(tokenString string) (*LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*LoginClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
