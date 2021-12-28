package controllers

import (
	"shorturlsrv/database"
	"shorturlsrv/models"
	"shorturlsrv/utils"

	"github.com/gin-gonic/gin"
)

type shorturl struct{}

var ShortUrlController = shorturl{}

func (u *shorturl) List(ctx *gin.Context) {
	var data []models.UrlModel
	var count int64
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	database.MysqlClient.Paging(page, limit).Order("id desc").Find(&data)
	database.MysqlClient.Model(&models.UrlModel{}).Count(&count)
	SuccessWithCount(ctx, data, count)
}

func (u *shorturl) Show(ctx *gin.Context) {

}

func (u *shorturl) Store(ctx *gin.Context) {
	data := struct {
		Url string `json:"url" form:"url" binding:"required"`
	}{}
	ctx.ShouldBind(&data)
	umodel := models.UrlModel{
		Url:  data.Url,
		Code: utils.Str.Random(8),
	}
	database.MysqlClient.FirstOrCreate(&umodel, models.UrlModel{Url: data.Url})
	Success(ctx, umodel)
}

func (u shorturl) Update(ctx *gin.Context) {

}

func (u shorturl) Destroy(ctx *gin.Context) {
	var req struct {
		Id string `uri:"id"`
	}
	ctx.ShouldBindUri(&req)
	database.MysqlClient.Delete(&models.UrlModel{}, req.Id)
	Success(ctx, nil)
}
