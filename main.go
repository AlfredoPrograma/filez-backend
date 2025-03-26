package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func bootstrap() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello world",
		})
	})

	return app
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	app := bootstrap()

	if err := app.Listen(":9000"); err != nil {
		panic(err)
	}
}
