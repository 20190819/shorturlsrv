package main

import (
	"log"
	_ "shorturlsrv/bootstrap"
	"shorturlsrv/database"
)

func main() {
	// 初始化 mysql
	err := database.InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	// 初始化 redis
	database.InitRedis()
	// 迁移文件
	database.Migrate()

}
