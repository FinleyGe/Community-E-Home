package middleware

import (
	"home-server/model"
	"home-server/utility"
	"home-server/utility/database"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		// there is not an authorization header
		c.JSON(401, gin.H{
			"status": "unauthorized",
		})
		c.Abort()
		return
	} else {
		// id, err := strconv.Atoi(utility.ParseToken(jwtToken))
		// if err != nil {
		// 	c.JSON(500, gin.H{
		// 		"status": "Jwt Error",
		// 	})
		// }
		id, err := utility.ParseToken(jwtToken)
		if err != nil {
			c.JSON(400, gin.H{
				"status": "Jwt Error",
			})
			c.Abort()
			return
		}
		user := model.User{}
		database.DB.Where("id = ?", id).First(&user)
		if (user == model.User{}) {
			c.JSON(400, gin.H{
				"status": "Unauthorized",
			})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
