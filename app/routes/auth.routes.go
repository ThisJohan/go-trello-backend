package routes

import (
	"github.com/ThisJohan/go-trello-clone/app/services"
	"github.com/ThisJohan/go-trello-clone/utils/middleware"
	"github.com/gofiber/fiber/v2"
)

func authRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", services.Signup)
	r.Post("/login", services.Login)
	r.Get("/", middleware.Auth, services.CheckUser)
}
