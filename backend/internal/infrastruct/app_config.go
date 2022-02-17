package infrastruct

import (
	"fmt"
	"os"
)

type AppConfig struct {
	DatabaseConfig *DatabaseConfig
}

func (c *AppConfig) BuildFromEnv() (*AppConfig, error) {
	if value := os.Getenv("DATABASE_URL"); value != "" {
		c.DatabaseConfig = &DatabaseConfig{URL: value}
	} else {
		return c, fmt.Errorf("[AppConfig.BuildFromEnv][1]: database url required")
	}

	return c, nil
}

type DatabaseConfig struct {
	URL string
}
