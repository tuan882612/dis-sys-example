package config

type JWT struct {
	Secret string `env:"SECRET,required"`
}
