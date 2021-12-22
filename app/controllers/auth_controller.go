package controllers

import (
	"errors"
	"net/http"
	"shorturlsrv/app/services"
	"shorturlsrv/database"
	"shorturlsrv/models"

	"github.com/gin-gonic/gin"
)

type auth struct{}

var AuthController = new(auth)

func (au auth) Register(ctx *gin.Context) {
	data := new(struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	})
	ctx.ShouldBind(data)
	oldUser := models.UserModel{}
	rows := database.MysqlClient.Where("email=?", data.Email).First(&oldUser).RowsAffected
	if rows > 0 {
		ApiException(ctx, http.StatusOK, CodeFailed, errors.New("邮箱被占用").Error())
		return
	}
	hashStr, err := models.UserModel{}.Hash(data.Password)
	if err != nil {
		ApiException(ctx, http.StatusOK, CodeFailed, err.Error())
		return
	}
	data.Password = hashStr

	user := models.UserModel{}
	user.Email = data.Email
	user.Password = data.Password
	err = database.MysqlClient.Create(&user).Error
	if err != nil {
		ApiException(ctx, http.StatusOK, CodeFailed, err.Error())
		return
	}
	Success(ctx, user)
}

func (au auth) Login(ctx *gin.Context) {
	data := struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}{}
	ctx.ShouldBind(&data)
	user := models.UserModel{}
	rows := database.MysqlClient.Where("email=?", data.Email).First(&user).RowsAffected
	if rows == 0 {
		ApiException(ctx, http.StatusOK, CodeFailed, errors.New("邮箱不正确").Error())
		return
	}

	err := user.CheckPwd(user.Password, data.Password)
	if err != nil {
		ApiException(ctx, http.StatusOK, CodeFailed, "密码不正确")
		return
	}

	// 校验通过
	token, err := services.CreateToken(user.Email)
	if err != nil {
		ApiException(ctx, http.StatusOK, CodeFailed, errors.New("创建token错误").Error())
		return
	}

	resp := struct {
		Token string           `json:"token"`
		User  models.UserModel `json:"user"`
	}{}
	resp.Token = token
	resp.User = user
	Success(ctx, resp)
}
