package config

import "os"

type Config struct {
	Port           string
	OrderingDBHost string
	OrderingDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:           os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		OrderingDBHost: os.Getenv("ACCOMMODATION_DB_HOST"),
		OrderingDBPort: os.Getenv("ACCOMMODATION_DB_PORT"),
	}
}
