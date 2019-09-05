package main

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/imshakthi/goland/controllers"
	"github.com/imshakthi/goland/models"
	"github.com/imshakthi/goland/services"
	"net/http"
)

func InitiateRoutes() *gin.Engine {
	router := gin.Default()

	configData := loadConfigFromToml()
	// service
	helloService := services.NewHelloService(configData)

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

		authGroup.POST("/user", helloController.CreateUser)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}

func loadConfigFromToml() models.Config {
	var config models.Config
	_, err := toml.DecodeFile("configurations/config.toml", &config)
	if err != nil {
		panic("Error while loading TOML data=" + err.Error())
	}
	return config
}
