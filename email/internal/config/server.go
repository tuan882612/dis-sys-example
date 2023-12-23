package config

type Server struct {
	Host string `env:"HOST,required"`
	Port string `env:"PORT,required"`
}
