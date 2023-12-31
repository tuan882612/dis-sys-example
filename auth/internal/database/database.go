package database

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"

	"dissys/internal/config"
)

type DataStores struct {
	Redis    *redis.Client
	Kafka    *kafka.Dialer
	Postgres *pgxpool.Pool
}

func NewDataStores(cfg *config.Database) (*DataStores, error) {
	log.Info().Msg("initializing database DataStores...")

	client, err := getRedis(cfg.RedisURL, cfg.RedisPSW)
	if err != nil {
		return nil, err
	}

	dialer, err := getKafka(cfg.KafkaADDR, cfg.KafkaUSR, cfg.KafkaPSW)
	if err != nil {
		return nil, err
	}

	postgres, err := getPostgres(context.Background(), cfg.PgURL)
	if err != nil {
		return nil, err
	}

	return &DataStores{
		Redis:    client,
		Kafka:    dialer,
		Postgres: postgres,
	}, nil
}
