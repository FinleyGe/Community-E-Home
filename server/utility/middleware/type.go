package middleware

import (
	"home-server/model"
	"home-server/utility"
	"home-server/utility/database"

	"github.com/gin-gonic/gin"
)

func IsElder(c *gin.Context) {
	id := c.GetString("id")
	user := model.User{}
	database.DB.Where("id = ?", id).Select("type").First(&user)
	if user.Type == 1 || user.Type == 2 {
		c.Next()
	} else {
		utility.ResponseError(c, "Wrong User Type")
		c.Abort()
	}
}

func IsVolunteer(c *gin.Context) {
	id := c.GetString("id")
	user := model.User{}
	database.DB.Where("id = ?", id).Select("type").First(&user)
	if user.Type == 0 || user.Type == 2 {
		c.Next()
	} else {
		utility.ResponseError(c, "Wrong User Type")
		c.Abort()
	}
}
