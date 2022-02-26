package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"rainbow-umbrella/internal/interfaces"
)

type sessionRepo struct {
	client *redis.Client
}

func NewSessionRepo(redisClient *redis.Client) interfaces.ISessionRepo {
	return &sessionRepo{client: redisClient}
}

func (r sessionRepo) InsertOne(sessionID, login string) error {
	ctx := context.Background()

	if err := r.client.HSet(ctx, sessionID, "login", login).Err(); err != nil {
		return fmt.Errorf("[sessionRepo.InsertOne][1]: %+v", err)
	}
	if err := r.client.Expire(ctx, sessionID, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("[sessionRepo.InsertOne][2]: %+v", err)
	}

	return nil
}

func (r sessionRepo) Exists(sessionID string) (bool, error) {
	ctx := context.Background()

	value, err := r.client.Exists(ctx, sessionID).Result()
	if err != nil {
		return false, fmt.Errorf("[sessionRepo.Exists][1]")
	}

	return value == 1, nil
}
