package main

import (
	"log"
	"shorturlsrv/app"
	_ "shorturlsrv/bootstrap"
	"shorturlsrv/database"
	"shorturlsrv/database/seed"
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
	seed.Seeder()

	// 注册路由
	router := gin.Default()
	// router.Use(middleware.CroseEnable())
	routes.Api(router)
	log.Fatalln(router.Run(app.EnvStr("APP_HOST")))
}
