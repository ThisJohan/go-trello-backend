package services

import (
	"github.com/ThisJohan/go-trello-clone/app/dal"
	"github.com/ThisJohan/go-trello-clone/app/types"
	"github.com/ThisJohan/go-trello-clone/utils"
	"github.com/gofiber/fiber/v2"
)

func CreteBoard(ctx *fiber.Ctx) error {
	body := new(types.BoardDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	board := dal.Board{Title: body.Title, Description: body.Description, UserID: utils.GetUser(ctx)}

	if res := dal.CreateBoard(&board); res.Error != nil {
		return res.Error
	}

	ctx.JSON(&board)

	return nil
}

func GetBoards(ctx *fiber.Ctx) error {

	boards := new([]dal.Board)

	if res := dal.FindBoardsByUser(boards, utils.GetUser(ctx)); res.Error != nil {
		return res.Error
	}

	ctx.JSON(&boards)

	return nil
}

func GetBoardById(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	board := new(dal.Board)

	if res := dal.FindBoardById(board, id); res.Error != nil {
		return res.Error
	}

	ctx.JSON(board)

	return nil

}

func DeleteBoardById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if res := dal.DeleteBoard(id, utils.GetUser(ctx)); res.Error != nil {
		return res.Error
	}

	ctx.SendString("Deleted Successfully")

	return nil
}

func UpdateBoard(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	body := new(types.BoardDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	if res := dal.UpdateBoard(id, utils.GetUser(ctx), body); res.Error != nil {
		return res.Error
	}

	ctx.SendString("Updated Successfully")

	return nil
}
