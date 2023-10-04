package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacob-cossette/cv-generator-services/internal/handler"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, handler.Skills)
}
