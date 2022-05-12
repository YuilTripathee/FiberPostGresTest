package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// Default middleware config
	app.Use(logger.New())

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// GET http://localhost:8080/hello%20world

	app.Get("/say/:value", func(c *fiber.Ctx) error {
		return c.SendString("You said: " + c.Params("value"))
		// => Get request with value: hello world
	})

	app.Static("/", "./public")

	app.Listen(":3000")

}
