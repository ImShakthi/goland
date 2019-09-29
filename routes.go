package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imshakthi/goland/controllers"
	"github.com/imshakthi/goland/repositories"
	"github.com/imshakthi/goland/services"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"net/http"
)

func InitiateRoutes(dbConnection *gorm.DB) *gin.Engine {
	router := gin.Default()

	// repository
	userRepo := repositories.NewUserRepo(dbConnection)

	// service
	configService := services.NewConfigService()
	helloService := services.NewHelloService()
	userService := services.NewUserService(configService, userRepo)

	// controller
	helloController := controllers.NewHelloController(helloService)
	userController := controllers.NewUserController(userService)

	authGroup := router.Group("/app")
	{
		// swagger:route GET /app/hello
		// Return Hello World as response.
		// StatusCode 200: Message is fetched successfully
		// responses:
		//  200: string
		authGroup.GET("/hello", helloController.HelloWorld)

		authGroup.GET("/user/:id", userController.GetUser)

		authGroup.POST("/user", userController.CreateUser)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}
