package controller

import (
	"home-server/model"
	"home-server/utility"
	"home-server/utility/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
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
	Type string `json:"type"`
}

func Login(c *gin.Context) {
	// Should use ShouldBindJson instead of PostForm.
	api := LoginAPI{}
	c.ShouldBind(&api)
	// email := c.PostForm("email")
	// method := c.PostForm("method")
	// phone := c.PostForm("phone")
	// pwd := c.PostForm("pwd")
	// fmt.Println(email, method, phone, pwd)
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
	user := model.User{}
	database.DB.
		Where("email = ?", API.Email).
		Or("phone = ?", API.Phone).First(&user)
	if (user != model.User{}) {
		c.JSON(406, gin.H{
			"status": -1, // exist
		})
	} else {
		user.Pwd = API.Pwd
		user.Email = API.Email
		user.Phone = API.Phone
		t, err := strconv.Atoi(API.Type)
		if err != nil {
			c.JSON(400, gin.H{
				"status": -100, // other error
			})
		}
		user.Type = uint8(t)

		e := database.DB.Create(&user)
		if e.Error != nil {
			c.JSON(500, gin.H{
				"status": -100, // other error
			})
		} else {
			c.JSON(200, gin.H{
				"status": 0, //OK
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
