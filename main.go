package main

import (
	"trading-bot/internal/application"
	"trading-bot/internal/controllers"
	"trading-bot/internal/exchanges"
	binance "trading-bot/internal/exchanges/binance/futures"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Configuration
	app := fiber.New()
	config := application.NewConfig()
	app.Use(logger.New())

	// Exchanges
	binanceFutures := binance.NewClientFutures(config)
	exchanges.NewMarketFutures(binanceFutures)

	// Controller
	connectionController := controllers.NewHealthController()
	connectionController.Route(app)

	app.Listen(":3000")
}
