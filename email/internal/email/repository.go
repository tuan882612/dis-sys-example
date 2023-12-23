package email

import "github.com/go-redis/redis"

type repository struct {
	cache *redis.Client
}

func NewRepository(cache *redis.Client) *repository {
	return &repository{
		cache: cache,
	}
}

func (r *repository) Store2FAPayload() {}
