package config

import "fmt"

type RedisPool struct {
	MaxConn  int `mapstructure:"max_conn"`
	MinConn  int `mapstructure:"min_conn"`
	PoolSize int `mapstructure:"pool_size"`
}

type Redis struct {
	Host     string    `mapstructure:"host"`
	Port     int       `mapstructure:"port"`
	User     string    `mapstructure:"user"`
	Password string    `mapstructure:"pass"`
	Database int       `mapstructure:"name"`
	Pool     RedisPool `mapstructure:"pool"`
}

func (c *Redis) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
