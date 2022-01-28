package handler

import (
	"net/http"

	"github.com/cfabrica46/go-crud/database/userdb"
	"github.com/cfabrica46/go-crud/structure"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	id := c.MustGet("id").(int)
	username := c.MustGet("username").(string)
	email := c.MustGet("email").(string)
	c.JSON(http.StatusOK, structure.ResponseHTTP{Code: http.StatusOK, Content: structure.User{ID: id, Username: username, Email: email}})
}

func GetAllUsers(c *gin.Context) {
	users, err := userdb.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, structure.ResponseHTTP{Code: http.StatusInternalServerError, ErrorText: "Error to get all users"})
		return
	}

	c.JSON(http.StatusOK, structure.ResponseHTTP{Code: http.StatusOK, Content: users})
}

func DeleteUser(c *gin.Context) {
	id := c.MustGet("id").(int)

	count, err := userdb.DeleteUserbByID(id)
	if err != nil {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to delete user"})
		return
	}
	if count == 0 {
		c.JSON(http.StatusConflict, structure.ResponseHTTP{Code: http.StatusConflict, ErrorText: "Conflict to delete user"})
		return
	}

	c.JSON(http.StatusOK, structure.ResponseHTTP{Code: http.StatusOK})
}
