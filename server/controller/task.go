package controller

import (
	"home-server/model"
	"home-server/utility/database"

	"github.com/gin-gonic/gin"
)

func NewTask(c *gin.Context) {
	task := model.Task{}
	id := c.GetInt("id")
	task.IssuerId = uint(id)
	c.ShouldBind(&task)
	// TODO new task
	database.DB.Create(&task)
}

type taskAPI struct {
	models []model.Task
}

func GetTaskList(c *gin.Context) {
	taskList := []model.Task{}
	database.DB.Find(&taskList)
	api := taskAPI{
		models: taskList,
	}
	c.JSON(200, gin.H{
		"tasks": api.models,
	})
}
