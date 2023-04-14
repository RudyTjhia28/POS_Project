package handlers

import (
	"net/http"

	"pos_project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		db.Find(&products)

		c.JSON(http.StatusOK, gin.H{"data": products})
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.CreateProductRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product := models.Product{Name: input.Name, Price: input.Price, Quantity: input.Quantity}
		db.Create(&product)

		c.JSON(http.StatusOK, gin.H{"data": product})
	}
}

func GetProductById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := db.First(&product, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": product})
	}
}

// func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var product models.Product
// 		if err := db.First(&product, c.Param("id")).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 			return
// 		}

// 		var input models.UpdateProductInput
// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		db.Model(&product).Updates(input)

// 		c.JSON(http.StatusOK, gin.H{"data": product})
// 	}
// }

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := db.First(&product, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		db.Delete(&product)

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
