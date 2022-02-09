package config

import (
	"log"

	"github.com/spf13/viper"
)

type database struct {
	Host     string
	User     string
	Passport string
	Port     string
	Name     string
}
type server struct {
	Port string
}

type jwt struct {
	Expires int
	Issuer  string
}
type ConfigData struct {
	Database database
	Server   server
	Jwt      jwt
}

var Config ConfigData

func ConfigInit() {
	var config = viper.New()
	config.SetConfigType("yaml")
	config.AddConfigPath("./config")
	config.SetConfigName("config")
	config.WatchConfig()

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalln("Config Error", err)
	}
	Config = ConfigData{
		Database: database{
			Host:     config.GetString("database.host"),
			User:     config.GetString("database.user"),
			Passport: config.GetString("database.passport"),
			Port:     config.GetString("database.port"),
			Name:     config.GetString("database.name"),
		},
		Server: server{
			Port: config.GetString("server.port"),
		},
		Jwt: jwt{
			Expires: config.GetInt("Jwt.expires"),
			Issuer:  config.GetString("Jwt.issuer"),
		},
	}
}
