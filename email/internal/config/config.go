package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-envconfig"
)

type Configuration struct {
	*Server   `env:",prefix=SERVER_"`
	*Database `env:",prefix=DB_"`
	*Email    `env:",prefix=EMAIL_"`
}

func New() (*Configuration, error) {
	if err := godotenv.Load(); err != nil {
		log.Error().Str("service", "auth").Msg("failed to load .env file")
		return nil, err
	}

	cfg := &Configuration{}

	if err := envconfig.Process(context.Background(), cfg); err != nil {
		log.Error().Msg("failed to process configuration")
		return nil, err
	}

	return cfg, nil
}
