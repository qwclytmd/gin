package middleware

import (
	"bcw/app/common/message"
	"bcw/config"
	"bcw/libs"
	"bcw/server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"strings"
)

var (
	tokenManager libs.TokenManager
)

func init() {
	tokenManager = libs.TokenManager{SecretKey: config.GetViper().GetString("server.jwt_secret_key")}
}

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//token 验证
		tokenString := strings.TrimPrefix(ctx.Request.Header.Get("Authorization"), "Bearer ")

		if tokenString == "" {
			message.ResponseAPIResult(ctx, &message.StatusTokenCannotBeEmpty, nil)
			ctx.Abort()
			return
		}

		claims, err := tokenManager.ParseToken(tokenString)
		if err != nil {
			message.ResponseAPIResult(ctx, &message.StatusTokenIsInvalid, nil)
			ctx.Abort()
			return
		}

		//权限验证
		ok, err := server.HttpServers.Casbin.Enforce(cast.ToString(claims.RId), ctx.Request.URL.Path, "*")

		if err != nil {
			log.Println(err)
			message.ResponseAPIResult(ctx, &message.StatusInternalError, nil)
			ctx.Abort()
			return
		}
		if !ok {
			message.ResponseAPIResult(ctx, &message.StatusUnAuthorized, nil)
			ctx.Abort()
			return
		}

		ctx.Set("operator_id", claims.ID)
		ctx.Next()

	}
}
