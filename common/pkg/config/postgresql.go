package config

import (
	"fmt"
	"time"
)

type PostgreSQLPool struct {
	MaxConn           int           `mapstructure:"max_conn"`
	MinConn           int           `mapstructure:"min_conn"`
	MaxConnIdleTime   time.Duration `mapstructure:"max_conn_idle_time"`
	MaxConnLifetime   time.Duration `mapstructure:"max_conn_lifetime"`
	HealthCheckPeriod time.Duration `mapstructure:"health_check_period"`
}

type PostgreSQL struct {
	Host     string         `mapstructure:"host"`
	Port     int            `mapstructure:"port"`
	User     string         `mapstructure:"user"`
	Password string         `mapstructure:"pass"`
	Database string         `mapstructure:"name"`
	Pool     PostgreSQLPool `mapstructure:"pool"`
}

func (c *PostgreSQL) Dsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Database)
}

func (c *PostgreSQL) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
