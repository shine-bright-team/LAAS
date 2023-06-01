package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenClaims struct {
	UserId   string `json:"userId,omitempty"`
	IsLender bool   `json:"isLender,omitempty"`
	jwt.RegisteredClaims
}

func SignToken(userId string, isLender bool) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		UserId:   userId,
		IsLender: isLender,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "api.laas.thistine.com",
			Subject:   "userCredential",
		},
	})
	key := GetEnv("JWT_SECRETS")
	s, err := t.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return s, nil
}

func ValidateToken(token string) (*TokenClaims, error) {
	t, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tk.Header["alg"])
		}
		key := GetEnv("JWT_SECRETS")
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(*TokenClaims); ok && t.Valid {

		return &TokenClaims{UserId: claims.UserId, IsLender: claims.IsLender}, nil
	} else {
		return nil, fmt.Errorf("there is an error verfying claims")
	}

}
