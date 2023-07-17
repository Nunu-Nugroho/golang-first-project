package route

import (
	"github.com/Nunu-Nugroho/golang-first-project/models"
	"github.com/Nunu-Nugroho/golang-first-project/controllers/productcontroller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	models.ConnectDatabase()	

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)
	return r		
}