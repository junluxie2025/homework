package main

import (
	"homework/go-basic/task4/config"
	"homework/go-basic/task4/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found, using system environment variables")
	}

	//init log
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Starting blog application...")

	//init database
	config.InitDatabase()

	// set up routers
	r := routers.SetupRouters()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	port = ":" + port

	logrus.WithField("port", port).Info("blog server starting")

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start blog application")
	}

}
