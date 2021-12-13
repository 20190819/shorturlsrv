package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess = 0
	CodeFailed  = -1
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "请求成功",
		"data": data,
	})
}

func ApiException(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"code": code,
		"msg":  msg,
	})
}
