package redisclient

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisClientInterface interface {
	AuthSessionCreate(ctx context.Context, req *AuthSession) error
	AuthSessionFind(ctx context.Context, username string) (*AuthSession, error)
	AuthSessionSetSessionId(ctx context.Context, req *AuthSession) (*AuthSession, error)
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(host string, port string, password string, db int) RedisClientInterface {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return &RedisClient{
		client: client,
	}
}
