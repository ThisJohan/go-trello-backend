package services

import (
	"errors"

	"github.com/ThisJohan/go-trello-clone/app/dal"
	"github.com/ThisJohan/go-trello-clone/app/types"
	"github.com/ThisJohan/go-trello-clone/utils"
	"github.com/ThisJohan/go-trello-clone/utils/jwt"
	"github.com/ThisJohan/go-trello-clone/utils/password"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	c.SendString("Login Works!")
	return nil
}

func Signup(ctx *fiber.Ctx) error {
	body := new(types.SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	err := dal.FindUserByEmail(&struct{ ID string }{}, body.Email).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	user := dal.User{Email: body.Email, Name: body.Name, Password: password.Generate(body.Password)}

	if result := dal.CreateUser(&user); result.Error != nil {
		return result.Error
	}

	t := jwt.Generate(jwt.TokenPayload{ID: user.ID})

	ctx.JSON(types.AuthResponse{
		User: types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: types.AccessResponse{Token: t},
	})
	return nil
}
