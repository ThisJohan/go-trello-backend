package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app fiber.Router) {
	authRoutes(app)
	boardRoutes(app)
	listRoutes(app)
}
