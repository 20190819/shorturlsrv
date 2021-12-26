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
	resp := gin.H{
		"code": 0,
		"msg":  "请求成功",
	}
	if data != nil {
		resp["data"] = data
	}
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func SuccessWithCount(ctx *gin.Context, data interface{}, count int64) {
	resp := gin.H{
		"code": 0,
		"msg":  "请求成功",
	}
	resp["count"] = count
	if data != nil {
		resp["data"] = data
	}
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}

func ApiException(ctx *gin.Context, status, code int, msg string) {
	ctx.AbortWithStatusJSON(status, gin.H{
		"code": code,
		"msg":  msg,
	})
}
