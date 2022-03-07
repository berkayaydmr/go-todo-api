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
	routerUrl := os.Getenv("ROUTER_URL")
	
	return &Enviroment{
		DatabaseUrl: databaseUrl,
		RouterUrl:   routerUrl,
	}
}