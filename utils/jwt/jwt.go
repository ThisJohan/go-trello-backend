package jwt

import (
	"errors"
	"fmt"
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

func parse(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.TOKEN_KEY), nil
	})
}

func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)

	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	id, ok := claims["ID"].(float64)
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &TokenPayload{
		ID: uint(id),
	}, nil
}
