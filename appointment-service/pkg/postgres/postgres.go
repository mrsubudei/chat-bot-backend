// Package postgres implements postgres connection.
package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	config "github.com/mrsubudei/chat-bot-backend/appointment-service/config"
)

// New -.
func New(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.NameDB)
	db, err := sql.Open("pgx", dsn)
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
