package config

const (
	// App config Environment Variables
	EnvAppName      = "APP_NAME"
	EnvAppEnviron   = "APP_ENV"
	EnvAppHost      = "APP_HOST"
	EnvAppPort      = "APP_PORT"
	EnvAppTimezone  = "APP_TIMEZONE"
	EnvAppDebugMode = "APP_DEBUG_MODE"

	// Database config Environment Variables
	EnvDBDriver       = "DB_DRIVER"
	EnvDBHost         = "DB_HOST"
	EnvDBPort         = "DB_PORT"
	EnvDBUser         = "DB_USER"
	EnvDBPassword     = "DB_PASSWORD"
	EnvDBName         = "DB_NAME"
	EnvDBSSLMode      = "DB_SSLMODE"
	EnvDBTimezone     = "DB_TIMEZONE"
	EnvDBMaxIdleConns = "DB_MAX_IDLE_CONNS"
	EnvDBMaxOpenConns = "DB_MAX_OPEN_CONNS"
	EnvDBMaxLifetime  = "DB_CONN_MAX_LIFETIME"
	EnvDBMaxIdleTime  = "DB_CONN_MAX_IDLE_TIME"
	EnvDBDialTimeout  = "DB_DIAL_TIMEOUT"
	EnvDBReadTimeout  = "DB_READ_TIMEOUT"
	EnvDBWriteTimeout = "DB_WRITE_TIMEOUT"
)
