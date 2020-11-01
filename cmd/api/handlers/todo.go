package handlers

import (
	"service/internal"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Index(t internal.Todoable) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		results, err := t.Get()
		if err != nil {
			return err
		}

		return ctx.JSON(results)
	}
}

func Store(t internal.Todoable) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		text := ctx.FormValue("text")
		if text == "" {
			return ctx.Status(422).JSON(fiber.Map{
				"message": "Validation Failed",
				"errors":  fiber.Map{"text": "The [text] field is required"},
			})
		}

		id, err := t.Store(text)
		if err != nil {
			return err
		}

		todo, err := t.Find(id)
		if err != nil {
			return err
		}

		return ctx.JSON(todo)
	}
}

func Complete(t internal.Todoable) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, _ := strconv.Atoi(ctx.Params("id"))

		if err := t.Complete(int64(id)); err != nil {
			return err
		}

		return ctx.SendStatus(fiber.StatusNoContent)
	}
}

func Uncomplete(t internal.Todoable) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, _ := strconv.Atoi(ctx.Params("id"))

		if err := t.Uncomplete(int64(id)); err != nil {
			return err
		}

		return ctx.SendStatus(fiber.StatusNoContent)
	}
}

func Destroy(t internal.Todoable) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, _ := strconv.Atoi(ctx.Params("id"))

		if err := t.Destroy(int64(id)); err != nil {
			return err
		}

		return ctx.SendStatus(fiber.StatusNoContent)
	}
}
