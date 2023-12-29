package app

import (
	ping_controller "bookstore/users/controllers/ping"
	users_controller "bookstore/users/controllers/users"
)

func setupRoutes() {
	api := router.Group("api")
	api.GET("/ping", ping_controller.Ping)

	api.GET("/users/:id", users_controller.GetUserById)
	api.GET("/search/users", users_controller.Search)
	api.POST("/users", users_controller.CreateUser)
	api.PUT("/users/:id", users_controller.UpdateUser)
	api.DELETE("/users/:id", users_controller.DeleteUser)
}
