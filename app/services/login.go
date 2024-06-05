package services

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Login struct{}

func (s Login) CheckLoginAuth(input any) any {
	log.Printf("%+v", input)
	return gin.H{"name": "duck"}
}
