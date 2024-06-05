package routes

import (
	"mate/app/controller"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	
	test := r.Group("/test")
	{
		test.GET("/index", controller.CheckLogin)
	}

}