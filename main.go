package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jacob-cossette/cv-generator-services/internal/handler"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handler.GetSkills)
	router.Run("localhost:8080")
}
