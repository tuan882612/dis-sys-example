package database

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func getKafka(addr, username, password string) (*kafka.Dialer, error) {
	log.Info().Msg("connecting kafka...")

	// create SASL/SCRAM mechanism
	mechanism, err := scram.Mechanism(scram.SHA512, username, password)
	if err != nil {
		log.Error().Err(err).Msg("failed to create mechanism")
		return nil, err
	}

	// create a dialer with SASL/SCRAM authentication
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	// create a new Kafka reader using the broker and dialer
	_, err = dialer.DialContext(context.Background(), "tcp", addr)
	if err != nil {
		log.Error().Err(err).Msg("failed to dial kafka")
		return nil, err
	}

	return dialer, nil
}
