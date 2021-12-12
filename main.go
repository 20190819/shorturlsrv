package main

import (
	_ "shorturlsrv/bootstrap"
	"shorturlsrv/database"
)

func main() {
	// 初始化 mysql
	database.InitDB()
	// 初始化 redis
	database.InitRedis()
	// 迁移文件
	database.Migrate()
}
