package utils

import (
	"fmt"
	"go-todo-api/common"

	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint64) (string, error) {
	envSecretKey := common.GetEnvironment().SecretKey // TODO: ask to how can we change it
	fmt.Println(envSecretKey)
	tokenOpt := jwt.MapClaims{
		"authorized":       true,
		"authorizedUserId": userID,
		"expiredTime":      time.Minute,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenOpt)
	signedToken, err := token.SignedString([]byte(envSecretKey))
	return signedToken, err
}

func IsTokenValid(tokenString string) (bool, error) { // TODO: Unused
	envSecretKey := common.GetEnvironment().SecretKey
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(envSecretKey), nil
	})

	if token.Valid {
		return true, err
	}

	return false, err
}
