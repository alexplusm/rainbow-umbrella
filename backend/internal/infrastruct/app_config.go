package infrastruct

import (
	"fmt"
	"os"
)

type AppConfig struct {
	DatabaseConfig *DatabaseConfig
	RedisConfig    *RedisConfig
}

func (c *AppConfig) BuildFromEnv() (*AppConfig, error) {
	if value := os.Getenv("DATABASE_URL"); value != "" {
		c.DatabaseConfig = &DatabaseConfig{URL: value}
	} else {
		return c, fmt.Errorf("[AppConfig.BuildFromEnv][1]: database url required")
	}

	if value := os.Getenv("REDIS_ADDR"); value != "" {
		c.RedisConfig = &RedisConfig{Addr: value}
	} else {
		return c, fmt.Errorf("[AppConfig.BuildFromEnv][2]: redis addr required")
	}

	return c, nil
}

type DatabaseConfig struct {
	URL string
}

type RedisConfig struct {
	Addr string
}
