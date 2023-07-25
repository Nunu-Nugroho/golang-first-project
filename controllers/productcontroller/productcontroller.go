package productcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Nunu-Nugroho/golang-first-project/models"
	"github.com/Nunu-Nugroho/golang-first-project/package"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product

	packages.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"product": products})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := packages.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	packages.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	// var product models.Product
	// id := c.Param("id")

	// if err := c.ShouldBindJSON(&product); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// if packages.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal update product"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "data berhasil terupdate"})
	// c.JSON(http.StatusOK, gin.H{"product": product})
	var product models.Product
	id := c.Param("id")
	
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if packages.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func Delete(c *gin.Context) {
	var product models.Product
	// input := map[string]string{"id": "0"}
	var input struct {
		Id json.Number
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	id, _ := strconv.Atoi(string(input.Id))
	// id, _ := input.Id.Int64()
	if packages.DB.Delete(&product, id).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "gagal hapus data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "hapus data berhasil"})
}
