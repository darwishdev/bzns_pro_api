package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type AuthCacheInterface interface {
	AuthSessionCreate(ctx context.Context, req *AuthSession) error
	AuthSessionFind(ctx context.Context, username string) (*AuthSession, error)
}

type AuthCache struct {
	client *redis.Client
}

func NewAuthCache(client *redis.Client) AuthCacheInterface {
	return &AuthCache{
		client: client,
	}
}
