package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imshakthi/goland/controllers"
	"github.com/imshakthi/goland/services"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"net/http"
)

func InitiateRoutes() *gin.Engine {
	router := gin.Default()

	// service
	configService := services.NewConfigService()
	helloService := services.NewHelloService()
	userService := services.NewUserService(configService)

	// controller
	helloController := controllers.NewHelloController(helloService)
	userController := controllers.NewUserController(userService)

	psqlInfo := getDatabaseInfo(configService)
	logrus.Info("psqlInfo=", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("Error while creating database connection " + err.Error())
	}
	defer db.Close()

	authGroup := router.Group("/app")
	{
		// swagger:route GET /app/hello
		// Return Hello World as response.
		// StatusCode 200: Message is fetched successfully
		// responses:
		//  200: string
		authGroup.GET("/hello", helloController.HelloWorld)

		authGroup.GET("/user", userController.GetUser)

		authGroup.POST("/user", userController.CreateUser)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}

func getDatabaseInfo(configService services.ConfigService) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configService.GetDatabaseHost(),
		configService.GetDatabasePort(),
		configService.GetDatabaseUserName(),
		configService.GetDatabasePassword(),
		configService.GetDatabaseName())

}
