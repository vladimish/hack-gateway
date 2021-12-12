package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vladimish/hack-gateway/internal/handlers"
	"github.com/vladimish/hack-gateway/internal/k8s"
	"github.com/vladimish/hack-gateway/internal/requests"
	"k8s.io/apimachinery/pkg/util/json"
)

func main() {
	k8s.InitKube()
	StartApi()
}

func StartApi() {
	app := fiber.New()

	app.Post("/add_restaurant", func(c *fiber.Ctx) error {
		req := requests.AddRestaurant{}
		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			return err
		}
		err = handlers.HandleAddRestaurant(req)
		if err != nil {
			return err
		}
		return nil
	})
	app.Listen(":3000")
}
