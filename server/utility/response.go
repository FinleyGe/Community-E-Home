package utility

import (
	"github.com/gin-gonic/gin"
)

func response(c *gin.Context, code int, message string, data gin.H) {
	c.JSON(code, gin.H{
		"message": message,
		"data":    data,
	})
}

func ResponseError(c *gin.Context, message string) {
	response(c, 400, message, nil)
}

func ResponseSuccess(c *gin.Context, data gin.H) {
	response(c, 200, "ok", data)
}

func ResponseServerError(c *gin.Context, message string) {
	response(c, 500, message, nil)
}
