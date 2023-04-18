package controllers

import (
	"pos_project/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddEndpoints(router *gin.Engine, db *gorm.DB) {
	// Define endpoints here...

	router.POST("/login", func(c *gin.Context) {
		handlers.Login(c, db)
	})

	router.GET("/products", handlers.JWTMiddleware(), handlers.GetProducts(db))
	router.POST("/addProduct", handlers.JWTMiddleware(), handlers.CreateProduct(db))
	router.GET("/products/:id", handlers.JWTMiddleware(), handlers.GetProductById(db))
	router.PUT("/updateProduct", handlers.JWTMiddleware(), handlers.UpdateProduct(db))

	// router.GET("/orders", handlers.GetOrders(db))
	// router.POST("/orders", handlers.CreateOrder(db))
	// router.GET("/orders/:id", handlers.GetOrderById(db))
	// router.PUT("/orders/:id", handlers.UpdateOrder(db))
	// router.DELETE("/orders/:id", handlers.DeleteOrder(db))
}
