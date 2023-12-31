package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func getPostgres(ctx context.Context, pgUrl string) (*pgxpool.Pool, error) {
	log.Info().Msg("connecting postgres...")

	// parse and set postgres config
	config, err := pgxpool.ParseConfig(pgUrl)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse postgres config")
		return nil, err
	}

	// connect to postgres database
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to postgres")
		return nil, err
	}

	// ping postgres database to check if connection is alive
	if err := pool.Ping(ctx); err != nil {
		log.Error().Err(err).Msg("failed to ping postgres")
		return nil, err
	}

	return pool, nil
}
