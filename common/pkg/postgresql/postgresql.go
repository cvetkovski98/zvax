package postgresql

import (
	"database/sql"
	"log"

	"github.com/cvetkovski98/zvax-common/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewPgDb(cfg *config.PostgreSQL) (*bun.DB, error) {
	log.Println("Connecting to PostgreSQL database...")
	var connector = pgdriver.NewConnector(pgdriver.WithDSN(cfg.Dsn()))
	var db = sql.OpenDB(connector)
	db.SetMaxOpenConns(cfg.Pool.MaxConn)
	db.SetMaxIdleConns(cfg.Pool.MinConn)
	db.SetConnMaxIdleTime(cfg.Pool.MaxConnIdleTime)
	db.SetConnMaxLifetime(cfg.Pool.MaxConnLifetime)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return bun.NewDB(db, pgdialect.New()), nil
}
