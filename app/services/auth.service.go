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

func Login(ctx *fiber.Ctx) error {

	body := new(types.LoginDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	user := new(types.UserResponse)

	err := dal.FindUserByEmail(user, body.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid Email or Password")
	}

	if err := password.Verify(user.Password, body.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid Email or Password")
	}

	t := jwt.Generate(jwt.TokenPayload{ID: user.ID})

	ctx.JSON(&types.AuthResponse{
		User: user,
		Auth: &types.AccessResponse{Token: t},
	})

	return nil
}

func CheckUser(ctx *fiber.Ctx) error {

	userId := utils.GetUser(ctx)

	user := new(dal.User)

	if err := dal.FindUser(user, "id = ?", userId).Error; err != nil {
		return err
	}

	ctx.JSON(&types.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email})

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

	ctx.JSON(&types.AuthResponse{
		User: &types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: &types.AccessResponse{Token: t},
	})
	return nil
}
