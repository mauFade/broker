package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/broker/application/handler/user"
)

func setupRoutes(app *fiber.App) {
	app.Post("/v1/users", user.CreateUserHandler)
}
