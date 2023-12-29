package users_controller

import (
	models "bookstore/users/models/users"
	"bookstore/users/services"
	"bookstore/users/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("User ID must be number.")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	user, restErr := services.UsersService.GetById(userId)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, restErr := services.UsersService.Search(status)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	req := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	restErr := services.UsersService.Create(req)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created User Successfully!",
	})
}

func UpdateUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("User ID must be number.")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	req := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	restErr := services.UsersService.Update(userId, req)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated User Successfully!",
	})
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("User ID must be number.")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	restErr := services.UsersService.Delete(userId)
	if restErr != nil {
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted User Successfully!",
	})
}
