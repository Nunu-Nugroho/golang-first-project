package main

import (
	"github.com/Nunu-Nugroho/golang-first-project/models"
	"github.com/Nunu-Nugroho/golang-first-project/controllers/productcontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()	

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)
	
	port := "3000" // Ganti dengan port yang diinginkan

	if err := r.Run(":" + port); err != nil {
		// Handle error jika terjadi
		panic(err)
	}
}