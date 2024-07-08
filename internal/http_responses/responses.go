package http_responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FailToReadBody(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"fail": "Failed to read body",
	})
}

func FailIncorrectEmail(c *gin.Context) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"fail": "Your email is not valid",
	})
}

func FailIncorrectUserName(c *gin.Context) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"fail": "Your username is not valid",
	})
}

func FailIncorrectPassword(c *gin.Context) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"fail": "Your password is not valid",
	})
}

func FailCurrentEmailAlreadyExists(c *gin.Context) {
	c.JSON(http.StatusConflict, gin.H{
		"fail": "This email already exists",
	})
}

func FailWrongPassword(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"fail": "You entered the wrong password or email",
	})
}

func ExecRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration was successful",
	})
}

func ErrorOnServer(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Error on the server. Please, try again later",
	})
}
