package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacob-cossette/cv-generator-services/internal/handler"
	"github.com/jacob-cossette/cv-generator-services/internal/middleware"
	"github.com/jacob-cossette/cv-generator-services/internal/util"
)

func main() {
	router := gin.Default()

	// Use the authentication middleware for all routes that require authentication.
	router.Use(middleware.AuthenticationMiddleware())

	router.GET("/albums", handler.GetSkills)

	user := User{ID: "143542634643643", Name: "jacob"}

	token, err := util.GenerateToken(user.ID)

	fmt.Printf("Error generating token: %v\n", err)

	fmt.Printf("the token is : %v\n", token)
	// Create a test request with the "Authorization" header.
	req, _ := http.NewRequest("GET", "/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token) // Set your authentication token here.

	router.Run("localhost:8080")
}

type User struct {
	ID   string
	Name string
	// Other fields...
}
