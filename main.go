package main

import (
	"trading-bot/internal/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	connectionController := controllers.NewHealthController()
	connectionController.Route(app)

	app.Listen(":3000")
}
