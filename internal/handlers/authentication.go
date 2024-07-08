package handlers

import (
	"dev/reglogauth/internal/config"
	"dev/reglogauth/internal/database"
	"dev/reglogauth/internal/http_responses"
	"dev/reglogauth/internal/models"
	"fmt"
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

func Authentication(c *gin.Context) {
	body := models.AuthenticationRequest{}
	if c.Bind(&body) != nil {
		http_responses.FailToReadBody(c)
		return
	}
	fmt.Println(config.CFG.JwtSecretKey)
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
		"AccessToken":  signedAccessToken,
		"RefreshToken": signedRefreshToken,
	})
}
