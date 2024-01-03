package handler

import (
	"bookstore/users/errs"
	"bookstore/users/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handleError(c, errs.NewBadRequestError("User ID must be number."))
		return
	}

	user, err := h.userSrv.GetUser(userId)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h userHandler) SearchUsers(c *gin.Context) {
	status := c.Query("status")

	users, err := h.userSrv.SearchUsers(status)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h userHandler) CreateUser(c *gin.Context) {
	req := services.CreateUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, errs.NewBadRequestError("Invalid JSON Body."))
		return
	}

	err := h.userSrv.CreateUser(req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created User Successfully!",
	})
}

func (h userHandler) UpdateUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handleError(c, errs.NewBadRequestError("User ID must be number."))
		return
	}

	req := services.UpdateUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleError(c, errs.NewBadRequestError("Invalid JSON Body."))
		return
	}

	err = h.userSrv.UpdateUser(userId, req)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated User Successfully!",
	})
}

func (h userHandler) DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		handleError(c, errs.NewBadRequestError("User ID must be number."))
		return
	}

	err = h.userSrv.DeleteUser(userId)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted User Successfully!",
	})
}
