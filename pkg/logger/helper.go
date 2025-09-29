package logger

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
)

const LoggerKey = "logger"

// GetFromFiberContext retrieves the logger from the context
func GetFromFiberContext(c fiber.Ctx) zerolog.Logger {
	if logger, ok := c.Locals(LoggerKey).(zerolog.Logger); ok {
		return logger
	}
	return zerolog.Nop()
}
