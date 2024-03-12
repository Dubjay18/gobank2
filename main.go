package main

import (
	"github.com/Dubjay18/gobank2/app"
	"github.com/Dubjay18/gobank2/logger"
)

func main() {
	//log.Println("Starting the application...")
	logger.Info("Starting the application...")
	app.Start()
}
