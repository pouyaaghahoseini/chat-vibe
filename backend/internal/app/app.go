package app

import (
	"context"
	"fmt"

	"chatvibe/backend/internal/config"
	"chatvibe/backend/internal/persistence"
	"chatvibe/backend/internal/store"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type App struct {
	Config       config.Config
	PostgresPool *pgxpool.Pool
	RedisClient  *redis.Client
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	postgresPool, err := store.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("initialize postgres: %w", err)
	}

	if err := persistence.RunMigrations(ctx, postgresPool); err != nil {
		postgresPool.Close()
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	redisClient, err := store.NewRedisClient(ctx, cfg.RedisURL)
	if err != nil {
		postgresPool.Close()
		return nil, fmt.Errorf("initialize redis: %w", err)
	}

	return &App{
		Config:       cfg,
		PostgresPool: postgresPool,
		RedisClient:  redisClient,
	}, nil
}

func (a *App) Close() {
	if a.PostgresPool != nil {
		a.PostgresPool.Close()
	}

	if a.RedisClient != nil {
		_ = a.RedisClient.Close()
	}
}
