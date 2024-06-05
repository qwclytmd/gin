package services

import (
	"log"
	"mate/app/dto"

	"github.com/gin-gonic/gin"
)

type Login struct{}

func (s Login) CheckLoginAuth(input dto.LoginReq) any {
	log.Printf("%+v", input)
	return gin.H{"name": "duck"}
}
