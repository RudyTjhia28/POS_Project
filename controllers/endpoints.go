package controllers

import (
	"pos_project/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddEndpoints(router *gin.Engine, db *gorm.DB) {
	// Define endpoints here...

	router.GET("/products", handlers.GetProducts(db))
	router.POST("/products", handlers.CreateProduct(db))
	router.GET("/products/:id", handlers.GetProductById(db))
	// router.PUT("/products/:id", handlers.UpdateProduct(db))
	// router.DELETE("/products/:id", handlers.DeleteProduct(db))

	// router.GET("/orders", handlers.GetOrders(db))
	// router.POST("/orders", handlers.CreateOrder(db))
	// router.GET("/orders/:id", handlers.GetOrderById(db))
	// router.PUT("/orders/:id", handlers.UpdateOrder(db))
	// router.DELETE("/orders/:id", handlers.DeleteOrder(db))
}
