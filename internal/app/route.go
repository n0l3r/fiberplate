package app

import (
	"github.com/n0l3r/fiberplate/internal/config"
	"github.com/n0l3r/fiberplate/internal/middleware"
	"github.com/n0l3r/fiberplate/internal/service"
	"github.com/n0l3r/fiberplate/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
)

type Router struct {
	AppCfg *config.AppCfg
	Svc    *service.Service
	Logger *logger.Logger
}

func NewRouter(opts ...Option) *Router {
	r := &Router{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *Router) Setup() *fiber.App {
	app := fiber.New()

	app.Use(middleware.RequestIDMiddleware())
	app.Use(middleware.RecoverMiddleware(r.Logger))
	app.Use(helmet.New())
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(logger.FiberMiddleware(r.Logger))

	registerDevRoutes(r, app)

	// Setup routes
	v1 := app.Group("/api/v1")
	_ = v1

	return app
}
