package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/broker/database"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":8081")
}
