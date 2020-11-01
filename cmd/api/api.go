package main

import (
	"service/pkg/validator"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (server *Server) listItems(ctx *fiber.Ctx) error {
	results, err := server.store.Get()
	if err != nil {
		return err
	}

	return ctx.JSON(results)
}

type storeItemRequest struct {
	Text string `json:"text" validate:"required,min=3,max=255"`
}

func (server *Server) storeItem(ctx *fiber.Ctx) error {
	var req storeItemRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	if errors := validator.ValidateStruct(req); len(errors) > 0 {
		return ctx.Status(422).JSON(errors)
	}

	id, err := server.store.Store(req.Text)
	if err != nil {
		return err
	}

	todo, err := server.store.Find(id)
	if err != nil {
		return err
	}

	return ctx.JSON(todo)
}

func (server *Server) toggleItem(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := server.store.Toggle(int64(id)); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (server *Server) destroyItem(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := server.store.Destroy(int64(id)); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
