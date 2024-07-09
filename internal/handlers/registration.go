package handlers

import (
	"dev/reglogauth/internal/database"
	"dev/reglogauth/internal/http_responses"
	"dev/reglogauth/internal/models"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /

// Registration godoc
// @Summary registers a user
// @Schemes application/json
// @Description accepts json sent by the user as input and registers it
// @Tags registration
// @Accept json
// @Produce json
// @Param input body models.RegisterRequest true "account info"
// @Success 200 "message: Registration was successful"
// @Failure 400 "error: Failed to read body"
// @Failure 409 "fail: This email already exists"
// @Failure 422 "fail: Your email is not valid"
// @Failure 422 "fail: Your username is not valid"
// @Failure 422 "fail: Your password is not valid"
// @Failure 500 "error: Error on the server. Please, try again later"
// @Router /api/auth/registration [post]
func Registration(c *gin.Context) {
	body := models.RegisterRequest{}
	if c.Bind(&body) != nil {
		http_responses.FailToReadBody(c)
		return
	}

	if len(body.Email) > 255 || !isEmailValid(body.Email) {
		http_responses.FailIncorrectEmail(c)
		return
	}

	if len(body.UserName) > 255 || len(body.UserName) < 4 {
		http_responses.FailIncorrectUserName(c)
		return
	}

	if len(body.Password) > 72 || len(body.Password) < 8 {
		http_responses.FailIncorrectPassword(c)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		http_responses.ErrorOnServer(c)
		return
	}

	user := models.User{
		Email:            body.Email,
		UserName:         body.UserName,
		Password:         string(hash),
		RegistrationTime: time.Now(),
	}

	err = database.InsertUser(user)
	if err != nil {
		http_responses.FailCurrentEmailAlreadyExists(c)
		return
	}

	http_responses.ExecRegister(c)
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
