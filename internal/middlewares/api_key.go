package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HeaderAPIKey ...
var HeaderAPIKey = "X-API-Key"

// APIKey ...
func APIKey(allowedKeys []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get(HeaderAPIKey)
		log.Println(apiKey)

		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Missing " + HeaderAPIKey + " header",
			})
			return
		}

		for _, key := range allowedKeys {
			if key == apiKey {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid API Key",
		})
		return
	}
}
