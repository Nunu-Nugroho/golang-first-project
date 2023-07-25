package route

import (
	"github.com/Nunu-Nugroho/golang-first-project/models"
	"github.com/Nunu-Nugroho/golang-first-project/package"
	"github.com/Nunu-Nugroho/golang-first-project/controllers/productcontroller"
	"github.com/Nunu-Nugroho/golang-first-project/controllers/objectcontroller"
	"github.com/Nunu-Nugroho/golang-first-project/controllers/miniocontroller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	models.ConnectDatabase()	
	minio.MinioConnection()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	
	// minio
	r.GET("/test-minio", miniocontroller.Test)
	r.POST("/upload-image", miniocontroller.CreateImage)
	r.GET("/image/:path", miniocontroller.GetImage)
	r.GET("/list-contents", miniocontroller.ListContent)
	r.GET("/list-contents/:folder", miniocontroller.ListContentbyFolder)
	return r		
}