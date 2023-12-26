package users_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "implement me!",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "implement me!",
	})
}
