package logging

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"

	"dissys/internal/config"
)

// handles the publishing of logs to kafka
type Publisher struct {
	writer *kafka.Writer
	topic  string
}

// creates a new log publisher
func NewLogPublisher(cfg *config.Database, dialer *kafka.Dialer) *Publisher {
	// initialize config values
	brokers := []string{cfg.KafkaADDR}
	topic := cfg.KafkaUSR + "-log"

	// setup kafka writer config
	writerCfg := kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
		Dialer:  dialer,
	}

	return &Publisher{
		writer: kafka.NewWriter(writerCfg),
		topic:  topic,
	}
}

func (s *Publisher) WriteLog(msg string) {
	go func() {
		kMsg := kafka.Message{
			Value: []byte(msg),
			Time:  time.Now(),
		}

		err := s.writer.WriteMessages(context.Background(), kMsg)
		if err != nil {
			log.Error().Err(err).Msg("failed to write log")
		}
	}()
}

func (s *Publisher) Setup() {
	// setup logger
	log.Logger = log.Logger.Hook(
		zerolog.HookFunc(
			func(e *zerolog.Event, level zerolog.Level, message string) {
				s.WriteLog(message)
			},
		),
	)
}
