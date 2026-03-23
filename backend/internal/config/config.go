package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppEnv                     string
	BackendPort                string
	DatabaseURL                string
	RedisURL                   string
	MaxSubscribersPerChannel   int
	AISummaryWindowSeconds     int
	ChannelIdleShutdownSeconds int
}

func Load() (Config, error) {
	cfg := Config{
		AppEnv:      getEnv("APP_ENV", "development"),
		BackendPort: getEnv("BACKEND_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		RedisURL:    getEnv("REDIS_URL", ""),
	}

	var err error

	cfg.MaxSubscribersPerChannel, err = getEnvAsInt("MAX_SUBSCRIBERS_PER_CHANNEL", 25)
	if err != nil {
		return Config{}, fmt.Errorf("invalid MAX_SUBSCRIBERS_PER_CHANNEL: %w", err)
	}

	cfg.AISummaryWindowSeconds, err = getEnvAsInt("AI_SUMMARY_WINDOW_SECONDS", 30)
	if err != nil {
		return Config{}, fmt.Errorf("invalid AI_SUMMARY_WINDOW_SECONDS: %w", err)
	}

	cfg.ChannelIdleShutdownSeconds, err = getEnvAsInt("CHANNEL_IDLE_SHUTDOWN_SECONDS", 0)
	if err != nil {
		return Config{}, fmt.Errorf("invalid CHANNEL_IDLE_SHUTDOWN_SECONDS: %w", err)
	}

	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}

	if cfg.RedisURL == "" {
		return Config{}, fmt.Errorf("REDIS_URL is required")
	}

	return cfg, nil
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getEnvAsInt(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return parsed, nil
}
