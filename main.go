package main

import (
	"log"
	"pos_project/apis"
	"pos_project/config"
	constant "pos_project/config/constants"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Initialize database connection
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Create Gin router
	router := gin.Default()

	router.ForwardedByClientIP = true

	router.SetTrustedProxies([]string{
		constant.ProxyLocal,
	})

	// Add endpoints to router
	apis.AddEndpoints(router, db)

	return router

}

func main() {
	// Initialize router
	router := SetupRouter()

	// Start server
	err := router.Run(constant.ServicePort)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
