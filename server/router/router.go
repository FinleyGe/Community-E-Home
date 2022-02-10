package router

import (
	"home-server/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func AddCors() {
	config := cors.DefaultConfig()
	config.AllowHeaders = append(
		config.AllowHeaders,
		"Authorization",
	)
	config.AllowAllOrigins = true
	Router.Use(cors.New(config))
}
func RouterInit() {
	AddCors()
	Router.POST("/api/login", controller.Login)
}
