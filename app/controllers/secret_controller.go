package controllers

import (
	"fmt"
	"shorturlsrv/database"
	"shorturlsrv/models"
	"shorturlsrv/utils"

	"github.com/gin-gonic/gin"
)

type secret struct{}

var SecretController = secret{}

func (sec secret) List(ctx *gin.Context) {
	var data []models.SecretModel
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	database.MysqlClient.Paging(page, limit).Find(&data)
	Success(ctx, data)
}

func (sec secret) Store(ctx *gin.Context) {
	data := models.SecretModel{
		Key:    utils.Str.Random(8),
		Secret: utils.Str.Random(16),
	}
	fmt.Println("Create secret", data)
	database.MysqlClient.Create(&data)
	Success(ctx, data)
}

func (sec secret) Destroy(ctx *gin.Context) {
	data := new(struct {
		Id string `uri:"id"`
	})

	ctx.ShouldBindUri(data)
	database.MysqlClient.Delete(new(models.SecretModel), data.Id)
	Success(ctx, nil)
}
