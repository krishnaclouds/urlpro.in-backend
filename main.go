package main

import (
	"github.com/gin-gonic/gin"
	c "urlpro/controller"
)

func main() {
	// Get the url port number from the configuration function
	_, port := getConfig()

	// Gin router with default middleware
	router := gin.Default()

	// Routes for the application
	router.POST("/generate", c.GenerateShortURL)
	router.GET("/:shortURL", c.RedirectToDesiredURL)

	// Start the Gin server on the specified port
	if err := router.Run(port); err != nil {
		// Handle any errors that occur while starting the server
		panic(err)
	}
}

func getConfig() (string, string) {
	return "", ":8084"
}
