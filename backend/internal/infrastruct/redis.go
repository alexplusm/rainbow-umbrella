package infrastruct

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewRedisConn(config *RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second) // TODO
	ctx := context.Background()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("[NewRedisConn][1]: %w", err)
	}

	return client, nil
}
