package common

type Enviroment struct {
	DatabaseUrl string
}

func GetEnviroment() *Enviroment {
	return &Enviroment{
		DatabaseUrl: "host=localhost user=postgres password=berkay1707 dbname=todo port=5432",
	}
}