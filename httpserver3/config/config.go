package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	FILE_PATH   string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	configuration := Configuration{}
	fileName := "/app/conf/config.json"
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
