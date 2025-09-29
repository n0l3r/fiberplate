package handler

import (
	"boilerplate/pkg/responder"

	"github.com/gofiber/fiber/v2"
)

// PingHandler responds with "pong" and the current date
func PingHandler(c *fiber.Ctx) error {
	response := responder.NewHTTPSuccess(c, fiber.StatusOK, fiber.Map{
		"date": c.Context().Time().String(),
	}, "pong")
	response.JSON()
	return nil
}
