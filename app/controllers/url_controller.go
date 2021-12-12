package controllers

import (
	"github.com/gin-gonic/gin"
)

type shorturl struct{}

var ShortUrl = new(shorturl)

func (u *shorturl) List(ctx *gin.Context) {
	Success(ctx, nil)
}

func (u *shorturl) Show(ctx *gin.Context) {

}

func (u *shorturl) Store(ctx *gin.Context) {

}

func (u shorturl) Update(ctx *gin.Context) {

}

func (u shorturl) Destroy(ctx *gin.Context) {

}
