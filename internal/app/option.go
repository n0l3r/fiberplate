package app

import (
	"github.com/n0l3r/fiberplate/internal/config"
	"github.com/n0l3r/fiberplate/internal/service"
	"github.com/n0l3r/fiberplate/pkg/logger"
)

type Option func(*Router)

func WithLogger(logger *logger.Logger) Option {
	return func(r *Router) {
		r.Logger = logger
	}
}

func WithAppConfig(cfg *config.AppCfg) Option {
	return func(r *Router) {
		r.AppCfg = cfg
	}
}

func WithService(svc *service.Service) Option {
	return func(r *Router) {
		r.Svc = svc
	}
}
