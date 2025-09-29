package database

import (
	"boilerplate/internal/config"
	"boilerplate/pkg/util"
)

type Driver string

const (
	MySQL      Driver = "mysql"
	PostgreSQL Driver = "postgres"
)

var supportedDrivers = map[Driver]struct{}{
	MySQL:      {},
	PostgreSQL: {},
}

type dsnOptions struct {
	driver, host, port, user, password, dbName, sslMode, timeZone string
}

func newDSN(cfg *config.DBCfg) dsnOptions {
	return dsnOptions{
		driver:   cfg.Driver,
		host:     cfg.Host,
		port:     cfg.Port,
		user:     cfg.User,
		password: cfg.Password,
		dbName:   cfg.Name,
		sslMode:  cfg.SSLMode,
		timeZone: cfg.TimeZone,
	}
}

func (o dsnOptions) toDSN() string {
	if Driver(o.driver) == MySQL {
		dsn := o.user + ":" + o.password + "@tcp(" + o.host + ":" + o.port + ")/" + o.dbName
		dsn += "?charset=utf8mb4&parseTime=True&loc=Local"
		if o.timeZone != "" {
			dsn += "&time_zone=" + o.timeZone
		}
		return dsn
	}
	if Driver(o.driver) == PostgreSQL {
		dsn := "postgres://" + o.user + ":" + o.password + "@" + o.host + ":" + o.port + "/" + o.dbName
		dsn += "?sslmode=" + util.DefaultIfEmpty(o.sslMode, "disable")
		if o.timeZone != "" {
			dsn += "&TimeZone=" + o.timeZone
		}
		return dsn
	}
	return ""
}

func isSupportedDriver(driver string) bool {
	_, ok := supportedDrivers[Driver(driver)]
	return ok
}
