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
	// 2 for Admin
	Profile string
	Gender  uint8 `json:"gender"`
	// 0 for male
	// 1 for female
	// 2 for others
	Province string `json:"province"`
	City     string `json:"city"`
}

type UserTask struct {
	// 存储 User 和 Task 之关系
	Model
	Uid    uint
	TaskId uint
}

type UserEmail struct {
	Email       string `json:"email"`
	VertifyCode string `json:"code" gorm:"primary_key"`
}
