package database

import (
	"github.com/n0l3r/fiberplate/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDBConnection(cfg *config.DBCfg) (*sqlx.DB, error) {
	if !isSupportedDriver(cfg.Driver) {
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}

	dsn := newDSN(cfg).toDSN()
	db, err := connectToDB(cfg.Driver, dsn)
	if err != nil {
		return nil, err
	}

	setConnectionPool(db, cfg)
	return db, nil
}

func connectToDB(driver, dsn string) (*sqlx.DB, error) {
	return sqlx.Connect(driver, dsn)
}

func setConnectionPool(db *sqlx.DB, cfg *config.DBCfg) {
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.MaxLifetime)
	db.SetConnMaxIdleTime(cfg.MaxIdleTime)
}
