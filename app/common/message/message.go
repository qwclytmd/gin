package message

import (
	"bcw/libs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseAPIResult(ctx *gin.Context, err *StatusCode, data any) {
	ctx.JSON(http.StatusOK, gin.H{"code": &err.Code, "msg": libs.GetLanguageMsg(err.Lang) + " " + err.Desc, "data": data})
}
