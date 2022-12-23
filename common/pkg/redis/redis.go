package redis

import (
	"context"

	"github.com/cvetkovski98/zvax-common/pkg/config"
	"github.com/go-redis/redis/v9"
)

// NewRedisConn returns a new redis connection
func NewRedisConn(cfg config.Redis) (*redis.Client, error) {
	opts := &redis.Options{
		// Username:     c.User,
		// Password:     c.Password,
		Addr:         cfg.Address(),
		DB:           cfg.Database,
		MinIdleConns: cfg.Pool.MinConn,
		MaxIdleConns: cfg.Pool.MaxConn,
		PoolSize:     cfg.Pool.PoolSize,
	}
	client := redis.NewClient(opts)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
