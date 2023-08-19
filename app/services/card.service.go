package services

import (
	"github.com/ThisJohan/go-trello-clone/app/dal"
	"github.com/ThisJohan/go-trello-clone/app/types"
	"github.com/ThisJohan/go-trello-clone/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateCard(ctx *fiber.Ctx) error {
	body := new(types.CardDTO)
	listId := utils.StringToUint(ctx.Params("listId"))

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	card := dal.Card{Title: body.Title, Description: body.Description, ListId: listId, AssignedTo: body.AssignedTo, DueDate: body.DueDate}

	dal.CreateCard(&card)

	ctx.JSON(card)

	return nil
}

func GetCards(ctx *fiber.Ctx) error {
	listId := ctx.Params("listId")

	cards := new([]dal.Card)

	if err := dal.FindCardsByList(cards, listId).Error; err != nil {
		return err
	}

	ctx.JSON(cards)

	return nil
}

func GetCard(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	Card := new(dal.Card)

	if err := dal.FindCardById(Card, id).Error; err != nil {
		return err
	}

	ctx.JSON(Card)

	return nil
}

func UpdateCard(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	body := new(types.CardDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	if err := dal.UpdateCard(id, body).Error; err != nil {
		return err
	}

	ctx.SendString("Updated Successfully")

	return nil
}

func DeleteCard(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if err := dal.DeleteCard(id).Error; err != nil {
		return err
	}

	ctx.SendString("Deleted Successfully")

	return nil
}

func AssignCardToUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	body := new(types.AssignCardToUserDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	if err := dal.UpdateCard(id, dal.Card{AssignedTo: body.UserId}).Error; err != nil {
		return err
	}

	return nil
}
