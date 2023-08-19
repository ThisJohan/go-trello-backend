package routes

import (
	"github.com/ThisJohan/go-trello-clone/app/services"
	"github.com/ThisJohan/go-trello-clone/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func cardRoutes(app fiber.App) {
	app.Post("/list/:listId/card", middleware.Auth, services.CreateCard)
	app.Get("/list/:listId/card", middleware.Auth, services.GetCards)

	r := app.Group("/list").Use(middleware.Auth)

	r.Get("/:id", services.GetCard)
	r.Put("/:id", services.UpdateCard)
	r.Delete("/:id", services.DeleteCard)
}
