package model

type User struct {
	Id        uint `gorm:"primary_key;autoIncrement"`
	Name      string
	Pwd       string
	Email     string
	Phone     string
	AvatarURL string
	UserType  uint8
	// 0 for volunteer
	// 1 for elder

}
