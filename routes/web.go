package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Web(router *gin.Engine) {
	// 加载静态文件
	router.Static("/css", "./public/css")
	router.Static("/js", "./public/js")
	router.Static("/fonts", "./public/fonts")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	// 加载 html
	router.LoadHTMLFiles("public/index.html")
	router.GET("/index.html/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
