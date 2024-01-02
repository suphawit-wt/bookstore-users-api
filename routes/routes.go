package routes

import (
	"bookstore/users/handler"
	"bookstore/users/repository"
	"bookstore/users/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(db *sqlx.DB, router *gin.Engine) {
	pingHandler := handler.NewPingHandler()

	userRepositoryDB := repository.NewUserRepositoryDB(db)
	userService := services.NewUserService(userRepositoryDB)
	userHandler := handler.NewUserHandler(userService)

	api := router.Group("api")
	api.GET("/ping", pingHandler.GetPing)

	api.GET("/users/:id", userHandler.GetUser)
	api.GET("/search/users", userHandler.SearchUsers)
	api.POST("/users", userHandler.CreateUser)
	api.PUT("/users/:id", userHandler.UpdateUser)
	api.DELETE("/users/:id", userHandler.DeleteUser)
}
