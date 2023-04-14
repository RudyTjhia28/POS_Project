package main

import (
	"log"
	"pos_project/config"

	"pos_project/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	// Initialize database connection
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Create Gin router
	router := gin.Default()

	// Add endpoints to router
	AddEndpoints(router, db)

	return router

}

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

func main() {
	// Initialize router
	router := SetupRouter()

	// Start server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
