package model

type User struct {
	Id        int `gorm:"primary_key"`
	Name      string
	Pwd       string
	Email     string
	AvatarURL string
}
