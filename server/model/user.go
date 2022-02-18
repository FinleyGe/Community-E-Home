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
	Valid   bool `json:"valid"`
}

type UserTask struct {
	// 存储 User 和 Task 之关系
	Model
	Uid    uint
	TaskId uint
}

type UserEmail struct {
	Uid         uint   `json:"uid"`
	VertifyCode string `json:"vertify_code" gorm:"primary_key"`
}
