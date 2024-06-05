package routes

import (
	"mate/app/controller"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	//登录
	r.POST("/login", controller.CheckLogin)
	test := r.Group("/test")
	{
		test.GET("/index", controller.CheckLogin)
	}

}
