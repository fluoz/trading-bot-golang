package controllers

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"
	"trading-bot/internal/consts"
	"trading-bot/internal/presentations"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app = setupTestHealthController()

func setupTestHealthController() *fiber.App {
	app := fiber.New()
	healthController := NewHealthController()
	healthController.Route(app)
	return app
}

func TestPing(t *testing.T) {
	req := httptest.NewRequest("GET", "/ping", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	result := presentations.Response{}
	json.Unmarshal(body, &result)
	assert.Equal(t, fiber.StatusOK, result.Code)
	assert.Equal(t, consts.RespStatusOK, result.Status)
}
