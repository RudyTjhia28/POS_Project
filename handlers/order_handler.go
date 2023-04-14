package handlers

import (
	"net/http"

	"pos_project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []models.Order
		db.Find(&orders)

		c.JSON(http.StatusOK, gin.H{"data": orders})
	}
}

// func CreateOrder(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var input models.CreateOrderInput
// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		order := models.Order{ProductID: input.ProductID, Quantity: input.Quantity}
// 		db.Create(&order)

// 		c.JSON(http.StatusOK, gin.H{"data": order})
// 	}
// }
