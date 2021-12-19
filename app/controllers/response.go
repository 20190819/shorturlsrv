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

func ApiException(ctx *gin.Context, code int, msg string) {
	status := 0
	if code == -1 {
		status = http.StatusOK
	} else if code == 500 {
		status = http.StatusInternalServerError
	}
	ctx.AbortWithStatusJSON(status, gin.H{
		"code": code,
		"msg":  msg,
	})
}
