package config

import (
	"boilerplate/pkg/envy"
	"time"
)

// DBCfg holds the database configuration details.
type DBCfg struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string

	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
	MaxIdleTime  time.Duration
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// WithDBConfig sets the database configuration.
func WithDBConfig() Option {
	return func(c *Config) {
		c.DB = DBCfg{
			Driver:   envy.Get(EnvDBDriver, DefaultDBDriver),
			Host:     envy.Get(EnvDBHost, DefaultDBHost),
			Port:     envy.Get(EnvDBPort, DefaultDBPort),
			User:     envy.Get(EnvDBUser, DefaultDBUser),
			Password: envy.Get(EnvDBPassword, DefaultDBPassword),
			Name:     envy.Get(EnvDBName, DefaultDBName),
			SSLMode:  envy.Get(EnvDBSSLMode, DefaultDBSSLMode),
			TimeZone: envy.Get(EnvDBTimezone, DefaultDBTimezone),
			// Optional settings with sensible defaults
			MaxIdleConns: envy.GetAsInt(EnvDBMaxIdleConns, DefaultDBMaxIdleConns),
			MaxOpenConns: envy.GetAsInt(EnvDBMaxOpenConns, DefaultDBMaxOpenConns),
			MaxLifetime:  envy.GetAsDuration(EnvDBMaxLifetime, DefaultDBMaxLifetime),
			MaxIdleTime:  envy.GetAsDuration(EnvDBMaxIdleTime, DefaultDBMaxIdleTime),
			DialTimeout:  envy.GetAsDuration(EnvDBDialTimeout, DefaultDBDialTimeout),
			ReadTimeout:  envy.GetAsDuration(EnvDBReadTimeout, DefaultDBReadTimeout),
			WriteTimeout: envy.GetAsDuration(EnvDBWriteTimeout, DefaultDBWriteTimeout),
		}
	}
}
