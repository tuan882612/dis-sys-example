package config

type Email struct {
	SmtpHost string `envconfig:"SMTP_HOST,required"`
	Sender   string `envconfig:"SENDER,required"`
	Password string `envconfig:"PASSWORD,required"`
}
