package database

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func getKafka(addr, username, password string) (*kafka.Conn, error) {
	log.Info().Msg("connecting kafka...")

	// Create SASL/SCRAM mechanism
	mechanism, err := scram.Mechanism(scram.SHA512, username, password)
	if err != nil {
		log.Error().Msgf("failed to create mechanism: %v", err)
		return nil, err
	}

	// Create a dialer with SASL/SCRAM authentication
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	// Create a new Kafka reader using the broker and dialer
	conn, err := dialer.DialContext(context.Background(), "tcp", addr)
	if err != nil {
		log.Error().Msgf("failed to dial: %v", err)
		return nil, err
	}

	return conn, nil
}
