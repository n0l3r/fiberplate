package config

import "boilerplate/pkg/envy"

// App is the configuration for the application.
type AppCfg struct {
	Name      string
	Environ   string
	Host      string
	Port      string
	Timezone  string
	DebugMode bool
}

// WithAppConfig sets the application configuration.
func WithAppConfig() Option {
	return func(c *Config) {
		c.App = AppCfg{
			Name:      envy.Get(EnvAppName, DefaultAppName),
			Environ:   envy.Get(EnvAppEnviron, DefaultAppEnv),
			Host:      envy.Get(EnvAppHost, DefaultAppHost),
			Port:      envy.Get(EnvAppPort, DefaultAppPort),
			Timezone:  envy.Get(EnvAppTimezone, DefaultAppTimezone),
			DebugMode: envy.GetAsBool(EnvAppDebugMode, DefaultAppDebugMode),
		}
	}
}
