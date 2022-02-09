package main

import (
	"home-server/config"
	"home-server/router"
	"home-server/utility/database"
)

func main() {
	// fmt.Println("Hello world")
	config.ConfigInit()
	database.DBInit()
	router.RouterInit()
}
