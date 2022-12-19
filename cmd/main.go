package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/pkg/routers"
)

func main() {
	app := fiber.New()
	router := routers.NewRouter()
	app.Post("/books", router.AddBookRouter)

	app.Listen(":3000")
}
