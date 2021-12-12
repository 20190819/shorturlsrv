package main

import (
	"log"
	"shorturlsrv/app"
	_ "shorturlsrv/bootstrap"
	"shorturlsrv/database"
	"shorturlsrv/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 mysql
	database.InitDB()
	// 初始化 redis
	database.InitRedis()
	// 迁移文件
	database.Migrate()
	database.Seed()

	// 注册路由
	router := gin.Default()
	routes.Api(router)
	log.Fatalln(router.Run(app.EnvStr("APP_HOST")))

}
