package utils

import (
	"go-todo-api/common"

	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId uint64) (string, error) {
	envSecretKey := common.GetEnviroment().SecretKey
	tokenOpt := jwt.MapClaims{
		"authorized":       true,
		"authorizedUserId": userId,
		"expiredTime":      time.Now().Add(time.Minute),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenOpt)
	signedToken, err := token.SignedString([]byte(envSecretKey))
	return signedToken, err
}

func IsTokenValid(tokenString string) (bool,error) {
	envSecretKey := common.GetEnviroment().SecretKey
	token,err := jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {
        return []byte(envSecretKey), nil
    })
	if err != nil {
		return false, err
	}
	
	if token.Valid {
		return true, nil
	}
	
	return false,nil
}
