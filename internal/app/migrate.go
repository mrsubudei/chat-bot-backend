package app

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/mrsubudei/chat-bot-backend/config"

	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func Migrate(cfg *config.Config) error {
	databaseURL := cfg.Postgres.URL

	databaseURL += "?sslmode=disable"

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		return fmt.Errorf("Migrate: postgres connect error: %w", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("Migrate: up error: %w", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("Migrate: no change: %w", err)
	}

	log.Printf("Migrate: up success")
	return nil
}
