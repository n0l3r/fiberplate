package middleware

import (
	"github.com/n0l3r/fiberplate/pkg/logger"
	"fmt"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

// RecoverMiddleware will catch panic, log it, and return 500
func RecoverMiddleware(logger *logger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {

				// // Log panic and stacktrace with request logger if available
				if reqLogger, ok := c.Locals("logger").(zerolog.Logger); ok {
					reqLogger.Error().
						Str("panic", fmt.Sprintf("%v", r)).
						Bytes("stacktrace", debug.Stack()).
						Msg("panic recovered")
				} else {
					logger.Error().
						Str("panic", fmt.Sprintf("%v", r)).
						Bytes("stacktrace", debug.Stack()).
						Msg("panic recovered")
				}

				// Set response
				_ = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":    fiber.StatusInternalServerError,
					"data":    nil,
					"message": "internal server error",
				})
			}
		}()
		return c.Next()
	}
}
