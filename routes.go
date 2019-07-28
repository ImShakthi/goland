package main

import (
	"github.com/gin-gonic/gin"
	"goland/controllers"
	"goland/services"
	"net/http"
)

func InitiateRoutes() *gin.Engine {
	router := gin.Default()

	// service
	helloService := services.NewHelloService()

	// controller
	helloController := controllers.NewHelloController(helloService)

	authGroup := router.Group("/app")
	{
		// swagger:route GET /app/hello
		// Return Hello World as response.
		// StatusCode 200: Message is fetched successfully
		// responses:
		//  200: string
		authGroup.GET("/hello", helloController.HelloWorld)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}
