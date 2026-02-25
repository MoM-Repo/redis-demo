package config

import (
	"os"
	"strconv"
	"time"

	"redis-demo/internal/config/section"
)

// Config содержит конфигурацию приложения
type Config struct {
	App        section.App
	Repository section.Repository
}

func LoadConfig() *Config {
	cfg := &Config{
		App: section.App{
			ServerPort: os.Getenv("SERVER_PORT"),
		},
		Repository: section.Repository{
			Postgres: section.RepositoryPostgres{
				Host:         os.Getenv("DB_HOST"),
				Port:         os.Getenv("DB_PORT"),
				Username:     os.Getenv("DB_USER"),
				Password:     os.Getenv("DB_PASSWORD"),
				Name:         os.Getenv("DB_NAME"),
				ReadTimeout:  30 * time.Second,
				WriteTimeout: 10 * time.Second,
			},
			Redis: section.RepositoryRedis{
				Address:  os.Getenv("REDIS_ADDRESS"),
				Password: os.Getenv("REDIS_PASSWORD"),
				DB:       mustAtoi(os.Getenv("REDIS_DB"), 0),
			},
		},
	}

	return cfg
}

func mustAtoi(s string, defaultVal int) int {
	if s == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return defaultVal
	}
	return n
}
