package database

import (
	"github.com/go-redis/redis"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"

	"dissys/internal/config"
)

type DataStores struct {
	Redis *redis.Client
	Kafka *kafka.Dialer
}

func NewDataStores(cfg *config.Database) (*DataStores, error) {
	log.Info().Msg("initializing database datastores...")

	client, err := getRedis(cfg.RedisURL, cfg.RedisPSW)
	if err != nil {
		return nil, err
	}

	dialer, err := getKafka(cfg.KafkaADDR, cfg.KafkaUSR, cfg.KafkaPSW)
	if err != nil {
		return nil, err
	}

	return &DataStores{
		Redis: client,
		Kafka: dialer,
	}, nil
}
