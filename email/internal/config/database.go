package config

type Database struct {
	// redis config
	RedisURL string `env:"REDIS_URL,required"`
	RedisPSW string `env:"REDIS_PSW,required"`
	// kafka config
	KafkaADDR string `env:"KAFKA_ADDR,required"`
	KafkaUSR  string `env:"KAFKA_USR,required"`
	KafkaPSW  string `env:"KAFKA_PSW,required"`
}
