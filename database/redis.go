package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisClient struct {
	rdb *redis.Client
}

func NewRedisClient(dsn, password string, opts redis.Options) *RedisClient {
	redisDbClient := redis.NewClient(&redis.Options{
		Network:               "",
		Addr:                  dsn,
		Username:              "",
		Password:              password,
		DB:                    0,
		ContextTimeoutEnabled: false,
		TLSConfig:             nil,
	})

	if err := redisDbClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Error connecting to redis: %v", err)
		return nil
	}

	return &RedisClient{
		rdb: redisDbClient,
	}
}

func (rdb RedisClient) Set(ctx context.Context, key string, value interface{}) error {
	return rdb.rdb.Set(ctx, key, value, 0).Err()
}

func (rdb RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := rdb.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
