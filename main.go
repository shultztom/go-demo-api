package main

import (
	"github.com/gin-gonic/gin"
	"shultzlab.com/go-demo-api/controllers"
)

// main go function
func main() {
	router := gin.Default()

	router.GET("/health", controllers.Health)

	// Run the server on port 8080 and log errors if any
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
