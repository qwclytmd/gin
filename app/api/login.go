package api

import (
	"bcw/app/common/message"
	"bcw/app/services/login_service"
	"github.com/gin-gonic/gin"
)

// 登录
func Login(ctx *gin.Context) {
	var input login_service.LoginUserRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		message.ResponseAPIResult(ctx, &message.StatusParameterError, nil)
		return
	}

	result, err := login_service.LoginAuthCheck(input)
	if err != nil {
		message.ResponseAPIResult(ctx, err, nil)
		return
	}

	message.ResponseAPIResult(ctx, &message.StatusOK, result)
}

// 登出
func Logout(ctx *gin.Context) {

}
