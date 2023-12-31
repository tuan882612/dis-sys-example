package deps

import (
	"github.com/rs/zerolog/log"

	"dissys/internal/auth"
	"dissys/internal/config"
	"dissys/internal/database"
	"dissys/internal/auth/jwt"
)

type Dependencies struct {
	Config   *config.Configuration
	Database *database.DataStores
	AuthRepo *auth.Repository
	JWTProv  *jwt.Provider
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

	authRepo := auth.NewRepository(db)

	jwtProv, err := jwt.NewProvider(cfg.JWT)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Config:   cfg,
		Database: db,
		AuthRepo: authRepo,
		JWTProv:  jwtProv,
	}, nil
}
