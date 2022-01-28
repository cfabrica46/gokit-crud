package middleware

import (
	"net/http"

	"github.com/cfabrica46/go-crud/database/cache"
	"github.com/cfabrica46/go-crud/structure"
	"github.com/cfabrica46/go-crud/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetUserFromBody(c *gin.Context) {
	var user structure.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, structure.ResponseHTTP{Code: http.StatusBadRequest, ErrorText: "Invalid Login"})
		return
	}

	c.Set("username", user.Username)
	c.Set("password", user.Password)
	c.Set("email", user.Email)
	c.Next()
}

func GetUserFromToken(c *gin.Context) {
	var tokenStructure structure.Token

	err := c.BindHeader(&tokenStructure)
	if err != nil {
		c.JSON(http.StatusBadRequest, structure.ResponseHTTP{Code: http.StatusBadRequest, ErrorText: "Invalid Header"})
		return
	}

	check, err := cache.TokenIsValid(tokenStructure.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structure.ResponseHTTP{Code: http.StatusInternalServerError, ErrorText: "Error Verifying Token"})
		return
	}
	if !check {
		c.JSON(http.StatusUnauthorized, structure.ResponseHTTP{Code: http.StatusUnauthorized, ErrorText: "Token Is Blacklisted"})
		return
	}

	id, username, email, err := token.ExtractClaims(tokenStructure.Token, "key.pem", jwt.SigningMethodHS256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, structure.ResponseHTTP{Code: http.StatusInternalServerError, ErrorText: "Error Verifying Token"})
		return
	}

	c.Set("id", id)
	c.Set("username", username)
	c.Set("email", email)
	c.Set("token", tokenStructure.Token)
	c.Next()
}
