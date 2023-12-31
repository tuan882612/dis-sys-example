package config

type Server struct {
	Host        string `env:"HOST,required"`
	Port        string `env:"PORT,required"`
	MailRpcHost string `env:"MAIL_RPC_HOST,required"`
	MailRpcPort string `env:"MAIL_RPC_PORT,required"`
}
