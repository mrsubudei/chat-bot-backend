package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
)

type EventsRepo struct {
	*sql.DB
}

func NewEventsRepo(pg *sql.DB) *EventsRepo {
	return &EventsRepo{pg}
}

func (er *EventsRepo) StoreDoctor(ctx context.Context, doctor entity.Doctor) error {
	res, err := er.DB.ExecContext(ctx, `
		INSERT INTO doctors(name, surname, phone)
			VALUES($1, $2, $3)
	`, doctor.Name, doctor.Surname, doctor.Phone)
	if err != nil {
		return fmt.Errorf("EventsRepo - StoreDoctor - ExecContext: %w", err)
	}

	rows, err := res.RowsAffected()
	if rows != 1 || err != nil {
		return fmt.Errorf("EventsRepo - StoreDoctor - RowsAffected: %w", err)
	}

	return nil
}

func (er *EventsRepo) GetDoctor(ctx context.Context, doctorId int32) (entity.Doctor, error) {
	doctor := entity.Doctor{}
	query := `
		SELECT id, name, surname, phone
		FROM doctors
		WHERE id = $1
	`

	res := er.DB.QueryRowContext(ctx, query, doctorId)

	var phone sql.NullString
	err := res.Scan(&doctor.Id, &doctor.Name, &doctor.Surname, &phone)

	if errors.Is(err, sql.ErrNoRows) {
		return doctor, entity.ErrEntityDoesNotExist
	} else if err != nil {
		return doctor, fmt.Errorf("EventsRepo - GetDoctor - Scan: %w", err)
	}
	doctor.Phone = phone.String

	return doctor, nil
}

func (er *EventsRepo) UpdateDoctor(ctx context.Context,
	doctor entity.Doctor) (entity.Doctor, error) {

	tx, err := er.DB.Begin()
	if err != nil {
		return doctor, fmt.Errorf("EventsRepo - UpdateDoctor - Begin: %w", err)
	}
	defer tx.Rollback()

	exist, err := er.GetDoctor(ctx, doctor.Id)
	if err != nil {
		return doctor, fmt.Errorf("EventsRepo - UpdateDoctor - GetDoctor #1 : %w", err)
	}

	if doctor.Name == "" {
		doctor.Name = exist.Name
	}
	if doctor.Surname == "" {
		doctor.Surname = exist.Surname
	}
	if doctor.Phone == "" {
		doctor.Phone = exist.Phone
	}

	query := `
		UPDATE doctors
		SET name = $1, surname = $2, phone = $3
		WHERE id = $4
	`

	res, err := tx.ExecContext(ctx, query, doctor.Name, doctor.Surname,
		doctor.Phone, doctor.Id)
	if err != nil {
		return doctor, fmt.Errorf("EventsRepo - UpdateDoctor - ExecContext: %w", err)
	}

	rows, err := res.RowsAffected()
	if rows != 1 || err != nil {
		return doctor, fmt.Errorf("EventsRepo - UpdateDoctor - RowsAffected: %w", err)
	}

	updated, err := er.GetDoctor(ctx, doctor.Id)
	if err != nil {
		return doctor, fmt.Errorf("EventsRepo - UpdateDoctor - GetDoctor #2: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return doctor, fmt.Errorf("EventsRepo - UpdateDoctor - Commit: %w", err)
	}

	return updated, nil
}

func (er *EventsRepo) DeleteDoctor(ctx context.Context, id int32) error {
	res, err := er.DB.ExecContext(ctx, `
		DELETE FROM doctors
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("EventsRepo - DeleteDoctor - ExecContext: %w", err)
	}

	rows, err := res.RowsAffected()
	if rows != 1 || err != nil {
		return fmt.Errorf("EventsRepo - DeleteDoctor - RowsAffected: %w", err)
	}

	return nil
}

func (er *EventsRepo) FetchDoctors(ctx context.Context) ([]entity.Doctor, error) {
	doctors := []entity.Doctor{}
	stmt, err := er.DB.PrepareContext(ctx, `
		SELECT id, name, surname, phone
		FROM doctors
		ORDER BY id
	`)
	if err != nil {
		return nil, fmt.Errorf("EventsRepo - FetchDoctors - PrepareContext: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventsRepo - FetchDoctors - QueryContext: %w", err)
	}

	for rows.Next() {
		var phone sql.NullString
		var doctor entity.Doctor
		err = rows.Scan(&doctor.Id, &doctor.Name, &doctor.Surname, &phone)
		if err != nil {
			return nil, fmt.Errorf("EventsRepo - FetchDoctors - Scan: %w", err)
		}
		doctor.Phone = phone.String
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (er *EventsRepo) StoreSchedule(ctx context.Context, events []entity.Event) error {
	tx, err := er.DB.Begin()
	if err != nil {
		return fmt.Errorf("EventsRepo - StoreSchedule - Begin: %w", err)
	}
	defer tx.Rollback()

	if err := er.checkConstraintViolation(ctx, tx, events); err != nil {
		return err
	}

	stmt, err := er.DB.PrepareContext(ctx, `
		INSERT INTO events(doctor_id, starts_at, ends_at)
		VALUES($1, $2, $3)
	`)
	if err != nil {
		return fmt.Errorf("EventsRepo - StoreSchedule - PrepareContext: %w", err)
	}
	defer stmt.Close()

	for _, v := range events {
		res, err := stmt.ExecContext(ctx, v.DoctorId, v.StartsAt, v.EndsAt)
		if err != nil {
			return fmt.Errorf("EventsRepo - StoreSchedule - ExecContext: %w", err)
		}
		affected, err := res.RowsAffected()
		if affected != 1 || err != nil {
			return fmt.Errorf("EventsRepo - StoreSchedule - RowsAffected: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("EventsRepo - StoreSchedule - Commit: %w", err)
	}

	return nil
}

func (er *EventsRepo) checkConstraintViolation(ctx context.Context, tx *sql.Tx, events []entity.Event) error {
	stmt, err := er.DB.PrepareContext(ctx, `
		SELECT EXISTS (SELECT 1 FROM events WHERE date(starts_at) = $1)
	`)
	if err != nil {
		return fmt.Errorf("EventsRepo - checkViolation - PrepareContext: %w", err)
	}
	defer stmt.Close()
	for _, v := range events {
		row := stmt.QueryRowContext(ctx, v.StartsAt)
		if err != nil {
			return fmt.Errorf("EventsRepo - checkViolation - QueryContext: %w", err)
		}
		exist := false
		err = row.Scan(&exist)
		if err != nil {
			return fmt.Errorf("EventsRepo - checkViolation - Scan: %w", err)
		}
		if exist {
			return entity.ErrDateAlreadyExists
		}
	}
	return nil
}

func (er *EventsRepo) FetchOpenEventsByDoctor(ctx context.Context,
	doctorId int32) ([]entity.Event, error) {
	events := []entity.Event{}
	rows, err := er.DB.QueryContext(ctx, `
		SELECT id, doctor_id, starts_at, ends_at 
		FROM events
		WHERE doctor_id = $1 AND starts_at > now() AND client_id is NULL
		ORDER BY id
	`, doctorId)
	if err != nil {
		return nil, fmt.Errorf("EventsRepo - FetchEventsByDoctor - QueryContext: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.Id, &event.DoctorId, &event.StartsAt, &event.EndsAt)
		if err != nil {
			return nil, fmt.Errorf("EventsRepo - FetchEventsByDoctor - Scan: %w", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func (er *EventsRepo) FetchReservedEventsByDoctor(ctx context.Context,
	doctorId int32) ([]entity.Event, error) {
	events := []entity.Event{}
	rows, err := er.DB.QueryContext(ctx, `
		SELECT id, client_id, doctor_id, starts_at, ends_at
		FROM events
		WHERE doctor_id = $1 AND starts_at > now() AND client_id >= 1
		ORDER BY id
	`, doctorId)
	if err != nil {
		return nil, fmt.Errorf("EventsRepo - FetchEventsByDoctor - QueryContext: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.Id, &event.ClientId, &event.DoctorId,
			&event.StartsAt, &event.EndsAt)
		if err != nil {
			return nil, fmt.Errorf("EventsRepo - FetchEventsByDoctor - Scan: %w", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func (er *EventsRepo) FetchReservedEventsByClient(ctx context.Context,
	clientId int32) ([]entity.Event, error) {
	events := []entity.Event{}
	rows, err := er.DB.QueryContext(ctx, `
		SELECT id, client_id, doctor_id, starts_at, ends_at 
		FROM events
		WHERE client_id = $1 AND starts_at > now()
		ORDER BY id
	`, clientId)
	if err != nil {
		return nil, fmt.Errorf("EventsRepo - FetchEventsByClient - QueryContext: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.Id, &event.ClientId, &event.DoctorId,
			&event.StartsAt, &event.EndsAt)
		if err != nil {
			return nil, fmt.Errorf("EventsRepo - FetchEventsByClient - Scan: %w", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func (er *EventsRepo) FetchAllEventsByClient(ctx context.Context,
	clientId int32) ([]entity.Event, error) {
	events := []entity.Event{}
	rows, err := er.DB.QueryContext(ctx, `
		SELECT id, client_id, doctor_id, starts_at, ends_at 
		FROM events
		WHERE client_id = $1
		ORDER BY id
	`, clientId)
	if err != nil {
		return nil, fmt.Errorf("EventsRepo - FetchAllEventsByClient - QueryContext: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.Id, &event.ClientId, &event.DoctorId,
			&event.StartsAt, &event.EndsAt)
		if err != nil {
			return nil, fmt.Errorf("EventsRepo - FetchAllEventsByClient - Scan: %w", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func (er *EventsRepo) GetEventById(ctx context.Context, eventId int32) (entity.Event, error) {
	event := entity.Event{}
	query := `
		SELECT id, client_id, doctor_id, starts_at, ends_at
		FROM events
		WHERE id = $1
	`

	res := er.DB.QueryRowContext(ctx, query, eventId)
	var clientId sql.NullInt32

	err := res.Scan(&event.Id, &clientId, &event.DoctorId, &event.StartsAt, &event.EndsAt)
	if errors.Is(err, sql.ErrNoRows) {
		return event, entity.ErrEntityDoesNotExist
	} else if err != nil {
		return event, fmt.Errorf("EventsRepo - GetEventById - Scan: %w", err)
	}
	event.ClientId = clientId.Int32

	return event, nil
}

func (er *EventsRepo) UpdateEvent(ctx context.Context, event entity.Event) (entity.Event, error) {
	tx, err := er.DB.Begin()
	if err != nil {
		return event, fmt.Errorf("EventsRepo - UpdateEvent - Begin: %w", err)
	}
	defer tx.Rollback()

	query1 := `
		UPDATE events
		SET client_id = NULL
		WHERE id = $1
	`

	query2 := `
		UPDATE events
		SET client_id = $1
		WHERE id = $2
	`

	switch event.ClientId {
	case 0:
		res, err := tx.ExecContext(ctx, query1, event.Id)
		if err != nil {
			return event, fmt.Errorf("EventsRepo - UpdateEvent - case #0 ExecContext: %w", err)
		}

		rows, err := res.RowsAffected()
		if rows != 1 || err != nil {
			return event, fmt.Errorf("EventsRepo - UpdateEvent - case #0 RowsAffected: %w", err)
		}
	default:
		res, err := tx.ExecContext(ctx, query2, event.ClientId, event.Id)
		if err != nil {
			return event, fmt.Errorf("EventsRepo - UpdateEvent - case #default ExecContext: %w", err)
		}

		rows, err := res.RowsAffected()
		if rows != 1 || err != nil {
			return event, fmt.Errorf("EventsRepo - UpdateEvent - case #default RowsAffected: %w", err)
		}
	}

	updated, err := er.GetEventById(ctx, event.Id)
	if err != nil {
		return event, fmt.Errorf("EventsRepo - UpdateEvent - %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return event, fmt.Errorf("EventsRepo - UpdateEvent - Commit: %w", err)
	}

	return updated, nil
}
