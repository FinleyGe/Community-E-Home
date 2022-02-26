package controller

import (
	"fmt"
	"home-server/model"
	"home-server/utility"
	"home-server/utility/database"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Pwd   string `json:"pwd"`
}

type LoginAPI struct {
	UserAPI
	Method string `json:"method"`
}
type RegisterAPI struct {
	UserAPI
	Type int `json:"type"`
}

func Login(c *gin.Context) {
	api := LoginAPI{}
	c.ShouldBind(&api)
	if (api.Email == "" && api.Method == "0") || (api.Phone == "" && api.Method == "1") {
		c.JSON(400,
			gin.H{
				"status": -100,
			})
		return
	} else {
		user := model.User{}
		if api.Method == "0" {
			database.DB.Where("email = ?", api.Email).First(&user)
		} else {
			database.DB.Where("phone = ?", api.Phone).First(&user)
		}

		if (model.User{} == user) {
			c.JSON(406, gin.H{
				"status": -1, // not found
			})
			return
		} else {
			if user.Pwd == api.Pwd {
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

func Register(c *gin.Context) {
	API := RegisterAPI{}
	c.ShouldBind(&API)
	fmt.Println(API)
	user := model.User{}
	database.DB.
		Where("email = ?", API.Email).
		Or("phone = ?", API.Phone).
		First(&user)

	if (user != model.User{}) {
		c.JSON(406, gin.H{
			// "status": -1, // exist
			"message": "User exist", //
		})
	} else {
		user.Name = API.Name
		user.Pwd = API.Pwd
		user.Email = API.Email
		user.Phone = API.Phone
		user.Type = uint8(API.Type)
		user.Valid = false

		e := database.DB.Create(&user)
		if e.Error != nil {
			c.JSON(500, gin.H{
				// "status": -100, // other error
				"message": "other error",
			})
		} else {
			c.JSON(200, gin.H{
				// "status": 0, //OK
				"message": "OK",
			})
		}
	}
}

type TestAPI struct {
	Jwt string `json:"jwt"`
}

func Test(c *gin.Context) {
	// test := TestAPI{}
	// c.ShouldBind(&test)
	// fmt.Println(test.Jwt)
	// fmt.Println(utility.ParseToken(test.Jwt))
	c.JSON(200, gin.H{
		"msg": "something",
	})
}

func UserInfo(c *gin.Context) {
	id, _ := c.Get("id")
	user := model.User{}
	database.DB.Where("id = ?", id).First(&user)
	if (user == model.User{}) {
		c.JSON(404, gin.H{
			"status": -100,
		})
	} else {
		c.JSON(200, gin.H{
			"status":     0,
			"name":       user.Name,
			"email":      user.Email,
			"phone":      user.Phone,
			"avatar_url": user.AvatarURL,
			"type":       user.Type,
			"profile":    user.Profile,
		})
	}
}

func EditUserInfo(c *gin.Context) {
	user := model.User{}
	c.ShouldBind(&user)

	id, _ := c.Get("id")
	rawUser := model.User{}
	database.DB.Where("id = ?", id).First(&rawUser)
	rawUser.Profile = user.Profile
	rawUser.Email = user.Email
	rawUser.Name = user.Name
	rawUser.Phone = user.Phone
	database.DB.Save(&rawUser)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func UploadAvatar(c *gin.Context) {
	file, _ := c.FormFile("avatar")
	id, _ := c.Get("id")
	if err := c.SaveUploadedFile(file, "../upload"+id.(string)); err != nil {
		log.Fatalln("Save uploaded file error:", err)
		c.JSON(500, gin.H{
			"message": "server error",
		})
		return
	}
	database.DB.Model(model.User{}).Where("id = ?", id).Update("avatarUrl", "/upload/"+id.(string))
	c.JSON(200, gin.H{
		"avatar_url": "/upload/" + id.(string),
	})
}

func SendEmail(c *gin.Context) {
	id := c.GetInt("id")
	userEmail := model.UserEmail{}
	user := model.User{}
	user.Id = uint(id)
	database.DB.Where("id = ?", id).First(&user)
	// ? Maybe using goroutine is better for next line ?
	userEmail.VertifyCode = utility.SendEmail(user.Email)
	userEmail.Uid = uint(id)
	database.DB.Create(&userEmail)
	c.JSON(200, gin.H{
		"message": "Email Sent",
	})
}

func VertifyEmail(c *gin.Context) {
	id := uint(c.GetInt("id"))
	fmt.Println(c.GetInt("id"))
	userEmail := model.UserEmail{}
	c.ShouldBind(&userEmail) // vertify_code
	database.DB.Find(&userEmail)
	if userEmail.Uid == id {
		// ok
		database.DB.Model(model.User{}).Where("id = ?", id).Update("valid", true)
		database.DB.Delete(&userEmail)
		c.JSON(200, gin.H{
			"message": "vertify success",
		})
	} else {
		c.JSON(400, gin.H{
			"message": "vertify error",
		})
	}
}
