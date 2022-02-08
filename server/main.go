package main

import (
	"home-server/config"
	"home-server/utility/database"
)

func main() {
	// fmt.Println("Hello world")
	config.ConfigInit() // 初始化配置文件
	database.DBInit()   // 初始化数据库

}
