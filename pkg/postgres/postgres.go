// Package postgres implements postgres connection.
package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/mrsubudei/chat-bot-backend/config"
)

// Postgres -.
type Postgres struct {
	DB *sql.DB
}

// New -.
func New(cfg *config.Config) (*Postgres, error) {
	db, err := sql.Open("pgx", cfg.Postgres.URL)
	fmt.Println(cfg.Postgres.URL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Postgres.PoolMax)

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &Postgres{db}, nil
}

func (pg *Postgres) Close() error {
	return pg.DB.Close()
}
