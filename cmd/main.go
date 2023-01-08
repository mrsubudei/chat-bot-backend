package main

import (
	"log"

	"github.com/mrsubudei/chat-bot-backend/config"
	"github.com/mrsubudei/chat-bot-backend/internal/app"
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
