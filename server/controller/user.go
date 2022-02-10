package controller

import (
	"home-server/model"
	"home-server/utility/database"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(json)
	if (json["e-mail"] == "" && json["method"] == "0") || (json["phone"] == "" && json["method"] == "1") {
		c.JSON(400,
			gin.H{
				"status": -100,
			})
		return
	} else {
		user := model.User{}
		if json["method"] == "0" {
			database.DB.Where("email = ?", json["e-mail"]).First(&user)
		} else {
			database.DB.Where("phone = ?", json["phone"]).First(&user)
		}

		if (model.User{} == user) {
			c.JSON(406, gin.H{
				"status": -1, // not found
			})
			return
		} else {
			if user.Pwd == json["pwd"] {
				// TODO OK here
				// Get JWT
			} else {
				c.JSON(406, gin.H{
					"status": -2, // pwd wrong
				})
			}
		}

	}
}
