package route

import (
	"boilerplate/internal/config"
	"boilerplate/internal/service"
	"boilerplate/pkg/logger"
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
