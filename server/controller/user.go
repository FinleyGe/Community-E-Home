package controller

import (
	"fmt"
	"home-server/model"
	"home-server/utility"
	"home-server/utility/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	email := c.PostForm("email")
	method := c.PostForm("method")
	phone := c.PostForm("phone")
	pwd := c.PostForm("pwd")
	fmt.Println(email, method, phone, pwd)
	if (email == "" && method == "0") || (phone == "" && method == "1") {
		c.JSON(400,
			gin.H{
				"status": -100,
			})
		return
	} else {
		user := model.User{}
		if method == "0" {
			database.DB.Where("email = ?", email).First(&user)
		} else {
			database.DB.Where("phone = ?", phone).First(&user)
		}

		if (model.User{} == user) {
			c.JSON(406, gin.H{
				"status": -1, // not found
			})
			return
		} else {
			if user.Pwd == pwd {
				// TODO OK here
				// Get JWT
				jwtData := utility.JwtData{
					ID: strconv.Itoa(int(user.Id)),
				}
				jwt := utility.GenerateStandardJwt(&jwtData)
				c.JSON(200, gin.H{
					"status": 0,   // OK
					"jwt":    jwt, //
				})
				return
			} else {
				c.JSON(406, gin.H{
					"status": -2, // pwd wrong
				})
			}
		}
	}
}
