package services

import (
	_ "dev/reglogauth/docs"
	"dev/reglogauth/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/api/auth")
	{
		auth.GET("/ping", Logger(), handlers.Ping)
		auth.POST("/reg", handlers.Registration)
		auth.POST("/log", handlers.Authentication)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
