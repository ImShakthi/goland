package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

)

type HelloController interface {
	HelloWorld(ctx *gin.Context)
}

type helloController struct {
}

func NewHelloController() HelloController {
	return &helloController{}
}

func (ctrl helloController) HelloWorld(ctx *gin.Context) {
	log.Info("Hello world!")
	return
}
