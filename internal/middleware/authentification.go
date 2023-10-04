package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacob-cossette/cv-generator-services/internal/util" // Import the token utility package.
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		userID, err := util.VerifyToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
