package main

import (
	"fmt"
	"github.com/imshakthi/goland/services"
	"github.com/jinzhu/gorm"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	startServer()
}

func startServer() {
	configureLogging()
	configService := services.NewConfigService()
	dbConnection, dbErr := gorm.Open("postgres", getGormDatabaseInfo(configService))
	if dbErr != nil {
		log.Fatalln("error in starting postgres server: ", dbErr)
	}
	defer func() {
		err := dbConnection.Close()
		if err != nil {
			log.Error("unable to close database connection due to: ", err)
		}
	}()

	router := InitiateRoutes(dbConnection)
	err := router.Run()
	if err != nil {
		log.Error("Server startup failed due to ", err)
	}
}

func configureLogging() *log.JSONFormatter {
	formatter := new(log.JSONFormatter)
	formatter.PrettyPrint = true
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	return formatter
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
