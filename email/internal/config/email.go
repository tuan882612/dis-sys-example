package config

type Email struct {
	SmtpHost string `env:"SMTP_HOST,required"`
	Sender   string `env:"SENDER,required"`
	Password string `env:"PASSWORD,required"`
}
