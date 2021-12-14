package controllers

import (
	"shorturlsrv/database"
	"shorturlsrv/models"
	"shorturlsrv/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type shorturl struct{}

var ShortUrl = new(shorturl)

type add struct {
	Url  string `json:"url" form:"url" binding:"required"`
	Code string `json:"code"`
}

func (u *shorturl) List(ctx *gin.Context) {
	var data []models.Urls
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	database.MysqlClient.Paging(pageInt, limitInt).Find(&data)
	Success(ctx, data)
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
	Success(ctx, umodel)
}

func (u shorturl) Update(ctx *gin.Context) {

}

func (u shorturl) Destroy(ctx *gin.Context) {
	var req struct {
		Id string `uri:"id"`
	}
	ctx.ShouldBindUri(&req)
	database.MysqlClient.Delete(&models.Urls{}, req.Id)
	Success(ctx, nil)
}
