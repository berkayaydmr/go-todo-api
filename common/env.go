package common

import (
	"os"

	"github.com/joho/godotenv"
)

type Enviroment struct {
	DatabaseUrl string
	RouterUrl   string
}

func GetEnviroment() *Enviroment {
	err := godotenv.Load("database.env")
	if err != nil {
		return nil
	}

	databaseUrl := os.Getenv("DSN")
	appHost := os.Getenv("APPLICATION_HOST")
	appPort := os.Getenv("APPLICATION_PORT")
	routerUrl := appHost + appPort
	
	return &Enviroment{
		DatabaseUrl: databaseUrl,
		RouterUrl:   routerUrl,
	}
}