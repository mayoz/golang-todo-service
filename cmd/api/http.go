package main

import (
	"service/internal/todo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// var cRet func(ctx *fiber.Ctx) error
type Server struct {
	app   *fiber.App
	store *todo.Store
}

func setupServer(store *todo.Store) *Server {
	server := &Server{store: store}
	app := fiber.New()

	// bind middleware
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	// register routes
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"ping": "pong"})
	})

	v1 := app.Group("/v1")
	{
		v1.Get("/todos", server.listItems)
		v1.Post("/todos", server.storeItem)
		v1.Post("/todos/:id/toggle", server.toggleItem)
		v1.Delete("/todos/:id", server.destroyItem)
	}

	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(404)
	})

	server.app = app
	return server
}

func (server *Server) Start(address string) error {
	return server.app.Listen(address)
}
