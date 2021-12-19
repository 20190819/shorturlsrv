package middleware

import (
	"errors"
	"shorturlsrv/app/controllers"
	"shorturlsrv/database"
	"shorturlsrv/models"

	"github.com/gin-gonic/gin"
)

func CheckSecret() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.GetHeader("x-key")
		sec := ctx.GetHeader("x-secret")
		err := validate(key, sec)
		if err != nil {
			controllers.ApiException(ctx, controllers.CodeFailed, err.Error())
			return
		}
		ctx.Next()
	}
}

func validate(key, sec string) error {
	if key == "" || sec == "" {
		return errors.New("Key or secret is missing")
	}
	secretRow := models.SecretModel{}

	if err := database.MysqlClient.Where("`key`=?", key).First(&secretRow).Error; err != nil {
		return err
	}
	if secretRow.Secret != sec {
		return errors.New("Key or secret not right")
	}
	return nil
}
