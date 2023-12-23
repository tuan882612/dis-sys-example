package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-envconfig"
)

type Configuration struct {
	*Server   `env:",prefix=SERVER_,required"`
	*Database `env:",prefix=DB_,required"`
	*Email    `env:",prefix=EMAIL_,required"`
}

func New() (*Configuration, error) {
	log.Info().Msg("initializing configuration...")

	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("failed to load .env file")
		return nil, err
	}

	cfg := &Configuration{}

	if err := envconfig.Process(context.Background(), cfg); err != nil {
		log.Error().Err(err).Msg("failed to process configuration")
		return nil, err
	}

	return cfg, nil
}
