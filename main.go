package main

import (
	"fmt"

	"github.com/ThisJohan/go-trello-clone/app/dal"
	"github.com/ThisJohan/go-trello-clone/app/routes"
	"github.com/ThisJohan/go-trello-clone/config"
	"github.com/ThisJohan/go-trello-clone/config/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadEnv()
	database.Connect()
	database.Migrate(dal.User{}, &dal.Board{}, &dal.List{}, &dal.Card{})

	app := fiber.New()

	app.Use(cors.New())

	r := app.Group("/api")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.SetupRoutes(r)

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}
