package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	username string
	password string
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusCreated, "get_user")
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, "create_user")
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "update_user")
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusCreated, "delete_user")
}
