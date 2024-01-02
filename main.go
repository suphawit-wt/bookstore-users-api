package main

import (
	"bookstore/users/database"
	"bookstore/users/logger"
	"bookstore/users/routes"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
)

func main() {
	envy.Load()
	db := database.InitDatabase()

	router := gin.Default()
	routes.SetupRoutes(db, router)
	logger.Info("Users Microservices starting at port 8080")
	router.Run(":8080")
}
