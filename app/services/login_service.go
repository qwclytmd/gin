package services

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Login struct{}

type LoginReq struct {
	Username string
	Password string
}

func (s Login) CheckLoginAuth(input LoginReq) any {
	log.Printf("%+v", input)
	return gin.H{"name": "duck"}
}
