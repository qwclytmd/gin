package middleware

import (
	"bcw/config"
	"github.com/gin-gonic/gin"
	"strings"
)

func LanguageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentLang := strings.TrimSpace(ctx.Request.Header.Get("Lang"))

		config.SetLang(currentLang)

		ctx.Next()
	}
}
