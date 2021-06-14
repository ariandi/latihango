package main

import (
	"latihan1/app"
	"latihan1/logger"
	"log"
	"os"
)

func SanityCheck()  {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("environment variable not defined")
	}
}

func main()  {

	SanityCheck()
	// log.Println("Application started....!!!!")
	logger.Info("Application started....!!!!")
	app.Start()
}
