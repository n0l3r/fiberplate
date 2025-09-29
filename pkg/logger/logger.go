package logger

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/n0l3r/fiberplate/pkg/envy"
	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

// NewLogger creates a logger instance based on the environment
func NewLogger() *Logger {
	if envy.Get("APP_ENV", "development") == "production" {
		return createLogger(os.Stdout, zerolog.DebugLevel, false)
	}
	return createLogger(os.Stdout, zerolog.InfoLevel, true)
}

// createLogger is a helper to configure the logger
func createLogger(output *os.File, level zerolog.Level, humanReadable bool) *Logger {
	var logger zerolog.Logger
	if humanReadable {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: output, TimeFormat: time.RFC3339}).With().Timestamp().Caller().Logger()
	} else {
		logger = zerolog.New(output).With().Timestamp().Caller().Logger()
	}
	logger = logger.Level(level)
	return &Logger{&logger}
}

// Extract logger from context
func LoggerFromFiber(c fiber.Ctx) zerolog.Logger {
	l, ok := c.Locals(LoggerKey).(zerolog.Logger)
	if ok {
		return l
	}
	return zerolog.Nop()
}

// FiberMiddleware to add logger to context and log requests
func FiberMiddleware(baseLogger *Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		reqID := c.Get(fiber.HeaderXRequestID)
		method := c.Method()
		path := c.Path()
		userAgent := c.Get(fiber.HeaderUserAgent)
		ip := c.IP()
		startTime := time.Now()

		// Logger with fields
		reqLogger := baseLogger.With().
			Str("req_id", reqID).
			Str("method", method).
			Str("path", path).
			Str("user_agent", userAgent).
			Str("ip", ip).
			Time("start_time", startTime).
			Logger()

		c.Locals(LoggerKey, reqLogger)

		err := c.Next()

		endTime := time.Now()
		duration := endTime.Sub(startTime)
		status := c.Response().StatusCode()
		bodySize := len(c.Response().Body())

		// Log access event
		reqLogger.Info().
			Time("end_time", endTime).
			Dur("duration", duration).
			Int("status", status).
			Int("body_size", bodySize).
			Msg("access")

		return err
	}
}
