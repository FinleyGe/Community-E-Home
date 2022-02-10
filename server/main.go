package main

import (
	"home-server/config"
	"home-server/router"
	"home-server/utility/database"
	"home-server/utility/server"
)

func main() {
	config.ConfigInit()
	database.DBInit()
	router.RouterInit()
	server.RunServer()
}
