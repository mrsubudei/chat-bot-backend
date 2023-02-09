// Package postgres implements postgres connection.
package postgres

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	config "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/config"
)

// New -.
func New(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.Postgres.URL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Postgres.PoolMax)

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
