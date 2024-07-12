package services

import (
	_ "dev/reglogauth/docs"
	"dev/reglogauth/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	auth := router.Group("/api/auth")
	{
		auth.GET("/ping", Logger(), handlers.Ping)
		auth.POST("/reg", handlers.Registration)
		auth.POST("/log", handlers.Authentication)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
