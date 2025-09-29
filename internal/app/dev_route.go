package app

import (
	"github.com/n0l3r/fiberplate/internal/handler"

	"github.com/gofiber/fiber/v2"
)

// registerDevRoutes registers development routes.
func registerDevRoutes(r *Router, app *fiber.App) {
	app.Get("/ping", handler.PingHandler)
}
