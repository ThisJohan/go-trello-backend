package routes

import (
	"github.com/ThisJohan/go-trello-clone/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func boardRoutes(app fiber.Router) {
	r := app.Group("/board").Use(middleware.Auth)

	r.Get("/", func(c *fiber.Ctx) error {
		c.SendString("Hi There")
		return nil
	})
}
