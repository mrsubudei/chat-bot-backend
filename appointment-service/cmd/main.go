package main

import (
	"log"

	config "github.com/mrsubudei/chat-bot-backend/appointment-service/config"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
