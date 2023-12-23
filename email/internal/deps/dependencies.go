package deps

import (
	"github.com/rs/zerolog/log"

	"dissys/internal/config"
	"dissys/internal/database"
)

type Dependencies struct {
	Config   *config.Configuration
	Database *database.Providers
}

func New() (*Dependencies, error) {
	log.Info().Msg("initializing dependencies...")

	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	db, err := database.New(cfg.Database)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Config:   cfg,
		Database: db,
	}, nil
}
