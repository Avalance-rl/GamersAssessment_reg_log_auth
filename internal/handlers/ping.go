package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /

// Ping godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags ping
// @Accept json
// @Produce plain
// @Success 200 plain pong
// @Router /api/auth/ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
