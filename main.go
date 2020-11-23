package main

import (
	log "github.com/sirupsen/logrus"
	"go-pdf-poc/db"
	"go-pdf-poc/server"
)

func main() {
	log.Println("Starting appliction....")
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	config := AppConfig{}

	config.LoadEnv()

	db.InitDb()
	log.Println("Intializing server")
	server.InitializeServer()
}
