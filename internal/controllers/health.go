package controllers

import (
	"trading-bot/internal/consts"
	"trading-bot/internal/presentations"

	"github.com/gofiber/fiber/v2"
)

type healthController struct {
}

func NewHealthController() healthController {
	return healthController{}
}

func (ctrl *healthController) Route(app *fiber.App) {
	app.Get("/ping", ctrl.Ping)
}

func (ctrl *healthController) Ping(c *fiber.Ctx) error {
	res := presentations.Response{
		Code:   fiber.StatusOK,
		Status: consts.RespStatusOK,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
