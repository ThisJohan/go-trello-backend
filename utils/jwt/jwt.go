package jwt

import (
	"time"

	"github.com/ThisJohan/go-trello-clone/config"
	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	ID uint
}

func Generate(payload TokenPayload) string {
	v, err := time.ParseDuration(config.TOKEN_EXP)

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(v).Unix(),
		"ID":  payload.ID,
	})

	tokenString, err := token.SignedString([]byte(config.TOKEN_KEY))

	if err != nil {
		panic(err)
	}

	return tokenString
}
