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

	app.Post("/add_table", func(c *fiber.Ctx) error {
		req := requests.AddTable{}
		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			return err
		}
		err = handlers.HandleAddTable(req)
		if err != nil {
			return err
		}

		return nil
	})

	app.Post("/delete_table", func(c *fiber.Ctx) error {
		req := requests.DeleteTable{}
		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			return err
		}
		err = handlers.HandleDeleteTable(req)
		if err != nil {
			return err
		}

		return nil
	})

	app.Post("/get_tables", func(c *fiber.Ctx) error {
		req := requests.GetTables{}
		err := json.Unmarshal(c.Body(), &req)
		if err != nil {
			return err
		}
		body, err := handlers.HandleGetTables(req)
		if err != nil {
			return err
		}
		_, err = c.Write(body)
		if err != nil {
			return err
		}
		return nil
	})

	app.Listen(":3000")
}
