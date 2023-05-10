package main

import (
	"first_init/startup"
	cfg "first_init/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
