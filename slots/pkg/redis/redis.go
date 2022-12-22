package redis

import (
	"context"

	"github.com/cvetkovski98/zvax-slots/internal/config"
	"github.com/go-redis/redis/v9"
)

// NewRedisConn returns a new redis connection
func NewRedisConn(c config.DbConfig, p config.PoolConfig) (*redis.Client, error) {
	opts := &redis.Options{
		// Username:     c.User,
		// Password:     c.Password,
		Addr:         c.Address(),
		DB:           c.Database,
		MinIdleConns: p.MinConn,
		MaxIdleConns: p.MaxConn,
		PoolSize:     p.PoolSize,
	}
	client := redis.NewClient(opts)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
