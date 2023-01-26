package postgres_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/config"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/logger"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/postgres"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
	p "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository/postgres"
)

func setup() *sql.DB {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	l := logger.New(cfg.Logger.Level)
	pg, err := postgres.New(cfg)
	if err != nil {
		l.Error(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer func() {
		err = pg.Close()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - pg.Close: %w", err))
		}
	}()
	return pg
}

func TestStoreDocktor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query := `
	INSERT INTO doctors(name, surname, phone)
		VALUES($1, $2, $3)
	`

	docktor := entity.Doctor{
		Name:    "Alice",
		Surname: "Mokovich",
		Phone:   "8-701-784-55-44",
	}
	mock.ExpectExec(query).WithArgs(docktor.Name, docktor.Surname, docktor.Phone).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	repo := p.NewClientsRepo(db)
	if err = repo.StoreDoctor(ctx, docktor); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
