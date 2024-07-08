package handlers

import (
	"dev/reglogauth/internal/config"
	"dev/reglogauth/internal/database"
	"dev/reglogauth/internal/http_responses"
	"dev/reglogauth/internal/models"
	"log"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	AccessLife  = 60 * 30
	RefreshLife = 3600 * 24
)

var jwtSecretKey = []byte(config.CFG.JwtSecretKey)

// @BasePath /auth/api/

// Authentication godoc
// @Summary authenticates the user
// @Schemes application/json
// @Description accepts json sent by the user as input and authorize it
// @Tags authentication
// @Accept json
// @Produce json
// @Param input body models.AuthenticationRequest true "account info"
// @Success 200 "message: Authentication was successful; access_token"
// @Failure 400 "error: Failed to read body"
// @Failure 403 "fail": "You entered the wrong password or email"
// @Router /api/auth/log [post]
func Authentication(c *gin.Context) {
	body := models.AuthenticationRequest{}
	if c.Bind(&body) != nil {
		http_responses.FailToReadBody(c)
		return
	}
	password := database.FindUser(body.Email)
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password))
	if err != nil {
		http_responses.FailWrongPassword(c)
		return
	}
	payload := jwt.MapClaims{
		"sub": body.Email,
		"exp": AccessLife,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedAccessToken, err := accessToken.SignedString(jwtSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	refreshPayload := jwt.MapClaims{
		"sub": rand.Int(),
		"exp": RefreshLife,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshPayload)
	signedRefreshToken, err := refreshToken.SignedString(jwtSecretKey)

	refreshID := database.InsertToken(body.Email, signedRefreshToken)

	jwtCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshID,
		MaxAge:   RefreshLife,
		Path:     "/api/auth",
		HttpOnly: true,
	}
	c.SetCookie(
		jwtCookie.Name,
		jwtCookie.Value,
		jwtCookie.MaxAge,
		jwtCookie.Path,
		jwtCookie.Domain,
		jwtCookie.Secure,
		jwtCookie.HttpOnly,
	)

	c.JSON(http.StatusOK, gin.H{
		"message":      "Authentication was successful",
		"access_token": signedAccessToken,
	})
}
