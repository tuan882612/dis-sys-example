package deps

import (
	"github.com/rs/zerolog/log"

	"dissys/internal/config"
	"dissys/internal/database"
)

type Dependencies struct {
	Config   *config.Configuration
	Database *database.DataStores
}

func NewDependencies() (*Dependencies, error) {
	log.Info().Msg("initializing dependencies...")

	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	db, err := database.NewDataStores(cfg.Database)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Config:   cfg,
		Database: db,
	}, nil
}
