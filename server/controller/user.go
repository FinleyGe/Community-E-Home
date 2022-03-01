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
	EmailPhone string `json:"emailPhone"`
	Pwd        string `json:"pwd"`
	Method     int    `json:"method"`
}
type RegisterAPI struct {
	UserAPI
	Type int `json:"type"`
}

func Login(c *gin.Context) {
	api := LoginAPI{}
	c.ShouldBind(&api)
	fmt.Println(api)
	user := model.User{}
	if api.Method == 0 {
		database.DB.Where("email = ?", api.EmailPhone).First(&user)
	} else {
		database.DB.Where("phone = ?", api.EmailPhone).First(&user)
	}

	if (model.User{} == user) {
		utility.ResponseError(c, "User not found")
		return
	} else {
		if user.Pwd == api.Pwd {
			jwtData := utility.JwtData{
				ID: strconv.Itoa(int(user.Id)),
			}
			jwt := utility.GenerateStandardJwt(&jwtData)
			utility.ResponseSuccess(c, gin.H{
				"jwt": jwt,
			})
			return
		} else {
			utility.ResponseError(c, "Wrong Pwd")
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
		utility.ResponseError(c, "User exist")
	} else {
		user.Name = API.Name
		user.Pwd = API.Pwd
		user.Email = API.Email
		user.Phone = API.Phone
		user.Type = uint8(API.Type)
		user.Valid = false

		e := database.DB.Create(&user)
		if e.Error != nil {
			utility.ResponseError(c, "Other Error")
		} else {
			utility.ResponseSuccess(c, nil)
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
		utility.ResponseSuccess(c, gin.H{
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
	// c.JSON(200, gin.H{
	// 	"message": "ok",
	// })
	utility.ResponseSuccess(c, nil)
}

func UploadAvatar(c *gin.Context) {
	file, _ := c.FormFile("avatar")
	id, _ := c.Get("id")
	if err := c.SaveUploadedFile(file, "../upload"+id.(string)); err != nil {
		log.Fatalln("Save uploaded file error:", err)
		utility.ResponseServerError(c, "Server Error")
		return
	}
	database.DB.Model(model.User{}).Where("id = ?", id).Update("avatarUrl", "/upload/"+id.(string))
	utility.ResponseSuccess(c, gin.H{
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
	// c.JSON(200, gin.H{
	// 	"message": "Email Sent",
	// })
	utility.ResponseSuccess(c, nil)
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
