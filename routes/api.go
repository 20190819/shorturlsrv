package routes

import (
	"shorturlsrv/app/controllers"
	"shorturlsrv/app/middleware"

	"github.com/gin-gonic/gin"
	// https://91porn.com/view_video.php?viewkey=e96a0531fa736f0c3e3d&page=&viewtype=&category=
)

func Api(router *gin.Engine) {
	apiGroup := router.Group("api")

	shortUrlGroup := apiGroup.Group("/short_urls")
	shortUrlGroup.Use(middleware.CheckSecret())
	shortUrlGroup.GET("", controllers.ShortUrl.List)
	shortUrlGroup.GET("/:id", controllers.ShortUrl.Show)
	shortUrlGroup.POST("/", controllers.ShortUrl.Store)
	shortUrlGroup.PUT("/:id", controllers.ShortUrl.Update)
	shortUrlGroup.DELETE("/:id", controllers.ShortUrl.Destroy)
}
