package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	var lastErr error

	for attempt := 1; attempt <= 10; attempt++ {
		pool, err := pgxpool.New(ctx, databaseURL)
		if err != nil {
			lastErr = fmt.Errorf("create postgres pool: %w", err)
		} else {
			pingErr := pool.Ping(ctx)
			if pingErr == nil {
				return pool, nil
			}

			lastErr = fmt.Errorf("ping postgres: %w", pingErr)
			pool.Close()
		}

		time.Sleep(2 * time.Second)
	}

	return nil, lastErr
}
