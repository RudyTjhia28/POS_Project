package apis

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

	product := handlers.NewProductHandler(db)
	router.GET("/products", handlers.JWTMiddleware(), product.GetProducts)
	// router.POST("/addProduct", handlers.JWTMiddleware(), product.CreateProduct)
	// router.GET("/products/:id", handlers.JWTMiddleware(), product.GetProductById)
	// router.PUT("/updateProduct", handlers.JWTMiddleware(), product.UpdateProduct)

	// router.GET("/orders", handlers.GetOrders(db))
	// router.POST("/orders", handlers.CreateOrder(db))
	// router.GET("/orders/:id", handlers.GetOrderById(db))
	// router.PUT("/orders/:id", handlers.UpdateOrder(db))
	// router.DELETE("/orders/:id", handlers.DeleteOrder(db))
}
