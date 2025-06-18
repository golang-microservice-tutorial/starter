package main

import (
	"starter-service/cmd"
	"starter-service/config"
	logger "starter-service/pkg"
)

func main() {
	config := config.LoadConfig()
	logger.Init(config)
	cmd.Run(cmd.ServerOptions{Config: config})
}
