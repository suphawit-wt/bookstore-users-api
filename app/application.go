package app

import (
	"bookstore/users/database"
	"bookstore/users/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	database.InitDatabase()
	setupRoutes()
	logger.Info("Users Microservices starting at port 8080")
	router.Run(":8080")
}
