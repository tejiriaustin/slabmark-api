package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisClient struct {
	rdb *redis.Client
}

func NewRedisClient(dsn, port string) (*RedisClient, error) {
	redisDbClient := redis.NewClient(&redis.Options{
		Network:   port,
		Addr:      dsn,
		TLSConfig: nil,
	})

	if err := redisDbClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Error connecting to redis: %v", err)
		return nil, err
	}

	return &RedisClient{
		rdb: redisDbClient,
	}, nil
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

func (rdb RedisClient) Disconnect(ctx context.Context) error {
	err := rdb.rdb.Close()
	if err != nil {
		return err
	}
	return nil
}
