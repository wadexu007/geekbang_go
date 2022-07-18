package config

import (
	"github.com/tkanos/gonfig"
)

var FILE_PATH string

var Conf *Configuration

type Configuration struct {
	FILE_PATH   string
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

// func GetConfig() Configuration {
// 	configuration := Configuration{}
// 	fileName := "/app/conf/config.json"
// 	gonfig.GetConf(fileName, &configuration)
// 	return configuration
// }

func init() {
	Conf = &Configuration{}
	fileName := "/app/conf/config.json"
	gonfig.GetConf(fileName, Conf)
}
