package database

import (
	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"

	"dissys/internal/config"
)

type Providers struct {
	Redis *redis.Client
	Kafka *kafka.Dialer
}

func New(d *config.Database) (*Providers, error) {
	log.Info().Msg("initializing database providers...")

	client, err := getRedis(d.RedisURL, d.RedisPSW)
	if err != nil {
		return nil, err
	}

	dialer, err := getKafka(d.KafkaADDR, d.KafkaUSR, d.KafkaPSW)
	if err != nil {
		return nil, err
	}

	return &Providers{
		Redis: client,
		Kafka: dialer,
	}, nil
}
