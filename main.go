package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacob-cossette/cv-generator-services/internal/handler"
	"github.com/jacob-cossette/cv-generator-services/internal/middleware"
)

func main() {
	router := gin.Default()

	// Use the authentication middleware for all routes that require authentication.
	router.Use(middleware.AuthenticationMiddleware())

	router.GET("/albums", handler.GetSkills)

	// Create a test request with the "Authorization" header.
	req, _ := http.NewRequest("GET", "/profile", nil)
	req.Header.Set("Authorization", "123") // Set your authentication token here.

	router.Run("localhost:8080")
}
