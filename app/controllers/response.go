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

func Failed(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"code": CodeFailed,
		"msg":  "系统异常",
	})
}
