package postgres_test

import (
	"context"
	"database/sql/driver"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
	p "github.com/mrsubudei/chat-bot-backend/appointment-service/internal/repository/postgres"
)

func TestStoredoctor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query := regexp.QuoteMeta(`
	INSERT INTO doctors(name, surname, phone)
		VALUES($1, $2, $3)
	`)

	doctor := entity.Doctor{
		Name:    "Alice",
		Surname: "Mokovich",
	}
	mock.ExpectExec(query).WithArgs(doctor.Name, doctor.Surname, doctor.Phone).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	repo := p.NewClientsRepo(db)
	if err = repo.StoreDoctor(ctx, doctor); err != nil {
		t.Fatalf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeletedoctor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := regexp.QuoteMeta(`
	DELETE FROM doctors
	WHERE id = $1
	`)

	doctorID := int32(1)

	mock.ExpectExec(query).WithArgs(doctorID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctx := context.Background()
	repo := p.NewClientsRepo(db)

	if err = repo.DeleteDoctor(ctx, doctorID); err != nil {
		t.Fatalf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestFetchdoctors(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "surname", "phone"}).
		AddRow(1, "Vasya", "Pupkin", "87776542154").
		AddRow(2, "Duhast", "Vicheslavovich", "87776542151")

	query := regexp.QuoteMeta(`
		SELECT doctor_id, name, surname, phone
		FROM doctors
		ORDER BY id
	`)

	prep := mock.ExpectPrepare(query)

	prep.ExpectQuery().WithArgs().WillReturnRows(rows)

	ctx := context.Background()
	repo := p.NewClientsRepo(db)
	doctors, err := repo.FetchDoctors(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, doctors)
}

func TestStoreSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	now := time.Now()
	events := []entity.Event{
		{
			DoctorId: 1,
			StartsAt: now,
			EndsAt:   now.Add(time.Second),
		},
	}
	query1 := regexp.QuoteMeta(`
		SELECT EXISTS (SELECT 1 FROM events WHERE date(starts_at) = $1)
	`)

	query2 := regexp.QuoteMeta(`
		INSERT INTO events(doctor_id, starts_at, ends_at)
		VALUES($1, $2, $3)
	`)

	rows := sqlmock.NewRows([]string{"exists"}).
		AddRow(false)
	ctx := context.Background()

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query1)
	prep.ExpectQuery().WithArgs().WillReturnRows(rows)
	prep = mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs(events[0].DoctorId, events[0].StartsAt, events[0].EndsAt).
		WillReturnResult(driver.RowsAffected(1))
	mock.ExpectCommit()

	repo := p.NewClientsRepo(db)
	if err = repo.StoreSchedule(ctx, events); err != nil {
		t.Fatalf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}

}

func TestStoreScheduleShouldRoleBack(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	now := time.Now()
	events := []entity.Event{
		{
			DoctorId: 1,
			StartsAt: now,
			EndsAt:   now.Add(time.Second),
		},
	}
	query1 := regexp.QuoteMeta(`
		SELECT EXISTS (SELECT 1 FROM events WHERE date(starts_at) = $1)
	`)

	query2 := regexp.QuoteMeta(`
		INSERT INTO events(doctor_id, starts_at, ends_at)
		VALUES($1, $2, $3)
	`)

	rows := sqlmock.NewRows([]string{"exists"}).
		AddRow(false)

	ctx := context.Background()

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query1)
	prep.ExpectQuery().WithArgs().WillReturnRows(rows)
	prep = mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs(events[0].DoctorId, events[0].StartsAt, events[0].EndsAt).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()
	repo := p.NewClientsRepo(db)
	if err = repo.StoreSchedule(ctx, events); err == nil {
		t.Fatalf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}

}

func TestFetchOpenEventsByDoctor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	now := time.Now()
	doctorId := int32(3)
	rows := sqlmock.NewRows([]string{"id", "doctor_id", "starts_at", "ends_at"}).
		AddRow(1, doctorId, now, now.Add(time.Hour)).
		AddRow(2, doctorId, now.Add(time.Hour*2), now.Add(time.Hour*3)).
		AddRow(3, doctorId, now.Add(time.Hour*4), now.Add(time.Hour*5))

	query := regexp.QuoteMeta(`
		SELECT id, doctor_id, starts_at, ends_at 
		FROM events
		WHERE doctor_id = $1 AND starts_at > now() AND client_id is NULL
		ORDER BY id
	`)
	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	ctx := context.Background()
	repo := p.NewClientsRepo(db)
	events, err := repo.FetchOpenEventsByDoctor(ctx, doctorId)
	assert.NoError(t, err)
	assert.NotNil(t, events)
}

func TestFetchReservedEventsByDoctor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	now := time.Now()
	doctorId := int32(2)
	rows := sqlmock.NewRows([]string{"id", "client_id", "doctor_id", "starts_at", "ends_at"}).
		AddRow(1, 3, doctorId, now, now.Add(time.Hour)).
		AddRow(2, 2, doctorId, now.Add(time.Hour*2), now.Add(time.Hour*3)).
		AddRow(3, 5, doctorId, now.Add(time.Hour*4), now.Add(time.Hour*5))

	query := regexp.QuoteMeta(`
		SELECT id, client_id, doctor_id, starts_at, ends_at
		FROM events
		WHERE doctor_id = $1 AND starts_at > now() AND client_id >= 1
		ORDER BY id
	`)
	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	ctx := context.Background()
	repo := p.NewClientsRepo(db)
	events, err := repo.FetchReservedEventsByDoctor(ctx, doctorId)
	assert.NoError(t, err)
	assert.NotNil(t, events)
}

func TestFetchReservedEventsByClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	now := time.Now()
	clientId := int32(2)
	rows := sqlmock.NewRows([]string{"id", "client_id", "doctor_id", "starts_at", "ends_at"}).
		AddRow(1, clientId, 2, now, now.Add(time.Hour)).
		AddRow(2, clientId, 5, now.Add(time.Hour*2), now.Add(time.Hour*3)).
		AddRow(3, clientId, 6, now.Add(time.Hour*4), now.Add(time.Hour*5))

	query := regexp.QuoteMeta(`
		SELECT id, client_id, doctor_id, starts_at, ends_at 
		FROM events
		WHERE client_id = $1 AND starts_at > now()
		ORDER BY id
	`)
	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	ctx := context.Background()
	repo := p.NewClientsRepo(db)
	events, err := repo.FetchReservedEventsByClient(ctx, clientId)
	assert.NoError(t, err)
	assert.NotNil(t, events)
}

func TestFetchAllEventsByClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	now := time.Now()
	clientId := int32(2)
	rows := sqlmock.NewRows([]string{"id", "client_id", "doctor_id", "starts_at", "ends_at"}).
		AddRow(1, clientId, 2, now, now.Add(time.Hour)).
		AddRow(2, clientId, 5, now.Add(time.Hour*2), now.Add(time.Hour*3)).
		AddRow(3, clientId, 6, now.Add(time.Hour*4), now.Add(time.Hour*5))

	query := regexp.QuoteMeta(`
		SELECT id, client_id, doctor_id, starts_at, ends_at 
		FROM events
		WHERE client_id = $1
		ORDER BY id
	`)
	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	ctx := context.Background()
	repo := p.NewClientsRepo(db)
	events, err := repo.FetchAllEventsByClient(ctx, clientId)
	assert.NoError(t, err)
	assert.NotNil(t, events)
}

func TestUpdateEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query1 := regexp.QuoteMeta(`
		UPDATE events
		SET client_id = NULL
		WHERE id = $1
	`)

	query2 := regexp.QuoteMeta(`
		UPDATE events
		SET client_id = $1
		WHERE id = $2
	`)

	event1 := entity.Event{
		Id: 1,
	}

	event2 := entity.Event{
		Id:       1,
		ClientId: 4,
	}

	t.Run("OK set client id", func(t *testing.T) {
		mock.ExpectExec(query1).WithArgs(event1.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		ctx := context.Background()
		repo := p.NewClientsRepo(db)

		if err = repo.UpdateEvent(ctx, event1); err != nil {
			t.Fatalf("error was not expected while updating stats: %s", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("OK remove client id", func(t *testing.T) {
		mock.ExpectExec(query2).WithArgs(event2.ClientId, event2.Id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		ctx := context.Background()
		repo := p.NewClientsRepo(db)

		if err = repo.UpdateEvent(ctx, event2); err != nil {
			t.Fatalf("error was not expected while updating stats: %s", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("there were unfulfilled expectations: %s", err)
		}
	})
}
