package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nunu-Nugroho/golang-first-project/route"
	"github.com/joho/godotenv"
)

func main() {
	r := route.SetupRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}	
	port := os.Getenv("ACTIVE_PORT")
	fmt.Println("Port:", port)

	if err := r.Run(":" + port); err != nil {
		// Handle error jika terjadi
		panic(err)
	}
	
}