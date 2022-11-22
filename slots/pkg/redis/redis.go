package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
)

// NewRedisConn returns a new redis connection
func NewRedisConn() (*redis.Client, error) {
	var client = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "", // no password set
		DB:           0,  // use default DB
		PoolSize:     100,
		MaxIdleConns: 100,
		MinIdleConns: 10,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
