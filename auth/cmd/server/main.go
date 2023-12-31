package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"dissys/internal/deps"
	"dissys/internal/observability/logging"
	"dissys/internal/server"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Str("service", "auth").Logger()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	dps, err := deps.NewDependencies()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize dependencies")
	}

	loggingSvc := logging.NewLogPublisher(dps.Config.Database, dps.Database.Kafka)
	loggingSvc.Setup()

	svr := server.New(dps)
	if err := svr.Start(); err != nil {
		os.Exit(1)
	}
}
