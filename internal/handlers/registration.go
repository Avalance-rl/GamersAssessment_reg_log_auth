package handlers

import (
	"dev/reglogauth/internal/http_responses"
	"dev/reglogauth/internal/models"

	"github.com/gin-gonic/gin"
)

func Registration(c *gin.Context) {
	body := models.RegisterRequest{}
	if c.Bind(&body) != nil {
		http_responses.FailToReadBody(c)
		return
	}

}
