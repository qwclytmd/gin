package controller

import (
	"mate/app/services"

	"github.com/gin-gonic/gin"
)

func CheckLogin(ctx *gin.Context) {
	var input services.LoginReq
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(200, "参数错误")
		return
	}

	s := services.Login{}
	result := s.CheckLoginAuth(input)

	ctx.JSON(200, result)
}
