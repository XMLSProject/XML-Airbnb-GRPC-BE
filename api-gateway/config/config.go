package config

import "os"

type Config struct {
	Address                     string
	GreeterServiceAddress       string
	LoginServiceAddress         string
	AccommodationServiceAddress string
	ReservationServiceAddress   string
}

func GetConfig() Config {
	return Config{
		GreeterServiceAddress:       os.Getenv("GREETER_SERVICE_ADDRESS"),
		Address:                     os.Getenv("GATEWAY_ADDRESS"),
		LoginServiceAddress:         os.Getenv("LOGIN_SERVICE_ADDRESS"),
		AccommodationServiceAddress: os.Getenv("ACCOMMODATION_SERVICE_ADDRESS"),
		ReservationServiceAddress:   os.Getenv("RESERVATION_SERVICE_ADDRESS"),
	}
}
