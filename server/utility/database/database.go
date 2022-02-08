package database

import (
	"home-server/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	dbHost := config.Config.GetString("data")
}
