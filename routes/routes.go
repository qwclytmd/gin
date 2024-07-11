package routes

import (
	"bcw/app/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORS(), middleware.LanguageHandler())

	RegisterWebRouter(r)

	return r
}
