package http_responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FailToReadBody(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Failed to read body",
	})
}
