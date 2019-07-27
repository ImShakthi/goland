package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imshakthi/goland/controllers"
	"github.com/imshakthi/goland/services"
	"net/http"
)

func Init() *gin.Engine {
	router := gin.Default()

	// service
	helloService := services.NewHelloService()

	// controller
	helloController := controllers.NewHelloController(helloService)

	authGroup := router.Group("/app")
	{
		authGroup.GET("/hello", helloController.HelloWorld)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}
