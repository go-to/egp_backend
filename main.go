package main

import (
	"github.com/go-to/egp-backend/router"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	loadEnv()

	apiPortStr := os.Getenv("API_PORT")
	apiPort, err := strconv.Atoi(apiPortStr)
	if err != nil {
		log.Fatal(err)
	}

	router.New(apiPort)
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
