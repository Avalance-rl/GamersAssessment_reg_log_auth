package services

import (
	"dev/reglogauth/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handlers.Ping)
	r.POST("/reg", handlers.Registration)
	r.POST("/log", handlers.Authentication)

	return r
}
