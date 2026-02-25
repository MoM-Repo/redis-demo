package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"redis-demo/internal/config/section"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	*redis.Client
	cfg section.RepositoryRedis
}

// RedisClient возвращает встроенный *redis.Client для передачи в сервисы
func (c *Client) RedisClient() *redis.Client {
	return c.Client
}

func NewConn(ctx context.Context, cfg section.RepositoryRedis) (*Client, error) {
	log.Printf("Connecting to Redis at %s, DB %d", cfg.Address, cfg.DB)

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	log.Println("Redis connected")

	return &Client{
		Client: client,
		cfg:    cfg,
	}, nil
}
