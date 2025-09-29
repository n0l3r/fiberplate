package route

import (
	"boilerplate/internal/handler"

	"github.com/gofiber/fiber/v2"
)

// registerDevRoutes registers development routes.
func registerDevRoutes(r *Router, app *fiber.App) {
	app.Get("/ping", handler.PingHandler)
}
