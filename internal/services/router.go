package services

import (
	"dev/reglogauth/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.GET("/ping", Logger(), handlers.Ping)
		auth.POST("/reg", handlers.Registration)
		auth.POST("/log", handlers.Authentication)
	}

	return r
}
