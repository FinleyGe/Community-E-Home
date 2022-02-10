package server

import (
	"home-server/config"
	"home-server/router"
)

func RunServer() {
	router.Router.Run(":" + config.Config.Server.Port)
}
