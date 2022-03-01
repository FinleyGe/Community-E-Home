package router

import (
	"home-server/controller"
	"home-server/utility/middleware"
	"net/http"

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
	AddCors() // 跨域
	// 设置
	Router.StaticFS("../upload", http.Dir("upload"))
	// without auth
	Router.POST("/api/login", controller.Login)
	Router.POST("/api/register", controller.Register)
	// with auth
	Router.GET("/api/user/info", middleware.Auth, controller.UserInfo)
	Router.POST("/api/user/edit", middleware.Auth, controller.EditUserInfo)
	Router.POST("/api/upload/avatar", middleware.Auth, controller.UploadAvatar)
	Router.POST("/api/task/new", middleware.Auth, middleware.IsElder, controller.NewTask)
	// TODO : 获取任务列表
	Router.GET("/api/task/list", middleware.Auth, controller.GetTaskList)
	// TODO : 通过特定条件检索任务
	// TODO : 接受任务
	Router.POST("/api/user/vertify", controller.SendEmail)
	// Router.POST("/api/user/vertify", middleware.EmailVertify, controller.VertifyEmail)
	// Router.GET("/api/test", controller.SendEmail)
}
