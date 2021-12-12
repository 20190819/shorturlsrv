package controllers

import (
	"github.com/gin-gonic/gin"
)

type shorturl struct{}

var ShortUrl = new(shorturl)

type addModel struct {
	Url  string `json:"url" form:"url" binding:"required"`
	Code string `json:"code"`
}

func (u *shorturl) List(ctx *gin.Context) {

}

func (u *shorturl) Show(ctx *gin.Context) {

}

func (u *shorturl) Store(ctx *gin.Context) {
	add := addModel{}
	ctx.ShouldBind(&add)

}

func (u shorturl) Update(ctx *gin.Context) {

}

func (u shorturl) Destroy(ctx *gin.Context) {

}
