package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imshakthi/goland/models"
	"github.com/imshakthi/goland/services"
	log "github.com/sirupsen/logrus"
)

type UserController interface {
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (ctrl userController) CreateUser(ctx *gin.Context) {
	var user models.UserRequest
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)
	log.Printf("After bind body=%v", err)
	ctx.Done()
}

func (ctrl userController) GetUser(ctx *gin.Context) {
	ctx.Done()
}
