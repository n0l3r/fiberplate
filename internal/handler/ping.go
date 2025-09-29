package handler

import (
	"github.com/n0l3r/fiberplate/pkg/responder"

	"github.com/gofiber/fiber/v3"
)

// PingHandler responds with "pong" and the current date
func PingHandler(c fiber.Ctx) error {
	response := responder.NewHTTPSuccess(c, fiber.StatusOK, fiber.Map{
		"date": c.RequestCtx().Time().String(),
	}, "pong")
	response.JSON()
	return nil
}
