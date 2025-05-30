package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logFile, err := os.Create("app.log")
	if err != nil {
		log.Fatal("failed to create log file")
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("App started")

	// load environment variable form file and set to file

	// retrieve and print the environment variable
	// Akses variabel environment
	appEnv := os.Getenv("APP_ENV")
	fmt.Println("APP_ENV:", appEnv)

	debug := os.Getenv("APP_DEBUG")
	fmt.Println("DEBUG:", debug)

}
