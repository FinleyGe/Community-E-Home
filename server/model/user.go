package model

type User struct {
	Id        uint `gorm:"primary_key;autoIncrement"`
	Name      string
	Pwd       string
	Email     string `gorm:"index;unique"`
	Phone     string `gorm:"index;unique"`
	AvatarURL string
	Type      uint8
	// 0 for volunteer
	// 1 for elder
	Profile string
}
