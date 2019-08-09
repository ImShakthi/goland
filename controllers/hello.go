package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imshakthi/goland/models"
	"github.com/imshakthi/goland/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HelloController interface {
	HelloWorld(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
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

func (ctrl helloController) CreateUser(ctx *gin.Context) {
	var user models.UserRequest
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)
	log.Printf("After bind body=%v", err)
	ctx.Done()
}
