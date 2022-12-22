package config

import (
	"fmt"
)

type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"pass"`
	Database int    `mapstructure:"name"`
}

func (c *DbConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type PoolConfig struct {
	MaxConn  int `mapstructure:"max_conn"`
	MinConn  int `mapstructure:"min_conn"`
	PoolSize int `mapstructure:"pool_size"`
}
