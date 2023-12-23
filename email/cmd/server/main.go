package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"dissys/internal/deps"
	"dissys/internal/server"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	dps, err := deps.New()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize dependencies")
	}

	svr := server.New(dps)
	if err := svr.Start(); err != nil {
		os.Exit(1)
	}
}
