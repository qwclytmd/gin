package routes

import (
	"bcw/app/api"
	"bcw/app/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterWebRouter(router *gin.Engine) {
	router.POST("/login", api.Login)

	router.Use(middleware.CheckAuth())
	{
		router.GET("/account/list", api.AccountList)
		router.GET("/player/list", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"code": 200, "msg": "/player/list", "data": nil})
		})
		router.GET("/player/detail", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"code": 200, "msg": "/player/detail", "data": nil})
		})
		router.POST("/logout", api.Logout)
	}
}
