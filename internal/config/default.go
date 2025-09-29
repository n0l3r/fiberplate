package config

var (
	DefaultAppName      = "Go Web API"
	DefaultAppEnv       = "local"
	DefaultAppHost      = "localhost"
	DefaultAppPort      = "8080"
	DefaultAppTimezone  = "UTC"
	DefaultAppDebugMode = true

	DefaultDBDriver       = "postgres"
	DefaultDBHost         = "localhost"
	DefaultDBPort         = "5432"
	DefaultDBUser         = "postgres"
	DefaultDBPassword     = "password"
	DefaultDBName         = "postgres"
	DefaultDBSSLMode      = "disable"
	DefaultDBTimezone     = "UTC"
	DefaultDBMaxIdleConns = 10
	DefaultDBMaxOpenConns = 100
	DefaultDBMaxLifetime  = "1h"
	DefaultDBMaxIdleTime  = "30m"
	DefaultDBDialTimeout  = "5s"
	DefaultDBReadTimeout  = "5s"
	DefaultDBWriteTimeout = "5s"
)
