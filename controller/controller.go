package controller

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type URLData struct {
	OriginalURL string `json:"originalURL"`
	BaseURL     string `json:"baseURL"`
	Expiry      int    `json:"expiry"`
}

// GenerateShortURL generates a short URL for the given OriginalURL.
// It uses SHA-256 hash to create a unique short URL.
// TODO: Handle BaseURL and Expiry Scenarios
func GenerateShortURL(c *gin.Context) {
	fmt.Println("In Short URL Generator")

	var urlData URLData

	// Bind JSON data from the request body to URLData struct
	if err := c.ShouldBindJSON(&urlData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	originalURL := urlData.OriginalURL
	fmt.Println("Original URL:", originalURL)

	// Generate a unique short URL using SHA-256 hash
	shortURL, err := generateShortURL(originalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"originalURL": originalURL,
		"shortURL":    "https://urlpros.in/" + shortURL,
	})
}

func generateShortURL(originalURL string) (string, error) {
	// Create a SHA-256 hash of the original URL
	hash := sha256.New()
	if _, err := hash.Write([]byte(originalURL)); err != nil {
		return "", err
	}

	// Encode the hash as a base64 URL-safe string
	encoded := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	// Use the first 8 characters of the encoded hash as the short URL
	shortURL := encoded[:8]

	return shortURL, nil
}

func RedirectToDesiredURL(c *gin.Context) {
	fmt.Println("Redirecting to Desired URL")
}
