package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// RequestIDMiddleware sets request ID if missing
func RequestIDMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		reqID := c.Get(fiber.HeaderXRequestID)
		if reqID == "" {
			reqID = uuid.NewString()
			c.Request().Header.Set(fiber.HeaderXRequestID, reqID)
		}

		c.Locals(fiber.HeaderXRequestID, reqID)
		return c.Next()
	}
}
