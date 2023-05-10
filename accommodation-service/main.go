package main

import (
	"accomm_module/startup"
	cfg "accomm_module/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
