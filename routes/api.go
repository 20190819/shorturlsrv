package routes

import (
	"shorturlsrv/app/controllers"
	"shorturlsrv/app/middleware"

	"github.com/gin-gonic/gin"
	// https://91porn.com/view_video.php?viewkey=e96a0531fa736f0c3e3d&page=&viewtype=&category=
)

func Api(router *gin.Engine) {
	router.Use(middleware.CroseEnable())
	apiGroup := router.Group("api")
	// 授权相关
	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/register", controllers.AuthController.Register)
	authGroup.POST("/login", controllers.AuthController.Login)
	// 短链密钥相关
	secGroup := apiGroup.Group("/secrets")
	secGroup.Use(middleware.AuthApi())
	secGroup.GET("/", controllers.SecretController.List)
	secGroup.POST("/", controllers.SecretController.Store)
	secGroup.DELETE("/", controllers.SecretController.Destroy)
	// 生成短链相关
	shortUrlGroup := apiGroup.Group("/surls")
	secGroup.Use(middleware.AuthApi())
	shortUrlGroup.GET("", controllers.ShortUrlController.List)
	shortUrlGroup.GET("/:id", controllers.ShortUrlController.Show)
	shortUrlGroup.POST("", controllers.ShortUrlController.Store)
	shortUrlGroup.PUT("/:id", controllers.ShortUrlController.Update)
	shortUrlGroup.DELETE("/:id", controllers.ShortUrlController.Destroy)
}
