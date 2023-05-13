package main

import (
	"res_init/startup"
	cfg "res_init/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
