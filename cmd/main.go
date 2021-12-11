package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vladimish/hack-gateway/internal/k8s"
)

func main() {
	k8s.InitKube()
	StartApi()
}

func StartApi() {
	app := fiber.New()

	app.Post("/create_restaurant", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
