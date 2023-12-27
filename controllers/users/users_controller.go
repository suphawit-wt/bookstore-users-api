package users_controller

import (
	models "bookstore/users/models/users"
	"bookstore/users/services"
	"bookstore/users/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	req := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	user, err := services.CreateUser(req)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("User ID must be number.")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	user, restErr := services.GetUserById(userId)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
