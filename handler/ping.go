package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pingHandler struct{}

func NewPingHandler() pingHandler {
	return pingHandler{}
}

func (h pingHandler) GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
