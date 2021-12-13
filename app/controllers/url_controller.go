package controllers

import (
	"shorturlsrv/database"
	"shorturlsrv/models"
	"shorturlsrv/utils"

	"github.com/gin-gonic/gin"
)

type shorturl struct{}

var ShortUrl = new(shorturl)

type add struct {
	Url  string `json:"url" form:"url" binding:"required"`
	Code string `json:"code"`
}

func (u *shorturl) List(ctx *gin.Context) {

}

func (u *shorturl) Show(ctx *gin.Context) {

}

func (u *shorturl) Store(ctx *gin.Context) {
	add := add{}
	ctx.ShouldBind(&add)
	umodel := models.Urls{
		Url:  add.Url,
		Code: utils.Str.Random(8),
	}
	database.MysqlClient.Create(&umodel)
	Success(ctx, nil)
}

func (u shorturl) Update(ctx *gin.Context) {

}

func (u shorturl) Destroy(ctx *gin.Context) {

}
