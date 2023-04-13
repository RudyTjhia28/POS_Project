package main

import (
	"log"
	"pos_project/config"
	"pos_project/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Custom middleware to log endpoint
	r.Use(func(c *gin.Context) {
		log.Println(c.Request.Method, c.Request.URL)
		c.Next()
	})

	r.GET("/products", func(c *gin.Context) {
		db, err := config.ConnectDatabase()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to connect to database",
			})
			return
		}

		var products []models.Product
		db.Find(&products)

		c.JSON(200, products)
	})

	return r
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.SetPrefix("INFO: ")
	log.Println("server runs in port 8080")
}
