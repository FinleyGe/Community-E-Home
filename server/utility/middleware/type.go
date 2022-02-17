package middleware

import (
	"fmt"
	"home-server/model"
	"home-server/utility/database"

	"github.com/gin-gonic/gin"
)

// type userType struct {
// 	Type uint8 `json:"type"`
// }

func IsElder(c *gin.Context) {
	id := c.GetString("id")
	user := model.User{}
	database.DB.Where("id = ?", id).Select("type").First(&user)
	// user := userType{}
	fmt.Println(user.Type)
	if user.Type == 1 || user.Type == 2 {
		c.Next()
	} else {
		c.JSON(406, gin.H{"message": "Wrong user type"})
		c.Abort()
	}
}
