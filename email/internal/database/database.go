package database

import (
	"github.com/go-redis/redis"
	"github.com/segmentio/kafka-go"

	"dissys/internal/config"
)

type Providers struct {
	Redis *redis.Client
	Kafka *kafka.Conn
}

func New(d *config.Database) (*Providers, error) {
	client, err := getRedis(d.RedisURL, d.RedisPSW)
	if err != nil {
		return nil, err
	}

	conn, err := getKafka(d.KafkaADDR, d.KafkaUSR, d.KafkaPSW)
	if err != nil {
		return nil, err
	}

	return &Providers{
		Redis: client,
		Kafka: conn,
	}, nil
}
