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
	Secret  string
}

type email struct {
	Sender   string
	Pwd      string
	SmtpAddr string
	SmtpPort int
}
type ConfigData struct {
	Database database
	Server   server
	Jwt      jwt
	Email    email
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
			Secret:  config.GetString("Jwt.secret"),
		},
		Email: email{
			Sender:   config.GetString("email.sender"),
			Pwd:      config.GetString("email.pwd"),
			SmtpAddr: config.GetString("email.smtpAddr"),
			SmtpPort: config.GetInt("email.smtpPort"),
		},
	}
}
