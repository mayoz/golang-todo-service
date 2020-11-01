package handlers

import "github.com/gofiber/fiber/v2"

func Ping() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"ping": "pong",
		})
	}
}
