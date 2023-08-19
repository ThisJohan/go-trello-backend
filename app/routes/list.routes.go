package routes

import (
	"github.com/ThisJohan/go-trello-clone/app/services"
	"github.com/ThisJohan/go-trello-clone/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func listRoutes(app fiber.Router) {

	app.Post("/board/:boardId/list", middleware.Auth, services.CreateList)
	app.Get("/board/:boardId/list", middleware.Auth, services.GetLists)

	r := app.Group("/list").Use(middleware.Auth)

	r.Get("/:id", services.GetList)
	r.Put("/:id", services.UpdateList)
	r.Delete("/:id", services.DeleteList)
}
