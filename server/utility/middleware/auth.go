package middleware

import "github.com/gin-gonic/gin"

func Auth(c *gin.Context) {
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {

	}
}
