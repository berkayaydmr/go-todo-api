package common

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Enviroment struct {
	DatabaseUrl string
	RouterUrl   string
	Debug	bool
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
	routerUrl := appHost + appPort
	
	var debug bool
	if os.Getenv("LOGGER") == "true" {
		debug = true
	}

	return &Enviroment{
		DatabaseUrl: databaseUrl,
		RouterUrl:   routerUrl,
		Debug: debug,
	}
}
