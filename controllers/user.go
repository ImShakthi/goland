package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/imshakthi/goland/models"
	"github.com/imshakthi/goland/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserController interface {
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
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
	var user models.UserDTO
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)

	err = ctrl.userService.AddUser(user)
	if err != nil {
		log.Printf("unable to create user : %+v", err)
		ctx.Abort()
		return
	}
	ctx.Done()
}

func (ctrl userController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := ctrl.userService.GetUser(id)
	if err != nil {
		log.Errorf("error response: %s", err)
		ctx.JSON(http.StatusNotFound, err)
	} else {
		ctx.JSON(http.StatusOK, user)
	}
}

func (ctrl userController) GetUsers(ctx *gin.Context) {
	users, err := ctrl.userService.GetUsers()
	if err != nil {
		log.Errorf("error response: %s", err)
		ctx.JSON(http.StatusNotFound, err)
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}
