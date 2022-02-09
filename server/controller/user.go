package controller

import (
	"github.com/gin-gonic/gin"
)

type userLoginData struct {
	email string `json:"email"`
	phone string `json:"phone"`
	pwd   string `json:"pwd"`
}

func Login(c *gin.Context) {

}
