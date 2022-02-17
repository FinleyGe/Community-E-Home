package database

import (
	"fmt"
	"home-server/config"
	"home-server/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Database.User,
		config.Config.Database.Passport,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Name)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Database Error!", err)
		panic(err)
	}
	// * : Auto Migrate here
	err = DB.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Fatalln("Database Error!", err)
		panic(err)
	}
}
