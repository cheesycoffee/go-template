package database

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	redisInstance struct {
		read, write *redis.Client
	}

	// RedisOption parameter
	RedisOption struct {
		ReadDSN, WriteDSN string
		ConnectTimeout    time.Duration
	}

	// Redis clients
	Redis interface {
		GetReadClient() *redis.Client
		GetWriteClient() *redis.Client
	}
)

func newRedis(option RedisOption) Redis {
	return &redisInstance{
		read:  getRedisClient(option.ReadDSN, option.ConnectTimeout),
		write: getRedisClient(option.WriteDSN, option.ConnectTimeout),
	}
}

func getRedisClient(dsn string, timeout time.Duration) *redis.Client {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	status := redisClient.Ping(ctx)
	if status.Err() != nil {
		panic(err)
	}

	return redisClient
}

func (r *redisInstance) GetReadClient() *redis.Client {
	return r.read
}

func (r *redisInstance) GetWriteClient() *redis.Client {
	return r.write
}
