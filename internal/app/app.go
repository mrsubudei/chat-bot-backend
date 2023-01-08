package app

import (
	"fmt"

	"github.com/mrsubudei/chat-bot-backend/config"
	"github.com/mrsubudei/chat-bot-backend/pkg/logger"
	"github.com/mrsubudei/chat-bot-backend/pkg/mysql"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

	// Mysql
	dbConn, err := mysql.New(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mysql.New: %w", err))
		return
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			l.Fatal(fmt.Errorf("app - Run - dbConn.Close: %w", err))
		}
	}()
	fmt.Println("vse good")
}
