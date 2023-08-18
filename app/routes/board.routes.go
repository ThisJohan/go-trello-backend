package routes

import (
	"github.com/ThisJohan/go-trello-clone/app/services"
	"github.com/ThisJohan/go-trello-clone/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func boardRoutes(app fiber.Router) {
	r := app.Group("/board").Use(middleware.Auth)

	r.Get("/", services.GetBoards)
	r.Get("/:id", services.GetBoardById)
	r.Post("/", services.CreteBoard)
	r.Delete("/:id", services.DeleteBoardById)
	r.Put("/:id", services.UpdateBoard)
}
