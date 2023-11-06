package routes

import (
	"github.com/gin-gonic/gin"
	c "urlpro/controllers/url"
)

// InitRoutes initializes the application routes and starts the server.
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

	// Start the Gin server on the specified host and port
	if err := router.Run(hostURL); err != nil {
		panic(err)
	}
}

// getConfig returns the host and port configuration.
func getConfig() (string, error) {
	// You can load the configuration from a file, environment variable, or any other source.
	// Here, we're using a hard-coded value for demonstration purposes.
	return ":8084", nil
}
