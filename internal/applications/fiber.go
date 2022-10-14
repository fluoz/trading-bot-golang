package applications

import "github.com/gofiber/fiber/v2"

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		AppName: "botchill",
	}
}
