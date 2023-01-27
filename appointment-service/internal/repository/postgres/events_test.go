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

func TestStoreDoctor(t *testing.T) {
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
	repo := p.NewEventsRepo(db)
	if err = repo.StoreDoctor(ctx, doctor); err != nil {
		t.Fatalf("error was not expected: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetDoctor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	doctorId := int32(3)
	rows := sqlmock.NewRows([]string{"id", "name", "surname", "phone"}).
		AddRow(doctorId, "Vasya", "Pupkin", "87776542154")

	query := regexp.QuoteMeta(`
		SELECT id, name, surname, phone
		FROM doctors
		WHERE id = $1
	`)

	mock.ExpectBegin()
	mock.ExpectQuery(query).WithArgs(doctorId).WillReturnRows(rows)
	mock.ExpectCommit()

	ctx := context.Background()
	repo := p.NewEventsRepo(db)
	doctors, err := repo.GetDoctor(ctx, doctorId)
	assert.NoError(t, err)
	assert.NotNil(t, doctors)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetDoctorShouldRollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	doctorId := int32(3)

	query := regexp.QuoteMeta(`
		SELECT id, name, surname, phone
		FROM doctors
		WHERE id = $1
	`)

	mock.ExpectBegin()
	mock.ExpectQuery(query).WithArgs(doctorId).WillReturnError(fmt.Errorf("some"))
	mock.ExpectRollback()

	ctx := context.Background()
	repo := p.NewEventsRepo(db)
	if _, err := repo.GetDoctor(ctx, doctorId); err == nil {
		t.Fatal("expected error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateDoctor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	doctorId := int32(1)
	rows := sqlmock.NewRows([]string{"id", "name", "surname", "phone"}).
		AddRow(doctorId, "Vasya", "Pupkin", "87776542154")
	doctor := entity.Doctor{
		Id:      doctorId,
		Name:    "Duhast",
		Surname: "Vicheslavovich",
		Phone:   "8-777-564-87-48",
	}
	rows2 := sqlmock.NewRows([]string{"id", "name", "surname", "phone"}).
		AddRow(doctorId, "Duhast", "Vicheslavovich", "8-777-564-87-48")
	getQuery := regexp.QuoteMeta(`
		SELECT id, name, surname, phone
		FROM doctors
		WHERE id = $1
	`)
	updateQuery := regexp.QuoteMeta(`
		UPDATE doctors
		SET name = $1, surname = $2, phone = $3
		WHERE id = $4
	`)

	mock.ExpectBegin()
	mock.ExpectQuery(getQuery).WithArgs(doctorId).WillReturnRows(rows)
	mock.ExpectExec(updateQuery).WithArgs(doctor.Name, doctor.Surname, doctor.Phone, doctor.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(getQuery).WithArgs(doctorId).WillReturnRows(rows2)
	mock.ExpectCommit()

	ctx := context.Background()
	repo := p.NewEventsRepo(db)

	if _, err := repo.UpdateDoctor(ctx, doctor); err != nil {
		t.Fatalf("error was not expected: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}

}

func TestUpdateDoctorShouldRollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	doctorId := int32(1)
	rows := sqlmock.NewRows([]string{"id", "name", "surname", "phone"}).
		AddRow(doctorId, "Vasya", "Pupkin", "87776542154")
	doctor := entity.Doctor{
		Id:      doctorId,
		Name:    "Duhast",
		Surname: "Vicheslavovich",
		Phone:   "8-777-564-87-48",
	}
	getQuery := regexp.QuoteMeta(`
		SELECT id, name, surname, phone
		FROM doctors
		WHERE id = $1
	`)
	updateQuery := regexp.QuoteMeta(`
		UPDATE doctors
		SET name = $1, surname = $2, phone = $3
		WHERE id = $4
	`)

	mock.ExpectBegin()
	mock.ExpectQuery(getQuery).WithArgs(doctorId).WillReturnRows(rows)
	mock.ExpectExec(updateQuery).WithArgs(doctor.Name, doctor.Surname, doctor.Phone, doctor.Id).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	ctx := context.Background()
	repo := p.NewEventsRepo(db)

	if _, err := repo.UpdateDoctor(ctx, doctor); err == nil {
		t.Fatal("expected error")
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
	repo := p.NewEventsRepo(db)

	if err = repo.DeleteDoctor(ctx, doctorID); err != nil {
		t.Fatalf("error was not expected: %s", err)
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
		SELECT id, name, surname, phone
		FROM doctors
		ORDER BY id
	`)

	prep := mock.ExpectPrepare(query)

	prep.ExpectQuery().WithArgs().WillReturnRows(rows)

	ctx := context.Background()
	repo := p.NewEventsRepo(db)
	doctors, err := repo.FetchDoctors(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, doctors)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
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
	getQuery := regexp.QuoteMeta(`
		SELECT EXISTS (SELECT 1 FROM events WHERE date(starts_at) = $1)
	`)

	insertQuery := regexp.QuoteMeta(`
		INSERT INTO events(doctor_id, starts_at, ends_at)
		VALUES($1, $2, $3)
	`)

	rows := sqlmock.NewRows([]string{"exists"}).
		AddRow(false)
	ctx := context.Background()

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(getQuery)
	prep.ExpectQuery().WithArgs().WillReturnRows(rows)
	prep = mock.ExpectPrepare(insertQuery)
	prep.ExpectExec().WithArgs(events[0].DoctorId, events[0].StartsAt, events[0].EndsAt).
		WillReturnResult(driver.RowsAffected(1))
	mock.ExpectCommit()

	repo := p.NewEventsRepo(db)
	if _, err := repo.StoreSchedule(ctx, events); err != nil {
		t.Fatalf("error was not expected: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}

}

func TestStoreScheduleShouldRollback(t *testing.T) {
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
	query := regexp.QuoteMeta(`
		SELECT EXISTS (SELECT 1 FROM events WHERE date(starts_at) = $1)
	`)

	ctx := context.Background()

	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs().WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()
	repo := p.NewEventsRepo(db)
	if _, err = repo.StoreSchedule(ctx, events); err == nil {
		t.Fatal("exptected error")
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
	repo := p.NewEventsRepo(db)
	events, err := repo.FetchOpenEventsByDoctor(ctx, doctorId)
	assert.NoError(t, err)
	assert.NotNil(t, events)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
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
	repo := p.NewEventsRepo(db)
	events, err := repo.FetchReservedEventsByDoctor(ctx, doctorId)
	assert.NoError(t, err)
	assert.NotNil(t, events)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
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
	repo := p.NewEventsRepo(db)
	events, err := repo.FetchReservedEventsByClient(ctx, clientId)
	assert.NoError(t, err)
	assert.NotNil(t, events)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
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
	repo := p.NewEventsRepo(db)
	events, err := repo.FetchAllEventsByClient(ctx, clientId)
	assert.NoError(t, err)
	assert.NotNil(t, events)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateEvent(t *testing.T) {
	updateQuery := regexp.QuoteMeta(`
		UPDATE events
		SET client_id = $1
		WHERE id = $2
	`)

	getQuery := regexp.QuoteMeta(`
		SELECT id, client_id, doctor_id, starts_at, ends_at
		FROM events
		WHERE id = $1
	`)

	eventId := int32(1)
	clientId := int32(3)

	now := time.Now()
	rows1 := sqlmock.NewRows([]string{"id", "client_id", "doctor_id", "starts_at", "ends_at"}).
		AddRow(eventId, 0, 6, now.Add(time.Hour*10), now.Add(time.Hour*11))
	rows2 := sqlmock.NewRows([]string{"id", "client_id", "doctor_id", "starts_at", "ends_at"}).
		AddRow(eventId, clientId, 6, now.Add(time.Hour*10), now.Add(time.Hour*11))

	event := entity.Event{
		Id:       eventId,
		ClientId: clientId,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectBegin()
	mock.ExpectQuery(getQuery).WithArgs(eventId).WillReturnRows(rows1)
	mock.ExpectExec(updateQuery).WithArgs(clientId, eventId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(getQuery).WithArgs(eventId).WillReturnRows(rows2)
	mock.ExpectCommit()

	ctx := context.Background()
	repo := p.NewEventsRepo(db)
	updated, err := repo.UpdateEvent(ctx, event)

	assert.NoError(t, err)
	assert.NotNil(t, updated)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateEventSouldRollback(t *testing.T) {
	query := regexp.QuoteMeta(`
		SELECT id, client_id, doctor_id, starts_at, ends_at
		FROM events
		WHERE id = $1
	`)

	event := entity.Event{
		Id:       int32(1),
		ClientId: int32(3),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectBegin()
	mock.ExpectQuery(query).WithArgs(event.Id).WillReturnError(entity.ErrEventDoesNotExist)
	mock.ExpectRollback()

	ctx := context.Background()
	repo := p.NewEventsRepo(db)
	if _, err := repo.UpdateEvent(ctx, event); err == nil {
		t.Fatal("expected error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestClearEvent(t *testing.T) {
	updateQuery := regexp.QuoteMeta(`
		UPDATE events
		SET client_id = NULL
		WHERE id = $1
	`)

	getQuery := regexp.QuoteMeta(`
		SELECT id, client_id, doctor_id, starts_at, ends_at
		FROM events
		WHERE id = $1
	`)

	eventId := int32(1)
	clientId := int32(3)

	now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "client_id", "doctor_id", "starts_at", "ends_at"}).
		AddRow(eventId, clientId, 6, now.Add(time.Hour*10), now.Add(time.Hour*11))

	event := entity.Event{
		Id: eventId,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectBegin()
	mock.ExpectQuery(getQuery).WithArgs(eventId).WillReturnRows(rows)
	mock.ExpectExec(updateQuery).WithArgs(eventId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	ctx := context.Background()
	repo := p.NewEventsRepo(db)
	if err = repo.ClearEvent(ctx, event); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}

func TestClearEventShouldRollback(t *testing.T) {
	query := regexp.QuoteMeta(`
		SELECT id, client_id, doctor_id, starts_at, ends_at
		FROM events
		WHERE id = $1
	`)

	event := entity.Event{
		Id: int32(1),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	mock.ExpectBegin()
	mock.ExpectQuery(query).WithArgs(event.Id).WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	ctx := context.Background()
	repo := p.NewEventsRepo(db)
	if err = repo.ClearEvent(ctx, event); err == nil {
		t.Fatal("expected error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %s", err)
	}
}
