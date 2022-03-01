package middleware

import (
	"home-server/model"
	"home-server/utility"
	"home-server/utility/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {

		utility.ResponseError(c, "Unauthorized")
		return
	} else {

		id, err := utility.ParseToken(jwtToken)
		if err != nil {
			utility.ResponseError(c, "Jwt Error")
			c.Abort()
			return
		}
		user := model.User{}
		database.DB.Where("id = ?", id).First(&user)
		if !user.Valid {
			utility.ResponseError(c, "Email is invalid")
			c.Abort()
			return
		} else {
			id_int, _ := strconv.Atoi(id)
			c.Set("id", id_int)
			c.Next()
		}
	}
}

func EmailVertify(c *gin.Context) {
	// Email 认证专用auth middleware
	jwtToken := c.GetHeader("Authorization")

	if jwtToken == "" {
		utility.ResponseError(c, "Unauthorized")
		c.Abort()
		return
	} else {
		id, err := utility.ParseToken(jwtToken)
		if err != nil {
			utility.ResponseError(c, "Jwt Error")
			c.Abort()
			return
		}
		id_int, _ := strconv.Atoi(id)
		c.Set("id", id_int)
		c.Next()
	}
}
