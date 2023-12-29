package app

import (
	"bookstore/users/database"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	database.InitDatabase()
	setupRoutes()
	router.Run(":8080")
}
