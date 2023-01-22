package postgres

import (
	"context"
	"fmt"

	"github.com/mrsubudei/chat-bot-backend/internal/entity"
	"github.com/mrsubudei/chat-bot-backend/pkg/postgres"
)

type ClientsRepo struct {
	*postgres.Postgres
}

func NewClientsRepo(pg *postgres.Postgres) *ClientsRepo {
	return &ClientsRepo{pg}
}

func (cr *ClientsRepo) StoreSchedule(ctx context.Context, events []entity.Event) error {
	tx, err := cr.DB.Begin()
	if err != nil {
		return fmt.Errorf("ClientsRepo - StoreSchedule - Begin: %w", err)
	}
	defer func() {
		err = tx.Rollback()
	}()

	stmt, err := cr.DB.PrepareContext(ctx, `
		INSERT INTO events(starts_at, ends_at)
			VALUES($1, $2)
	`)
	if err != nil {
		return fmt.Errorf("ClientsRepo - StoreSchedule - PrepareContext: %w", err)
	}
	defer stmt.Close()

	for _, v := range events {
		res, err := stmt.ExecContext(ctx, v.StartsAt, v.EndsAt)
		if err != nil {
			return fmt.Errorf("ClientsRepo - StoreSchedule - Exec: %w", err)
		}
		affected, err := res.RowsAffected()
		if affected != 1 || err != nil {
			return fmt.Errorf("ClientsRepo - StoreSchedule - RowsAffected: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("ClientsRepo - StoreSchedule - Commit: %w", err)
	}

	return nil
}
