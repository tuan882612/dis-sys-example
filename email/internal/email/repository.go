package email

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"

	"dissys/internal/email/smtp"
)

type repository struct {
	cache *redis.Client
}

func NewRepository(cache *redis.Client) *repository {
	log.Info().Msg("initializing repo (cache) repository...")
	return &repository{
		cache: cache,
	}
}

func (r *repository) StoreTwoFAData(userID, status string, code int) error {
	// initialize and serialize the 2FA data
	data, err := smtp.NewTwoFAData(code, 5, status).Serialize()
	if err != nil {
		return err
	}

	// set the key and ttl 
	ttl := time.Duration(60*5) * time.Second // 5 minutes
	key := "auth:" + userID

	// store the data
	cmd := r.cache.Set(key, data, ttl)
	if err := cmd.Err(); err != nil {
		log.Error().Err(err).Msgf("%v: failed to store two factor data", userID)
		return err
	}

	return nil
}
