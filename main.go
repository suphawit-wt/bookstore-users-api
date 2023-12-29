package main

import (
	"bookstore/users/app"

	"github.com/gobuffalo/envy"
)

func main() {
	envy.Load()
	app.StartApplication()
}
