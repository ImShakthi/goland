package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	startServer()
}

func startServer() {
	configureLogging()
	router := InitiateRoutes()
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
