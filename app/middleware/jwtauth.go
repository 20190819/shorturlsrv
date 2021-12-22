package middleware

import (
	"fmt"
	"net/http"
	"shorturlsrv/app/controllers"
	"shorturlsrv/app/services"

	"github.com/gin-gonic/gin"
)

func AuthApi() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		token := string([]byte(authorization)[7:])
		fmt.Println("token", token)
		if token == "" {
			controllers.ApiException(ctx, http.StatusUnauthorized, controllers.CodeFailed, "授权认证失败")
			return
		} else {
			_, cliams, err := services.ParseToken(token)
			if err != nil {
				controllers.ApiException(ctx, http.StatusUnauthorized, controllers.CodeFailed, "授权认证失败")
				return
			}
			ctx.Set("cliams", cliams)
			ctx.Next()
		}
	}
}
