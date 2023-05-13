package config

import "os"

type Config struct {
	Port           string
	OrderingDBHost string
	OrderingDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:           os.Getenv("RESERVATION_SERVICE_PORT"),
		OrderingDBHost: os.Getenv("RESERVATION_DB_HOST"),
		OrderingDBPort: os.Getenv("RESERVATION_DB_PORT"),
	}
}
