package main

import (
	"github.com/Dubjay18/gobank2/app"
	"github.com/Dubjay18/gobank2/logger"
)

func main() {
	//log.Println("Starting the application...")
	app.GetEnvVar()
	logger.Info("Starting the application...")
	app.SanityCheck()
	app.Start()
}
