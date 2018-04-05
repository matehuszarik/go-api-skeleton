package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthcheckHandler ..
func HealthcheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
