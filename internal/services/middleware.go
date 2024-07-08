package services

import (
	"dev/reglogauth/internal/database"
	"dev/reglogauth/internal/handlers"
	"dev/reglogauth/internal/http_responses"
	"log/slog"
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshID, err := c.Cookie("refresh_token")
		if err != nil {
			http_responses.FailRefreshTokenMissing(c)
			slog.Error("Logger() error = %v", err)
			return
		}
		token := database.FindToken(refreshID)
		if token == nil {
			return
		}
		if math.Abs(token["creation_time"].(time.Time).Sub(time.Now()).Seconds()) >= handlers.RefreshLife {
			http_responses.FailRefreshTokenIsObsolete(c)
			return
		}

		c.Next()
	}
}
