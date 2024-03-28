package main

import (
	"hacktiv8-techrawih-go-product-sale-concurrency/config"
	"hacktiv8-techrawih-go-product-sale-concurrency/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db := config.GetDBConnection()

	// Set up Gin engine
	r := gin.Default()

	// Register all application routes
	router.RegisterAPIService(r, db)

	// Start the server
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
