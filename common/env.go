package common

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Enviroment struct {
	DatabaseUrl string
	RouterUrl   string
	RedisUrl    string
	Debug       bool
	SecretKey   string
}

func GetEnviroment() *Enviroment {
	err := godotenv.Load(".env")
	if err != nil {
		zap.S().Error(err)
		return nil
	}

	databaseUrl := os.Getenv("DSN")
	appHost := os.Getenv("APPLICATION_HOST")
	appPort := os.Getenv("APPLICATION_PORT")
	routerUrl := appHost + ":" + appPort
	redisUrl := appHost + ":" + "6379"
	secretKet := os.Getenv("ACCESS_KEY")
	var debug bool
	if os.Getenv("DEBUG") == "true" {
		debug = true
	}

	return &Enviroment{
		DatabaseUrl: databaseUrl,
		RouterUrl:   routerUrl,
		RedisUrl:    redisUrl,
		Debug:       debug,
		SecretKey:   secretKet,
	}
}
