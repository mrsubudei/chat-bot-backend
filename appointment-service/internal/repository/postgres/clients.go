package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/entity"
)

type ClientsRepo struct {
	*sql.DB
}

func NewClientsRepo(pg *sql.DB) *ClientsRepo {
	return &ClientsRepo{pg}
}

func (cr *ClientsRepo) StoreDoctor(ctx context.Context, doctor entity.Doctor) error {
	res, err := cr.DB.ExecContext(ctx, `
		INSERT INTO doctors(name, surname, phone)
			VALUES($1, $2, $3)
	`, doctor.Name, doctor.Surname, doctor.Phone)
	if err != nil {
		return fmt.Errorf("ClientsRepo - StoreDoctor - ExecContext: %w", err)
	}

	rows, err := res.RowsAffected()
	if rows != 1 || err != nil {
		return fmt.Errorf("ClientsRepo - StoreDoctor - RowsAffected: %w", err)
	}

	return nil
}

func (cr *ClientsRepo) DeleteDoctor(ctx context.Context, id int32) error {
	res, err := cr.DB.ExecContext(ctx, `
		DELETE FROM doctors
		WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("ClientsRepo - DeleteDoctor - ExecContext: %w", err)
	}

	rows, err := res.RowsAffected()
	if rows != 1 || err != nil {
		return fmt.Errorf("ClientsRepo - DeleteDoctor - RowsAffected: %w", err)
	}

	return nil
}

func (cr *ClientsRepo) FetchDoctors(ctx context.Context) ([]entity.Doctor, error) {
	doctors := []entity.Doctor{}
	stmt, err := cr.DB.PrepareContext(ctx, `
		SELECT * FROM doctors
		ORDER BY id
	`)
	if err != nil {
		return nil, fmt.Errorf("ClientsRepo - FetchDoctors - PrepareContext: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ClientsRepo - FetchDoctors - QueryContext: %w", err)
	}

	for rows.Next() {
		var phone sql.NullString
		var doctor entity.Doctor
		err = rows.Scan(&doctor.Id, &doctor.Name, &doctor.Surname, &phone)
		if err == sql.ErrNoRows {
			return nil, entity.ErrNoData
		} else if err != nil {
			return nil, fmt.Errorf("ClientsRepo - FetchDoctors - Scan: %w", err)
		}
		doctor.Phone = phone.String
		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (cr *ClientsRepo) StoreSchedule(ctx context.Context, events []entity.Event) error {
	tx, err := cr.DB.Begin()
	if err != nil {
		return fmt.Errorf("ClientsRepo - StoreSchedule - Begin: %w", err)
	}
	defer tx.Rollback()

	if err := cr.checkConstraintViolation(ctx, tx, events); err != nil {
		return err
	}

	stmt, err := cr.DB.PrepareContext(ctx, `
		INSERT INTO events(doctor_id, starts_at, ends_at)
		VALUES($1, $2, $3)
	`)
	if err != nil {
		return fmt.Errorf("ClientsRepo - StoreSchedule - PrepareContext: %w", err)
	}
	defer stmt.Close()

	for _, v := range events {
		res, err := stmt.ExecContext(ctx, v.DoctorId, v.StartsAt, v.EndsAt)
		if err != nil {
			return fmt.Errorf("ClientsRepo - StoreSchedule - ExecContext: %w", err)
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

func (cr *ClientsRepo) checkConstraintViolation(ctx context.Context, tx *sql.Tx, events []entity.Event) error {
	stmt, err := cr.DB.PrepareContext(ctx, `
		SELECT EXISTS (SELECT 1 FROM events WHERE date(starts_at) = $1)
	`)
	if err != nil {
		return fmt.Errorf("ClientsRepo - checkViolation - PrepareContext: %w", err)
	}
	defer stmt.Close()
	for _, v := range events {
		row := stmt.QueryRowContext(ctx, v.StartsAt)
		if err != nil {
			return fmt.Errorf("ClientsRepo - checkViolation - QueryContext: %w", err)
		}
		exist := false
		err = row.Scan(&exist)
		if err != nil {
			return fmt.Errorf("ClientsRepo - checkViolation - Scan: %w", err)
		}
		if exist {
			return entity.ErrUniqueDateViolation
		}
	}
	return nil
}

func (cr *ClientsRepo) FetchOpenEventsByDoctor(ctx context.Context,
	doctorId int32) ([]entity.Event, error) {
	events := []entity.Event{}
	rows, err := cr.DB.QueryContext(ctx, `
		SELECT id, doctor_id, starts_at, ends_at 
		FROM events
		WHERE doctor_id = $1 AND starts_at > now() AND client_id is NULL
		ORDER BY id
	`, doctorId)
	if err != nil {
		return nil, fmt.Errorf("ClientsRepo - FetchEventsByDoctor - QueryContext: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.Id, &doctorId, &event.StartsAt, &event.EndsAt)
		if err == sql.ErrNoRows {
			return nil, entity.ErrNoData
		} else if err != nil {
			return nil, fmt.Errorf("ClientsRepo - FetchEventsByDoctor - Scan: %w", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func (cr *ClientsRepo) FetchReservedEventsByDoctor(ctx context.Context,
	doctorId int32) ([]entity.Event, error) {
	events := []entity.Event{}
	rows, err := cr.DB.QueryContext(ctx, `
		SELECT * FROM events
		FROM events
		WHERE doctor_id = $1 AND starts_at > now() AND client_id >= 1
		ORDER BY id
	`, doctorId)
	if err != nil {
		return nil, fmt.Errorf("ClientsRepo - FetchEventsByDoctor - QueryContext: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.Id, &event.ClientId, &event.DoctorId,
			&event.StartsAt, &event.EndsAt)
		if err == sql.ErrNoRows {
			return nil, entity.ErrNoData
		} else if err != nil {
			return nil, fmt.Errorf("ClientsRepo - FetchEventsByDoctor - Scan: %w", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func (cr *ClientsRepo) FetchReservedEventsByClient(ctx context.Context,
	clientId int32) ([]entity.Event, error) {
	events := []entity.Event{}
	rows, err := cr.DB.QueryContext(ctx, `
		SELECT * FROM events
		WHERE client_id = $1 AND starts_at > now()
		ORDER BY id
	`, clientId)
	if err != nil {
		return nil, fmt.Errorf("ClientsRepo - FetchEventsByClient - QueryContext: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.Id, &event.ClientId, &event.DoctorId,
			&event.StartsAt, &event.EndsAt)
		if err == sql.ErrNoRows {
			return nil, entity.ErrNoData
		} else if err != nil {
			return nil, fmt.Errorf("ClientsRepo - FetchEventsByClient - Scan: %w", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func (cr *ClientsRepo) FetchAllEventsByClient(ctx context.Context,
	client entity.Client) ([]entity.Event, error) {
	events := []entity.Event{}
	rows, err := cr.DB.QueryContext(ctx, `
		SELECT * FROM events
		WHERE client_id = $1
		ORDER BY id
	`, client.Id)
	if err != nil {
		return nil, fmt.Errorf("ClientsRepo - FetchAllEventsByClient - QueryContext: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		event := entity.Event{}
		err = rows.Scan(&event.Id, &event.ClientId, &event.DoctorId,
			&event.StartsAt, &event.EndsAt)
		if err == sql.ErrNoRows {
			return nil, entity.ErrNoData
		} else if err != nil {
			return nil, fmt.Errorf("ClientsRepo - FetchAllEventsByClient - Scan: %w", err)
		}

		events = append(events, event)
	}

	return events, nil
}

func (cr *ClientsRepo) UpdateEvent(ctx context.Context, event entity.Event) error {
	query := `
		UPDATE events
		SET client_id = $1
		WHERE id = $2
	`

	switch event.ClientId {
	case 0:
		res, err := cr.DB.ExecContext(ctx, query, "NULL", event.Id)
		if err != nil {
			return fmt.Errorf("ClientsRepo - UpdateEvent - ExecContext case #0: %w", err)
		}

		rows, err := res.RowsAffected()
		if rows != 1 || err != nil {
			return fmt.Errorf("ClientsRepo - UpdateEvent - RowsAffected case #0: %w", err)
		}
	default:
		res, err := cr.DB.ExecContext(ctx, query, event.ClientId, event.Id)
		if err != nil {
			return fmt.Errorf("ClientsRepo - UpdateEvent - ExecContext case #default: %w", err)
		}

		rows, err := res.RowsAffected()
		if rows != 1 || err != nil {
			return fmt.Errorf("ClientsRepo - UpdateEvent - RowsAffected case #default: %w", err)
		}
	}

	return nil
}
