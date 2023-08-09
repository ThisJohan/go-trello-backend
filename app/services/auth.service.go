package services

import (
	"github.com/ThisJohan/go-trello-clone/app/dal"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	c.SendString("Login Works!")
	return nil
}

func Signup(c *fiber.Ctx) error {

	result := dal.CreateUser(&dal.User{Email: "khosravijohan@gmail.com", Name: "Johan", Password: "12345"})

	if result.Error != nil {
		c.SendStatus(400)
	}

	c.SendString("Signup Works!")
	return nil
}
