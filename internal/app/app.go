package app

import (
	"dev/reglogauth/internal/config"
	"dev/reglogauth/internal/database"
	"dev/reglogauth/internal/services"
)

func init() {
	config.Init()
	database.Init()
}

func Run() {
	r := services.SetupRouter()
	r.Run("localhost:8080")
}
