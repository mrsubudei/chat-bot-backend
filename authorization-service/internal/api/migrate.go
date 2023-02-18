package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/internal/config"
	"github.com/mrsubudei/chat-bot-backend/authorization-service/pkg/logger"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(cfg *config.Config, l *logger.Logger) error {
	databaseURL := cfg.Postgres.URL

	databaseURL += "?sslmode=disable"

	var (
		attempts = DefaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)
		if err == nil {
			break
		}

		l.Info("postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(DefaultTimeout)
		attempts--
	}

	if err != nil {
		return fmt.Errorf("postgres connect error: %w", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("up error: %w", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("no change: %w", err)
	}

	l.Info("Migrate: up success")
	return nil
}
