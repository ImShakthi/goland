package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imshakthi/goland/controllers"
	"github.com/imshakthi/goland/repositories"
	"github.com/imshakthi/goland/services"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	//_ "github.com/go-pg/pg"
	_ "github.com/lib/pq"
	"net/http"
)

func InitiateRoutes() *gin.Engine {
	router := gin.Default()

	configService := services.NewConfigService()
	dbConnection := getGormConnection(configService)

	userRepo := repositories.NewUserRepo(dbConnection)

	// service
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

		authGroup.GET("/user", userController.GetUser)

		authGroup.POST("/user", userController.CreateUser)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})

	return router
}

func getGormConnection(configService services.ConfigService) *gorm.DB {
	db, err := gorm.Open("postgres", getGormDatabaseInfo(configService))
	if err != nil {
		log.Fatalln("error in starting postgres server: ", err)
	}
	defer db.Close()

	return db
}

func getGormDatabaseInfo(configService services.ConfigService) string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		configService.GetDatabaseHost(),
		configService.GetDatabasePort(),
		configService.GetDatabaseUserName(),
		configService.GetDatabaseName(),
		configService.GetDatabasePassword())
}

//func getDatabaseConnection(configService services.ConfigService) *sql.DB {
//	connStr := getDatabaseHostInfo(configService)
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	return db
//}
//func getDatabaseHostInfo(configService services.ConfigService) string {
//	return fmt.Sprintf("user=%s dbname=%s sslmode=verify-full",
//		configService.GetDatabaseUserName(),
//		configService.GetDatabaseName())
//}

//func getDatabaseConnection(configService services.ConfigService) *pg.DB {
//	options := pg.Options{
//		Addr:     getDatabaseHostInfo(configService),
//		User:     configService.GetDatabaseUserName(),
//		Password: configService.GetDatabasePassword(),
//		Database: configService.GetDatabaseName(),
//	}
//	return pg.Connect(&options)
//}
//
//func getDatabaseHostInfo(configService services.ConfigService) string {
//	return fmt.Sprintf("%s:%d",
//		configService.GetDatabaseHost(),
//		configService.GetDatabasePort())
//
//}
