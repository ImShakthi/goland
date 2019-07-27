package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imshakthi/goland/controller"
	"net/http"
)

func Init() *gin.Engine {
	router := gin.Default()

	helloController := controller.NewHelloController()

	authGroup := router.Group("/app")
	{
		authGroup.GET("/hello", helloController.HelloWorld)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}
