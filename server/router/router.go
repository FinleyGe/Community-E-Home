package router

import (
	"home-server/controller"
	"home-server/utility/middleware"

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
	Router.POST("/api/register", controller.Register)
	Router.POST("/api/test", middleware.Auth, controller.Test)
}
