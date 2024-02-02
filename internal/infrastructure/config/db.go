package config

type Database struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME"`
	SSLMode  string `env:"SSL_MODE"`
}
