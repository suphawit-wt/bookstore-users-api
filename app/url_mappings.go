package app

import (
	ping_controller "bookstore/users/controllers/ping"
	users_controller "bookstore/users/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping_controller.Ping)

	router.GET("/users/:id", users_controller.GetUser)
	router.POST("/users", users_controller.CreateUser)
}
