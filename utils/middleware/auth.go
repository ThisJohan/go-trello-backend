package middleware

import (
	"strings"

	"github.com/ThisJohan/go-trello-clone/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {

	h := ctx.Get("Authorization")

	if h == "" {
		return fiber.ErrUnauthorized
	}

	chunks := strings.Split(h, " ")

	if len(chunks) < 2 {
		return fiber.ErrUnauthorized
	}

	user, err := jwt.Verify(chunks[1])

	if err != nil {
		return fiber.ErrUnauthorized
	}

	ctx.Locals("USER", user.ID)

	return ctx.Next()
}
