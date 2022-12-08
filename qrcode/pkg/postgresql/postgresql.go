package postgresql

import (
	"database/sql"

	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewPgDb(c *config.DbConfig, pc *config.PoolConfig) (*bun.DB, error) {
	var connector = pgdriver.NewConnector(pgdriver.WithDSN(c.Dsn()))
	var db = sql.OpenDB(connector)
	db.SetMaxOpenConns(pc.MaxConn)
	db.SetMaxIdleConns(pc.MinConn)
	db.SetConnMaxIdleTime(pc.MaxConnIdleTime)
	db.SetConnMaxLifetime(pc.MaxConnLifetime)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return bun.NewDB(db, pgdialect.New()), nil
}
