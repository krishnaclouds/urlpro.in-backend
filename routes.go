package main

import (
	"github.com/gin-gonic/gin"
	c "urlpro/controllers/url"
)

func InitRoutes() {

	// Get the url port number from the configuration function
	hostURL, err := getConfig()
	if err != nil {
		panic(err)
	}

	// Gin router with default middleware
	router := gin.Default()

	// Routes for the application
	router.POST("/generate", c.GenerateShortURL)
	router.GET("/:shortURL", c.RedirectToDesiredURL)

	// Start the Gin server on the specified port
	if err := router.Run(hostURL); err != nil {
		// Handle any errors that occur while starting the server
		panic(err)
	}

}

// getConfig returns url of the application.
func getConfig() (string, error) {
	return ":8084", nil
}
