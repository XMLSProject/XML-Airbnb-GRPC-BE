package config

import "os"

type Config struct {
	Port           string
	OrderingDBHost string
	OrderingDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:           os.Getenv("LOGIN_SERVICE_PORT"),
		OrderingDBHost: os.Getenv("LOGIN_DB_HOST"),
		OrderingDBPort: os.Getenv("LOGIN_DB_PORT"),
	}
}
