package logging

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"

	"dissys/internal/deps"
)

type Service struct {
	writer *kafka.Writer
	topic  string
}

func New(d *deps.Dependencies) *Service {
	// initialize config values
	brokers := []string{d.Config.Database.KafkaADDR}
	topic := d.Config.Database.KafkaUSR + "-log"

	// setup kafka writer config
	writerCfg := kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
		Dialer:  d.Database.Kafka,
	}

	return &Service{
		writer: kafka.NewWriter(writerCfg),
		topic:  topic,
	}
}

func (s *Service) WriteLog(msg string) {
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

func (s *Service) Setup() {
	log.Logger = log.Logger.Hook(zerolog.HookFunc(func(e *zerolog.Event, level zerolog.Level, message string) {
		s.WriteLog(message)
	}))
}
