package main

import (
	"server/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to social app")
	})

	app.Listen(":5000")
}
