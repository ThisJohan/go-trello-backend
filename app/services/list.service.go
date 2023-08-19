package services

import (
	"github.com/ThisJohan/go-trello-clone/app/dal"
	"github.com/ThisJohan/go-trello-clone/app/types"
	"github.com/ThisJohan/go-trello-clone/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateList(ctx *fiber.Ctx) error {

	body := new(types.ListDTO)
	boardId := utils.StringToUint(ctx.Params("boardId"))

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	list := dal.List{Title: body.Title, Description: body.Description, BoardId: boardId}

	dal.CreateList(&list)

	ctx.JSON(list)

	return nil
}

func GetLists(ctx *fiber.Ctx) error {
	boardId := ctx.Params("boardId")

	lists := new([]dal.List)

	if err := dal.FindListsByBoard(lists, boardId).Error; err != nil {
		return err
	}

	ctx.JSON(lists)

	return nil
}

func GetList(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	list := new(dal.List)

	if err := dal.FindListById(list, id).Error; err != nil {
		return err
	}

	ctx.JSON(list)

	return nil
}

func UpdateList(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	body := new(types.ListDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	if err := dal.UpdateList(id, body).Error; err != nil {
		return err
	}

	ctx.SendString("Updated Successfully")

	return nil
}

func DeleteList(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	if err := dal.DeleteList(id).Error; err != nil {
		return err
	}

	ctx.SendString("Deleted Successfully")

	return nil
}
