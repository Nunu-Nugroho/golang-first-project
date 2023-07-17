package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nunu-Nugroho/golang-first-project/models"
	"github.com/Nunu-Nugroho/golang-first-project/controllers/productcontroller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()	

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)
	
	// port := "1005" // Ganti dengan port yang diinginkan
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