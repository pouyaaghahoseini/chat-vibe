package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context, redisURL string) (*redis.Client, error) {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("parse redis url: %w", err)
	}

	var lastErr error

	for attempt := 1; attempt <= 10; attempt++ {
		client := redis.NewClient(options)

		if err := client.Ping(ctx).Err(); err == nil {
			return client, nil
		} else {
			lastErr = fmt.Errorf("ping redis: %w", err)
			_ = client.Close()
		}

		time.Sleep(2 * time.Second)
	}

	return nil, lastErr
}
