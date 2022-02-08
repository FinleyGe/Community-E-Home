package database

import (
	"fmt"
	"home-server/config"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		config.Config.Database.User,
		config.Config.Database.Passport,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Name)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Database Error!", err)
		os.Exit(-1)
	}

}
