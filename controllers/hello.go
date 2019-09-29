package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imshakthi/goland/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HelloController interface {
	HelloWorld(ctx *gin.Context)
}

type helloController struct {
	helloService services.HelloService
}

func NewHelloController(helloService services.HelloService) *helloController {
	return &helloController{helloService: helloService}
}

func (ctrl helloController) HelloWorld(ctx *gin.Context) {
	message := ctrl.helloService.Hello()
	log.Info(message)
	SendMessageWithStatus(ctx, http.StatusOK, message)
}
