package services

import (
	"github.com/ThisJohan/go-trello-clone/app/dal"
	"github.com/ThisJohan/go-trello-clone/app/types"
	"github.com/ThisJohan/go-trello-clone/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	c.SendString("Login Works!")
	return nil
}

func Signup(c *fiber.Ctx) error {
	body := new(types.SignupDTO)

	if err := c.BodyParser(body); err != nil {
		return err
	}

	err := utils.Validate(body)
	if err != nil {
		return err
	}

	user := dal.User{Email: body.Email, Name: body.Name, Password: body.Password}

	if result := dal.CreateUser(&user); result.Error != nil {
		return result.Error
	}

	c.JSON(user)
	return nil
}
