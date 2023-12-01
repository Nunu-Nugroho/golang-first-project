package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nunu-Nugroho/golang-first-project/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Set mode to "release" for better performance in production
	gin.SetMode(gin.ReleaseMode)

	// Setup Gin router
	r := router.SetupRouter()

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the port from the environment variables
	port := os.Getenv("ACTIVE_PORT")
	if port == "" {
		log.Fatal("ACTIVE_PORT is not set in the .env file")
	}
	fmt.Println("Port:", port)

	// Run the server
	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
