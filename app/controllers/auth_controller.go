package controllers

import (
	"errors"
	"fmt"
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
		ApiException(ctx, CodeFailed, errors.New("邮箱被占用").Error())
		return
	}
	hashStr, err := models.UserModel{}.Hash(data.Password)
	if err != nil {
		ApiException(ctx, CodeFailed, err.Error())
		return
	}
	data.Password = hashStr

	user := models.UserModel{}
	user.Email = data.Email
	user.Password = data.Password
	err = database.MysqlClient.Create(&user).Error
	if err != nil {
		ApiException(ctx, CodeFailed, err.Error())
		return
	}
	Success(ctx, user)
}

func (au auth) Login(ctx *gin.Context) {
	data := new(struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	})
	ctx.ShouldBind(data)

	fmt.Printf("Login receive data:%+v\n", *data)
}
