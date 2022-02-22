package repos

import (
	"github.com/go-redis/redis/v8"

	"rainbow-umbrella/internal/interfaces"
)

type sessionRepo struct {
	client *redis.Client
}

func NewSessionRepo(redisClient *redis.Client) interfaces.ISessionRepo {
	return &sessionRepo{client: redisClient}
}
